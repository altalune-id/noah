package lambda

import (
	"github.com/altalune-id/noah/cdk/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func NewWebsocketLambda(stack awscdk.Stack, cfg *config.Config) awslambda.Function {
	functionName := config.AppName + "-websocket-" + cfg.StageLowerCase()

	function := awslambda.NewFunction(stack, jsii.String("WebsocketLambda"), &awslambda.FunctionProps{
		FunctionName: jsii.String(functionName),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./bin/websocket.zip"), &awss3assets.AssetOptions{}),
	})

	awscdk.NewCfnOutput(stack, jsii.String("WebsocketLambdaOutputArn"), &awscdk.CfnOutputProps{
		Value:      function.FunctionArn(),
		ExportName: jsii.String(functionName + "-arn"),
	})

	return function
}
