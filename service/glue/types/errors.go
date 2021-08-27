// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Access to a resource was denied.
type AccessDeniedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource to be created or added already exists.
type AlreadyExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string             { return "AlreadyExistsException" }
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Two processes are trying to modify a resource simultaneously.
type ConcurrentModificationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Too many jobs are being run concurrently.
type ConcurrentRunsExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentRunsExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentRunsExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentRunsExceededException) ErrorCode() string {
	return "ConcurrentRunsExceededException"
}
func (e *ConcurrentRunsExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A specified condition was not satisfied.
type ConditionCheckFailureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConditionCheckFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConditionCheckFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConditionCheckFailureException) ErrorCode() string             { return "ConditionCheckFailureException" }
func (e *ConditionCheckFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The CreatePartitions API was called on a table that has indexes enabled.
type ConflictException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string             { return "ConflictException" }
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified crawler is not running.
type CrawlerNotRunningException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CrawlerNotRunningException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CrawlerNotRunningException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CrawlerNotRunningException) ErrorCode() string             { return "CrawlerNotRunningException" }
func (e *CrawlerNotRunningException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation cannot be performed because the crawler is already running.
type CrawlerRunningException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CrawlerRunningException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CrawlerRunningException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CrawlerRunningException) ErrorCode() string             { return "CrawlerRunningException" }
func (e *CrawlerRunningException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified crawler is stopping.
type CrawlerStoppingException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CrawlerStoppingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CrawlerStoppingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CrawlerStoppingException) ErrorCode() string             { return "CrawlerStoppingException" }
func (e *CrawlerStoppingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A specified entity does not exist
type EntityNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *EntityNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityNotFoundException) ErrorCode() string             { return "EntityNotFoundException" }
func (e *EntityNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An encryption operation failed.
type GlueEncryptionException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *GlueEncryptionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *GlueEncryptionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *GlueEncryptionException) ErrorCode() string             { return "GlueEncryptionException" }
func (e *GlueEncryptionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The same unique identifier was associated with two different records.
type IdempotentParameterMismatchException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	return "IdempotentParameterMismatchException"
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

type IllegalBlueprintStateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *IllegalBlueprintStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IllegalBlueprintStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IllegalBlueprintStateException) ErrorCode() string             { return "IllegalBlueprintStateException" }
func (e *IllegalBlueprintStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The workflow is in an invalid state to perform a requested operation.
type IllegalWorkflowStateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *IllegalWorkflowStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IllegalWorkflowStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IllegalWorkflowStateException) ErrorCode() string             { return "IllegalWorkflowStateException" }
func (e *IllegalWorkflowStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal service error occurred.
type InternalServiceException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceException) ErrorCode() string             { return "InternalServiceException" }
func (e *InternalServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The input provided was not valid.
type InvalidInputException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The machine learning transform is not ready to run.
type MLTransformNotReadyException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MLTransformNotReadyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MLTransformNotReadyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MLTransformNotReadyException) ErrorCode() string             { return "MLTransformNotReadyException" }
func (e *MLTransformNotReadyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is no applicable schedule.
type NoScheduleException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *NoScheduleException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoScheduleException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoScheduleException) ErrorCode() string             { return "NoScheduleException" }
func (e *NoScheduleException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation timed out.
type OperationTimeoutException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *OperationTimeoutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationTimeoutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationTimeoutException) ErrorCode() string             { return "OperationTimeoutException" }
func (e *OperationTimeoutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource numerical limit was exceeded.
type ResourceNumberLimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNumberLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNumberLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNumberLimitExceededException) ErrorCode() string {
	return "ResourceNumberLimitExceededException"
}
func (e *ResourceNumberLimitExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified scheduler is not running.
type SchedulerNotRunningException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SchedulerNotRunningException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchedulerNotRunningException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchedulerNotRunningException) ErrorCode() string             { return "SchedulerNotRunningException" }
func (e *SchedulerNotRunningException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified scheduler is already running.
type SchedulerRunningException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SchedulerRunningException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchedulerRunningException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchedulerRunningException) ErrorCode() string             { return "SchedulerRunningException" }
func (e *SchedulerRunningException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified scheduler is transitioning.
type SchedulerTransitioningException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *SchedulerTransitioningException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SchedulerTransitioningException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SchedulerTransitioningException) ErrorCode() string {
	return "SchedulerTransitioningException"
}
func (e *SchedulerTransitioningException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A value could not be validated.
type ValidationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string             { return "ValidationException" }
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a version conflict.
type VersionMismatchException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *VersionMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *VersionMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *VersionMismatchException) ErrorCode() string             { return "VersionMismatchException" }
func (e *VersionMismatchException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
