// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom language model. Use Amazon S3 prefixes to provide the
// location of your input files. The time it takes to create your model depends on
// the size of your training data.
func (c *Client) CreateLanguageModel(ctx context.Context, params *CreateLanguageModelInput, optFns ...func(*Options)) (*CreateLanguageModelOutput, error) {
	if params == nil {
		params = &CreateLanguageModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLanguageModel", params, optFns, c.addOperationCreateLanguageModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLanguageModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLanguageModelInput struct {

	// The Amazon Transcribe standard language model, or base model used to create your
	// custom language model. If you want to use your custom language model to
	// transcribe audio with a sample rate of 16,000 Hz or greater, choose Wideband. If
	// you want to use your custom language model to transcribe audio with a sample
	// rate that is less than 16,000 Hz, choose Narrowband.
	//
	// This member is required.
	BaseModelName types.BaseModelName

	// Contains the data access role and the Amazon S3 prefixes to read the required
	// input files to create a custom language model.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input text you're using to train your custom language model.
	//
	// This member is required.
	LanguageCode types.CLMLanguageCode

	// The name you choose for your custom language model when you create it.
	//
	// This member is required.
	ModelName *string

	// Adds one or more tags, each in the form of a key:value pair, to a new language
	// model at the time you create this new model.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateLanguageModelOutput struct {

	// The Amazon Transcribe standard language model, or base model you've used to
	// create a custom language model.
	BaseModelName types.BaseModelName

	// The data access role and Amazon S3 prefixes you've chosen to create your custom
	// language model.
	InputDataConfig *types.InputDataConfig

	// The language code of the text you've used to create a custom language model.
	LanguageCode types.CLMLanguageCode

	// The name you've chosen for your custom language model.
	ModelName *string

	// The status of the custom language model. When the status is COMPLETED the model
	// is ready to use.
	ModelStatus types.ModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLanguageModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLanguageModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLanguageModel{}, middleware.After)
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
	if err = addOpCreateLanguageModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLanguageModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLanguageModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateLanguageModel",
	}
}
