package main

import (
	"github.com/altalune-id/noah/cdk/config"
	"github.com/altalune-id/noah/cdk/lambda"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewApp(scope constructs.Construct, props *awscdk.StackProps, cfg *config.Config) awscdk.Stack {
	stackID := jsii.String(config.AppName)
	stack := awscdk.NewStack(scope, stackID, props)

	_ = lambda.NewRestapiLambda(stack, cfg)

	return stack
}
