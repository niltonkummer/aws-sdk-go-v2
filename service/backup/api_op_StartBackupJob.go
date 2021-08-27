// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts an on-demand backup job for the specified resource.
func (c *Client) StartBackupJob(ctx context.Context, params *StartBackupJobInput, optFns ...func(*Options)) (*StartBackupJobOutput, error) {
	if params == nil {
		params = &StartBackupJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBackupJob", params, optFns, c.addOperationStartBackupJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBackupJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBackupJobInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	//
	// This member is required.
	IamRoleArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	//
	// This member is required.
	ResourceArn *string

	// Specifies the backup option for a selected resource. This option is only
	// available for Windows Volume Shadow Copy Service (VSS) backup jobs. Valid
	// values: Set to "WindowsVSS":"enabled" to enable the WindowsVSS backup option and
	// create a Windows VSS backup. Set to "WindowsVSS""disabled" to create a regular
	// backup. The WindowsVSS option is not enabled by default.
	BackupOptions map[string]string

	// A value in minutes during which a successfully started backup must complete, or
	// else AWS Backup will cancel the job. This value is optional. This value begins
	// counting down from when the backup was scheduled. It does not add additional
	// time for StartWindowMinutes, or if the backup started later than scheduled.
	CompleteWindowMinutes *int64

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to StartBackupJob. Retrying a successful request with the same
	// idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. Backup will transition and expire backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the “expire
	// after days” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold. Only Amazon EFS file system
	// backups can be transitioned to cold storage.
	Lifecycle *types.Lifecycle

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair.
	RecoveryPointTags map[string]string

	// A value in minutes after a backup is scheduled before a job will be canceled if
	// it doesn't start successfully. This value is optional, and the default is 8
	// hours.
	StartWindowMinutes *int64

	noSmithyDocumentSerde
}

type StartBackupJobOutput struct {

	// Uniquely identifies a request to Backup to back up a resource.
	BackupJobId *string

	// The date and time that a backup job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBackupJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBackupJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBackupJob{}, middleware.After)
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
	if err = addOpStartBackupJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBackupJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBackupJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartBackupJob",
	}
}
