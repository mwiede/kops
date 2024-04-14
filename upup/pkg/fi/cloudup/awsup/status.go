/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsup

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"k8s.io/klog/v2"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/protokube/pkg/etcd"
	"k8s.io/kops/upup/pkg/fi"
)

// FindClusterStatus discovers the status of the cluster, by looking for the tagged etcd volumes
func (c *awsCloudImplementation) FindClusterStatus(cluster *kops.Cluster) (*kops.ClusterStatus, error) {
	etcdStatus, err := findEtcdStatus(c, cluster)
	if err != nil {
		return nil, err
	}
	status := &kops.ClusterStatus{
		EtcdClusters: etcdStatus,
	}
	klog.V(2).Infof("Cluster status (from cloud): %v", fi.DebugAsJsonString(status))
	return status, nil
}

// FindEtcdStatus discovers the status of the cluster, by looking for the tagged etcd volumes
func (c *MockAWSCloud) FindClusterStatus(cluster *kops.Cluster) (*kops.ClusterStatus, error) {
	etcdStatus, err := findEtcdStatus(c, cluster)
	if err != nil {
		return nil, err
	}
	return &kops.ClusterStatus{
		EtcdClusters: etcdStatus,
	}, nil
}

// findEtcdStatus discovers the status of etcd, by looking for the tagged etcd volumes
func findEtcdStatus(c AWSCloud, cluster *kops.Cluster) ([]kops.EtcdClusterStatus, error) {
	klog.V(2).Infof("Querying AWS for etcd volumes")
	statusMap := make(map[string]*kops.EtcdClusterStatus)

	tags := c.Tags()

	request := &ec2.DescribeVolumesInput{}
	for k, v := range tags {
		request.Filters = append(request.Filters, NewEC2Filter("tag:"+k, v))
	}

	var volumes []ec2types.Volume
	klog.V(2).Infof("Listing EC2 Volumes")
	paginator := ec2.NewDescribeVolumesPaginator(c.EC2(), request)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("error describing volumes: %v", err)
		}
		volumes = append(volumes, page.Volumes...)
	}

	var err error
	for _, volume := range volumes {
		volumeID := aws.ToString(volume.VolumeId)

		etcdClusterName := ""
		var etcdClusterSpec *etcd.EtcdClusterSpec
		master := false
		for _, tag := range volume.Tags {
			k := aws.ToString(tag.Key)
			v := aws.ToString(tag.Value)

			if strings.HasPrefix(k, TagNameEtcdClusterPrefix) {
				etcdClusterName = strings.TrimPrefix(k, TagNameEtcdClusterPrefix)
				etcdClusterSpec, err = etcd.ParseEtcdClusterSpec(etcdClusterName, v)
				if err != nil {
					return nil, fmt.Errorf("error parsing etcd cluster tag %q on volume %q: %v", v, volumeID, err)
				}
			} else if k == TagNameRolePrefix+TagRoleMaster || k == TagNameRolePrefix+TagRoleControlPlane {
				master = true
			}
		}
		if etcdClusterName == "" || etcdClusterSpec == nil || !master {
			continue
		}

		status := statusMap[etcdClusterName]
		if status == nil {
			status = &kops.EtcdClusterStatus{
				Name: etcdClusterName,
			}
			statusMap[etcdClusterName] = status
		}

		memberName := etcdClusterSpec.NodeName
		status.Members = append(status.Members, &kops.EtcdMemberStatus{
			Name:     memberName,
			VolumeID: aws.ToString(volume.VolumeId),
		})
	}

	var status []kops.EtcdClusterStatus
	for _, v := range statusMap {
		status = append(status, *v)
	}
	return status, nil
}
