package lambda

import (
	"github.com/altalune-id/noah/cdk/config"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func NewRestapiLambda(stack awscdk.Stack, cfg *config.Config) awslambda.Function {
	functionName := config.AppName + "-restapi-" + cfg.StageLowerCase()

	function := awslambda.NewFunction(stack, jsii.String("RestapiLambda"), &awslambda.FunctionProps{
		FunctionName: jsii.String(functionName),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		MemorySize:   jsii.Number(128),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./bin/restapi.zip"), &awss3assets.AssetOptions{}),
	})

	if cfg.NoahApiID != "" {
		apiRef := awsapigatewayv2.HttpApi_FromHttpApiAttributes(stack, jsii.String("NoahApiRef"), &awsapigatewayv2.HttpApiAttributes{
			HttpApiId: jsii.String(cfg.NoahApiID),
		})

		integration := awsapigatewayv2integrations.NewHttpLambdaIntegration(
			jsii.String("RestapiLambdaIntegration"),
			function,
			&awsapigatewayv2integrations.HttpLambdaIntegrationProps{
				Timeout: awscdk.Duration_Seconds(jsii.Number(15)),
			},
		)

		if api, ok := apiRef.(awsapigatewayv2.HttpApi); ok {
			api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
				Integration: integration,
				Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_ANY},
				Path:        jsii.String("/{proxy+}"),
			})
		}
	}

	awscdk.NewCfnOutput(stack, jsii.String("RestapiLambdaOutputArn"), &awscdk.CfnOutputProps{
		Value:      function.FunctionArn(),
		ExportName: jsii.String(functionName + "-arn"),
	})

	return function
}
