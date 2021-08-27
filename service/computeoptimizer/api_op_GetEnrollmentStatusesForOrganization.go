// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the Compute Optimizer enrollment (opt-in) status of organization member
// accounts, if your account is an organization management account. To get the
// enrollment status of standalone accounts, use the GetEnrollmentStatus action.
func (c *Client) GetEnrollmentStatusesForOrganization(ctx context.Context, params *GetEnrollmentStatusesForOrganizationInput, optFns ...func(*Options)) (*GetEnrollmentStatusesForOrganizationOutput, error) {
	if params == nil {
		params = &GetEnrollmentStatusesForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEnrollmentStatusesForOrganization", params, optFns, c.addOperationGetEnrollmentStatusesForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEnrollmentStatusesForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEnrollmentStatusesForOrganizationInput struct {

	// An array of objects to specify a filter that returns a more specific list of
	// account enrollment statuses.
	Filters []types.EnrollmentFilter

	// The maximum number of account enrollment statuses to return with a single
	// request. You can specify up to 100 statuses to return with each request. To
	// retrieve the remaining results, make another request with the returned nextToken
	// value.
	MaxResults *int32

	// The token to advance to the next page of account enrollment statuses.
	NextToken *string

	noSmithyDocumentSerde
}

type GetEnrollmentStatusesForOrganizationOutput struct {

	// An array of objects that describe the enrollment statuses of organization member
	// accounts.
	AccountEnrollmentStatuses []types.AccountEnrollmentStatus

	// The token to use to advance to the next page of account enrollment statuses.
	// This value is null when there are no more pages of account enrollment statuses
	// to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEnrollmentStatusesForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEnrollmentStatusesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEnrollmentStatusesForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEnrollmentStatusesForOrganization(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opGetEnrollmentStatusesForOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetEnrollmentStatusesForOrganization",
	}
}
