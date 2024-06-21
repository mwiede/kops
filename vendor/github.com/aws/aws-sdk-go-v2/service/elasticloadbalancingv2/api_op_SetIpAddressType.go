// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the type of IP addresses used by the subnets of the specified load
// balancer.
func (c *Client) SetIpAddressType(ctx context.Context, params *SetIpAddressTypeInput, optFns ...func(*Options)) (*SetIpAddressTypeOutput, error) {
	if params == nil {
		params = &SetIpAddressTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIpAddressType", params, optFns, c.addOperationSetIpAddressTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIpAddressTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetIpAddressTypeInput struct {

	// Note: Internal load balancers must use the ipv4 IP address type.
	//
	// [Application Load Balancers] The IP address type. The possible values are ipv4
	// (for only IPv4 addresses), dualstack (for IPv4 and IPv6 addresses), and
	// dualstack-without-public-ipv4 (for IPv6 only public addresses, with private IPv4
	// and IPv6 addresses).
	//
	// [Network Load Balancers] The IP address type. The possible values are ipv4 (for
	// only IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses). You can’t
	// specify dualstack for a load balancer with a UDP or TCP_UDP listener.
	//
	// [Gateway Load Balancers] The IP address type. The possible values are ipv4 (for
	// only IPv4 addresses) and dualstack (for IPv4 and IPv6 addresses).
	//
	// This member is required.
	IpAddressType types.IpAddressType

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string

	noSmithyDocumentSerde
}

type SetIpAddressTypeOutput struct {

	// The IP address type.
	IpAddressType types.IpAddressType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetIpAddressTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetIpAddressType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetIpAddressType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetIpAddressType"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpSetIpAddressTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIpAddressType(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetIpAddressType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetIpAddressType",
	}
}
