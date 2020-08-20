// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of ContributorInsightsSummary for a table and all its global
// secondary indexes.
func (c *Client) ListContributorInsights(ctx context.Context, params *ListContributorInsightsInput, optFns ...func(*Options)) (*ListContributorInsightsOutput, error) {
	stack := middleware.NewStack("ListContributorInsights", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListContributorInsightsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListContributorInsights(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "ListContributorInsights",
			Err:           err,
		}
	}
	out := result.(*ListContributorInsightsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContributorInsightsInput struct {
	// Maximum number of results to return per page.
	MaxResults *int32
	// A token to for the desired page, if there is one.
	NextToken *string
	// The name of the table.
	TableName *string
}

type ListContributorInsightsOutput struct {
	// A list of ContributorInsightsSummary.
	ContributorInsightsSummaries []*types.ContributorInsightsSummary
	// A token to go to the next page if there is one.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListContributorInsightsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListContributorInsights{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListContributorInsights{}, middleware.After)
}

func newServiceMetadataMiddleware_opListContributorInsights(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "ListContributorInsights",
	}
}
