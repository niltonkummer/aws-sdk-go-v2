// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new custom vocabulary that you can use to modify how Amazon Transcribe
// Medical transcribes your audio file.
func (c *Client) CreateMedicalVocabulary(ctx context.Context, params *CreateMedicalVocabularyInput, optFns ...func(*Options)) (*CreateMedicalVocabularyOutput, error) {
	if params == nil {
		params = &CreateMedicalVocabularyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMedicalVocabulary", params, optFns, c.addOperationCreateMedicalVocabularyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMedicalVocabularyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMedicalVocabularyInput struct {

	// The language code for the language used for the entries in your custom
	// vocabulary. The language code of your custom vocabulary must match the language
	// code of your transcription job. US English (en-US) is the only language code
	// available for Amazon Transcribe Medical.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The location in Amazon S3 of the text file you use to define your custom
	// vocabulary. The URI must be in the same Amazon Web Services Region as the
	// resource that you're calling. Enter information about your VocabularyFileUri in
	// the following format:  https://s3..amazonaws.com///  The following is an example
	// URI for a vocabulary file that is stored in Amazon S3:
	// https://s3.us-east-1.amazonaws.com/AWSDOC-EXAMPLE-BUCKET/vocab.txt For more
	// information about Amazon S3 object names, see Object Keys
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html#object-keys)
	// in the Amazon S3 Developer Guide. For more information about custom
	// vocabularies, see Medical Custom Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-it-works.html#how-vocabulary-med).
	//
	// This member is required.
	VocabularyFileUri *string

	// The name of the custom vocabulary. This case-sensitive name must be unique
	// within an Amazon Web Services account. If you try to create a vocabulary with
	// the same name as a previous vocabulary, you get a ConflictException error.
	//
	// This member is required.
	VocabularyName *string

	// Adds one or more tags, each in the form of a key:value pair, to a new medical
	// vocabulary at the time you create this new vocabulary.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMedicalVocabularyOutput struct {

	// If the VocabularyState field is FAILED, this field contains information about
	// why the job failed.
	FailureReason *string

	// The language code for the entries in your custom vocabulary. US English (en-US)
	// is the only valid language code for Amazon Transcribe Medical.
	LanguageCode types.LanguageCode

	// The date and time that you created the vocabulary.
	LastModifiedTime *time.Time

	// The name of the vocabulary. The name must be unique within an Amazon Web
	// Services account and is case sensitive.
	VocabularyName *string

	// The processing state of your custom vocabulary in Amazon Transcribe Medical. If
	// the state is READY, you can use the vocabulary in a StartMedicalTranscriptionJob
	// request.
	VocabularyState types.VocabularyState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMedicalVocabularyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMedicalVocabulary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMedicalVocabulary{}, middleware.After)
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
	if err = addOpCreateMedicalVocabularyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMedicalVocabulary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMedicalVocabulary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateMedicalVocabulary",
	}
}
