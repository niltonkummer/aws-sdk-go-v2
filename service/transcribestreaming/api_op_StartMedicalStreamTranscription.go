// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribestreaming

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream/eventstreamapi"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribestreaming/types"
	"github.com/aws/smithy-go/middleware"
	smithysync "github.com/aws/smithy-go/sync"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"sync"
	"time"
)

// Starts a bidirectional HTTP/2 stream where audio is streamed to Amazon
// Transcribe Medical and the transcription results are streamed to your
// application.
func (c *Client) StartMedicalStreamTranscription(ctx context.Context, params *StartMedicalStreamTranscriptionInput, optFns ...func(*Options)) (*StartMedicalStreamTranscriptionOutput, error) {
	if params == nil {
		params = &StartMedicalStreamTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMedicalStreamTranscription", params, optFns, c.addOperationStartMedicalStreamTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMedicalStreamTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMedicalStreamTranscriptionInput struct {

	// Indicates the source language used in the input audio stream. For Amazon
	// Transcribe Medical, this is US English (en-US).
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The encoding used for the input audio.
	//
	// This member is required.
	MediaEncoding types.MediaEncoding

	// The sample rate of the input audio in Hertz.
	//
	// This member is required.
	MediaSampleRateHertz *int32

	// The medical specialty of the clinician or provider.
	//
	// This member is required.
	Specialty types.Specialty

	// The type of input audio. Choose DICTATION for a provider dictating patient
	// notes. Choose CONVERSATION for a dialogue between a patient and one or more
	// medical professionanls.
	//
	// This member is required.
	Type types.Type

	// Set this field to PHI to identify personal health information in the
	// transcription output.
	ContentIdentificationType types.MedicalContentIdentificationType

	// When true, instructs Amazon Transcribe Medical to process each audio channel
	// separately and then merge the transcription output of each channel into a single
	// transcription. Amazon Transcribe Medical also produces a transcription of each
	// item. An item includes the start time, end time, and any alternative
	// transcriptions. You can't set both ShowSpeakerLabel and
	// EnableChannelIdentification in the same request. If you set both, your request
	// returns a BadRequestException.
	EnableChannelIdentification bool

	// The number of channels that are in your audio stream.
	NumberOfChannels *int32

	// Optional. An identifier for the transcription session. If you don't provide a
	// session ID, Amazon Transcribe generates one for you and returns it in the
	// response.
	SessionId *string

	// When true, enables speaker identification in your real-time stream.
	ShowSpeakerLabel bool

	// The name of the medical custom vocabulary to use when processing the real-time
	// stream.
	VocabularyName *string

	noSmithyDocumentSerde
}

type StartMedicalStreamTranscriptionOutput struct {

	// If the value is PHI, indicates that you've configured your stream to identify
	// personal health information.
	ContentIdentificationType types.MedicalContentIdentificationType

	// Shows whether channel identification has been enabled in the stream.
	EnableChannelIdentification bool

	// The language code for the response transcript. For Amazon Transcribe Medical,
	// this is US English (en-US).
	LanguageCode types.LanguageCode

	// The encoding used for the input audio stream.
	MediaEncoding types.MediaEncoding

	// The sample rate of the input audio in Hertz.
	MediaSampleRateHertz *int32

	// The number of channels identified in the stream.
	NumberOfChannels *int32

	// An identifier for the streaming transcription.
	RequestId *string

	// Optional. An identifier for the transcription session. If you don't provide a
	// session ID, Amazon Transcribe generates one for you and returns it in the
	// response.
	SessionId *string

	// Shows whether speaker identification was enabled in the stream.
	ShowSpeakerLabel bool

	// The specialty in the medical domain.
	Specialty types.Specialty

	// The type of audio that was transcribed.
	Type types.Type

	// The name of the vocabulary used when processing the stream.
	VocabularyName *string

	eventStream *StartMedicalStreamTranscriptionEventStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

// GetStream returns the type to interact with the event stream.
func (o *StartMedicalStreamTranscriptionOutput) GetStream() *StartMedicalStreamTranscriptionEventStream {
	return o.eventStream
}

func (c *Client) addOperationStartMedicalStreamTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMedicalStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMedicalStreamTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addEventStreamStartMedicalStreamTranscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddRequireMinimumProtocol(stack, 2, 0); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddStreamingEventsPayload(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = eventstreamapi.AddInitializeStreamWriter(stack); err != nil {
		return err
	}
	if err = addOpStartMedicalStreamTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMedicalStreamTranscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMedicalStreamTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartMedicalStreamTranscription",
	}
}

// StartMedicalStreamTranscriptionEventStream provides the event stream handling for the StartMedicalStreamTranscription operation.
//
// For testing and mocking the event stream this type should be initialized via
// the NewStartMedicalStreamTranscriptionEventStream constructor function. Using the functional options
// to pass in nested mock behavior.
type StartMedicalStreamTranscriptionEventStream struct {
	// AudioStreamWriter is the EventStream writer for the AudioStream events. This
	// value is automatically set by the SDK when the API call is made Use this member
	// when unit testing your code with the SDK to mock out the EventStream Writer.
	//
	// Must not be nil.
	Writer AudioStreamWriter

	// MedicalTranscriptResultStreamReader is the EventStream reader for the
	// MedicalTranscriptResultStream events. This value is automatically set by the SDK
	// when the API call is made Use this member when unit testing your code with the
	// SDK to mock out the EventStream Reader.
	//
	// Must not be nil.
	Reader MedicalTranscriptResultStreamReader

	done      chan struct{}
	closeOnce sync.Once
	err       *smithysync.OnceErr
}

// NewStartMedicalStreamTranscriptionEventStream initializes an StartMedicalStreamTranscriptionEventStream.
// This function should only be used for testing and mocking the StartMedicalStreamTranscriptionEventStream
// stream within your application.
//
// The Writer member must be set before writing events to the stream.
//
// The Reader member must be set before reading events from the stream.
func NewStartMedicalStreamTranscriptionEventStream(optFns ...func(*StartMedicalStreamTranscriptionEventStream)) *StartMedicalStreamTranscriptionEventStream {
	es := &StartMedicalStreamTranscriptionEventStream{
		done: make(chan struct{}),
		err:  smithysync.NewOnceErr(),
	}
	for _, fn := range optFns {
		fn(es)
	}
	return es
}

// Send writes the event to the stream blocking until the event is written.
// Returns an error if the event was not written.
func (es *StartMedicalStreamTranscriptionEventStream) Send(ctx context.Context, event types.AudioStream) error {
	return es.Writer.Send(ctx, event)
}

// Events returns a channel to read events from.
func (es *StartMedicalStreamTranscriptionEventStream) Events() <-chan types.MedicalTranscriptResultStream {
	return es.Reader.Events()
}

// Close closes the stream. This will also cause the stream to be closed.
// Close must be called when done using the stream API. Not calling Close
// may result in resource leaks.
//
// Will close the underlying EventStream writer and reader, and no more events can be
// sent or received.
func (es *StartMedicalStreamTranscriptionEventStream) Close() error {
	es.closeOnce.Do(es.safeClose)
	return es.Err()
}

func (es *StartMedicalStreamTranscriptionEventStream) safeClose() {
	close(es.done)

	t := time.NewTicker(time.Second)
	defer t.Stop()
	writeCloseDone := make(chan error)
	go func() {
		if err := es.Writer.Close(); err != nil {
			es.err.SetError(err)
		}
		close(writeCloseDone)
	}()
	select {
	case <-t.C:
	case <-writeCloseDone:
	}

	es.Reader.Close()
}

// Err returns any error that occurred while reading or writing EventStream Events
// from the service API's response. Returns nil if there were no errors.
func (es *StartMedicalStreamTranscriptionEventStream) Err() error {
	if err := es.err.Err(); err != nil {
		return err
	}

	if err := es.Writer.Err(); err != nil {
		return err
	}

	if err := es.Reader.Err(); err != nil {
		return err
	}

	return nil
}

func (es *StartMedicalStreamTranscriptionEventStream) waitStreamClose() {
	type errorSet interface {
		ErrorSet() <-chan struct{}
	}

	var inputErrCh <-chan struct{}
	if v, ok := es.Writer.(errorSet); ok {
		inputErrCh = v.ErrorSet()
	}

	var outputErrCh <-chan struct{}
	if v, ok := es.Reader.(errorSet); ok {
		outputErrCh = v.ErrorSet()
	}
	var outputClosedCh <-chan struct{}
	if v, ok := es.Reader.(interface{ Closed() <-chan struct{} }); ok {
		outputClosedCh = v.Closed()
	}

	select {
	case <-es.done:
	case <-inputErrCh:
		es.err.SetError(es.Writer.Err())
		es.Close()

	case <-outputErrCh:
		es.err.SetError(es.Reader.Err())
		es.Close()

	case <-outputClosedCh:
		if err := es.Reader.Err(); err != nil {
			es.err.SetError(es.Reader.Err())
		}
		es.Close()

	}
}
