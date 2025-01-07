package config

import (
	"os"
	"strings"
)

const AppName = "noah"

var (
	AwsAccountID = os.Getenv("AWS_ACCOUNT_ID")
	AwsRegion    = os.Getenv("AWS_REGION")
)

func stageName(stage string) string {
	return strings.ToLower(stage)
}

func ExportedLambdaARN(function string, stage string) string {
	return AppName + "-" + function + "-" + stageName(stage) + "-arn"
}
