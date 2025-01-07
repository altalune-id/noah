package main

import (
	"github.com/altalune-id/noah/cdk/config"
	"github.com/altalune-id/noah/cdk/lambda"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppProps struct {
	awscdk.StackProps
	Stage string
}

func NewApp(scope constructs.Construct, props *AppProps) awscdk.Stack {
	stackID := jsii.String(config.AppName)
	stack := awscdk.NewStack(scope, stackID, &props.StackProps)

	lambda.NewRestapiLambda(stack, props.Stage)

	return stack
}
