// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates the copy of an AMI. You can copy an AMI from one Region to another,
// or from a Region to an Outpost. You can't copy an AMI from an Outpost to a
// Region, from one Outpost to another, or within the same Outpost. To copy an AMI
// to another partition, see [CreateStoreImageTask].
//
// To copy an AMI from one Region to another, specify the source Region using the
// SourceRegion parameter, and specify the destination Region using its endpoint.
// Copies of encrypted backing snapshots for the AMI are encrypted. Copies of
// unencrypted backing snapshots remain unencrypted, unless you set Encrypted
// during the copy operation. You cannot create an unencrypted copy of an encrypted
// backing snapshot.
//
// To copy an AMI from a Region to an Outpost, specify the source Region using the
// SourceRegion parameter, and specify the ARN of the destination Outpost using
// DestinationOutpostArn. Backing snapshots copied to an Outpost are encrypted by
// default using the default encryption key for the Region, or a different key that
// you specify in the request using KmsKeyId. Outposts do not support unencrypted
// snapshots. For more information, [Amazon EBS local snapshots on Outposts]in the Amazon EBS User Guide.
//
// For more information about the prerequisites and limits when copying an AMI,
// see [Copy an AMI]in the Amazon EC2 User Guide.
//
// [CreateStoreImageTask]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateStoreImageTask.html
// [Amazon EBS local snapshots on Outposts]: https://docs.aws.amazon.com/ebs/latest/userguide/snapshots-outposts.html#ami
// [Copy an AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/CopyingAMIs.html
func (c *Client) CopyImage(ctx context.Context, params *CopyImageInput, optFns ...func(*Options)) (*CopyImageOutput, error) {
	if params == nil {
		params = &CopyImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyImage", params, optFns, c.addOperationCopyImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CopyImage.
type CopyImageInput struct {

	// The name of the new AMI in the destination Region.
	//
	// This member is required.
	Name *string

	// The ID of the AMI to copy.
	//
	// This member is required.
	SourceImageId *string

	// The name of the Region that contains the AMI to copy.
	//
	// This member is required.
	SourceRegion *string

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see [Ensuring idempotency]in the Amazon EC2 API Reference.
	//
	// [Ensuring idempotency]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html
	ClientToken *string

	// Indicates whether to include your user-defined AMI tags when copying the AMI.
	//
	// The following tags will not be copied:
	//
	//   - System tags (prefixed with aws: )
	//
	//   - For public and shared AMIs, user-defined tags that are attached by other
	//   Amazon Web Services accounts
	//
	// Default: Your user-defined AMI tags are not copied.
	CopyImageTags *bool

	// A description for the new AMI in the destination Region.
	Description *string

	// The Amazon Resource Name (ARN) of the Outpost to which to copy the AMI. Only
	// specify this parameter when copying an AMI from an Amazon Web Services Region to
	// an Outpost. The AMI must be in the Region of the destination Outpost. You cannot
	// copy an AMI from an Outpost to a Region, from one Outpost to another, or within
	// the same Outpost.
	//
	// For more information, see [Copy AMIs from an Amazon Web Services Region to an Outpost] in the Amazon EBS User Guide.
	//
	// [Copy AMIs from an Amazon Web Services Region to an Outpost]: https://docs.aws.amazon.com/ebs/latest/userguide/snapshots-outposts.html#copy-amis
	DestinationOutpostArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// Specifies whether the destination snapshots of the copied image should be
	// encrypted. You can encrypt a copy of an unencrypted snapshot, but you cannot
	// create an unencrypted copy of an encrypted snapshot. The default KMS key for
	// Amazon EBS is used unless you specify a non-default Key Management Service (KMS)
	// KMS key using KmsKeyId . For more information, see [Amazon EBS encryption] in the Amazon EBS User
	// Guide.
	//
	// [Amazon EBS encryption]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption.html
	Encrypted *bool

	// The identifier of the symmetric Key Management Service (KMS) KMS key to use
	// when creating encrypted volumes. If this parameter is not specified, your Amazon
	// Web Services managed KMS key for Amazon EBS is used. If you specify a KMS key,
	// you must also set the encrypted state to true .
	//
	// You can specify a KMS key using any of the following:
	//
	//   - Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//   - Key alias. For example, alias/ExampleAlias.
	//
	//   - Key ARN. For example,
	//   arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//   - Alias ARN. For example,
	//   arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// Amazon Web Services authenticates the KMS key asynchronously. Therefore, if you
	// specify an identifier that is not valid, the action can appear to complete, but
	// eventually fails.
	//
	// The specified KMS key must exist in the destination Region.
	//
	// Amazon EBS does not support asymmetric KMS keys.
	KmsKeyId *string

	// The tags to apply to the new AMI and new snapshots. You can tag the AMI, the
	// snapshots, or both.
	//
	//   - To tag the new AMI, the value for ResourceType must be image .
	//
	//   - To tag the new snapshots, the value for ResourceType must be snapshot . The
	//   same tag is applied to all the new snapshots.
	//
	// If you specify other values for ResourceType , the request fails.
	//
	// To tag an AMI or snapshot after it has been created, see [CreateTags].
	//
	// [CreateTags]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTags.html
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

// Contains the output of CopyImage.
type CopyImageOutput struct {

	// The ID of the new AMI.
	ImageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCopyImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCopyImage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CopyImage"); err != nil {
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
	if err = addOpCopyImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyImage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopyImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CopyImage",
	}
}
