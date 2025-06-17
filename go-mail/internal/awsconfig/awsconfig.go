package awsconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// NewAWSConfigはregionとendpointを受け取り、AWS configを返します。
func NewAWSConfig(ctx context.Context, region, endpoint string) (aws.Config, error) {
	return config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithBaseEndpoint(endpoint),
	)
}

// NewDefaultAWSConfigはデフォルトのregionとendpointを使用してAWS configを返します。
func NewDefaultAWSConfig(ctx context.Context) (aws.Config, error) {
	return NewAWSConfig(ctx, "ap-northeast-1", "http://localhost:8005")
}
