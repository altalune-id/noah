package lambda

import (
	"github.com/altalune-id/noah/cdk/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func NewRestapiLambda(stack awscdk.Stack, stage string) awslambda.Function {
	functionID := jsii.String("restapi")
	functionArn := jsii.String("restapi-arn")

	function := awslambda.NewFunction(stack, functionID, &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./bin/restapi.zip"), &awss3assets.AssetOptions{}),
	})

	awscdk.NewCfnOutput(stack, functionArn, &awscdk.CfnOutputProps{
		Value:      function.FunctionArn(),
		ExportName: jsii.String(config.ExportedLambdaARN("restapi", stage)),
	})

	return function
}
