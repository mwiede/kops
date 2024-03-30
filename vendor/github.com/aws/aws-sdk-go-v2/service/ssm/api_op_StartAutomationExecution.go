// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates execution of an Automation runbook.
func (c *Client) StartAutomationExecution(ctx context.Context, params *StartAutomationExecutionInput, optFns ...func(*Options)) (*StartAutomationExecutionOutput, error) {
	if params == nil {
		params = &StartAutomationExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAutomationExecution", params, optFns, c.addOperationStartAutomationExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAutomationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAutomationExecutionInput struct {

	// The name of the SSM document to run. This can be a public document or a custom
	// document. To run a shared document belonging to another account, specify the
	// document ARN. For more information about how to use shared documents, see
	// Sharing SSM documents (https://docs.aws.amazon.com/systems-manager/latest/userguide/documents-ssm-sharing.html)
	// in the Amazon Web Services Systems Manager User Guide.
	//
	// This member is required.
	DocumentName *string

	// The CloudWatch alarm you want to apply to your automation.
	AlarmConfiguration *types.AlarmConfiguration

	// User-provided idempotency token. The token must be unique, is case insensitive,
	// enforces the UUID format, and can't be reused.
	ClientToken *string

	// The version of the Automation runbook to use for this execution.
	DocumentVersion *string

	// The maximum number of targets allowed to run this task in parallel. You can
	// specify a number, such as 10, or a percentage, such as 10%. The default value is
	// 10 .
	MaxConcurrency *string

	// The number of errors that are allowed before the system stops running the
	// automation on additional targets. You can specify either an absolute number of
	// errors, for example 10, or a percentage of the target set, for example 10%. If
	// you specify 3, for example, the system stops running the automation when the
	// fourth error is received. If you specify 0, then the system stops running the
	// automation on additional targets after the first error result is returned. If
	// you run an automation on 50 resources and set max-errors to 10%, then the system
	// stops running the automation on additional targets when the sixth error is
	// received. Executions that are already running an automation when max-errors is
	// reached are allowed to complete, but some of these executions may fail as well.
	// If you need to ensure that there won't be more than max-errors failed
	// executions, set max-concurrency to 1 so the executions proceed one at a time.
	MaxErrors *string

	// The execution mode of the automation. Valid modes include the following: Auto
	// and Interactive. The default mode is Auto.
	Mode types.ExecutionMode

	// A key-value map of execution parameters, which match the declared parameters in
	// the Automation runbook.
	Parameters map[string][]string

	// Optional metadata that you assign to a resource. You can specify a maximum of
	// five tags for an automation. Tags enable you to categorize a resource in
	// different ways, such as by purpose, owner, or environment. For example, you
	// might want to tag an automation to identify an environment or operating system.
	// In this case, you could specify the following key-value pairs:
	//   - Key=environment,Value=test
	//   - Key=OS,Value=Windows
	// To add tags to an existing automation, use the AddTagsToResource operation.
	Tags []types.Tag

	// A location is a combination of Amazon Web Services Regions and/or Amazon Web
	// Services accounts where you want to run the automation. Use this operation to
	// start an automation in multiple Amazon Web Services Regions and multiple Amazon
	// Web Services accounts. For more information, see Running Automation workflows
	// in multiple Amazon Web Services Regions and Amazon Web Services accounts (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-automation-multiple-accounts-and-regions.html)
	// in the Amazon Web Services Systems Manager User Guide.
	TargetLocations []types.TargetLocation

	// A key-value mapping of document parameters to target resources. Both Targets
	// and TargetMaps can't be specified together.
	TargetMaps []map[string][]string

	// The name of the parameter used as the target resource for the rate-controlled
	// execution. Required if you specify targets.
	TargetParameterName *string

	// A key-value mapping to target resources. Required if you specify
	// TargetParameterName.
	Targets []types.Target

	noSmithyDocumentSerde
}

type StartAutomationExecutionOutput struct {

	// The unique ID of a newly scheduled automation execution.
	AutomationExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAutomationExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartAutomationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartAutomationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartAutomationExecution"); err != nil {
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
	if err = addOpStartAutomationExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAutomationExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartAutomationExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartAutomationExecution",
	}
}
