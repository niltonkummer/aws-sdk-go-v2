module github.com/aws/aws-sdk-go-v2/service/dynamodb

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.11.2
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.2
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.0.2
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.5.0
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.3.3
	github.com/aws/smithy-go v1.9.0
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../service/internal/endpoint-discovery/
