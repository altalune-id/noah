package awscli

import (
	"github.com/altalune-id/noah/config"
	"github.com/aws/aws-sdk-go/aws"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
)

type ApiGatewayClient struct {
	session *awssession.Session
	apigw   *apigatewaymanagementapi.ApiGatewayManagementApi
}

func NewApiGatewayClient(cfg *config.Config) *ApiGatewayClient {
	return &ApiGatewayClient{
		session: Session,
		apigw: apigatewaymanagementapi.New(
			Session,
			aws.NewConfig().WithEndpoint(cfg.WebSocket.Endpoint),
		),
	}
}

func (c *ApiGatewayClient) PostToConnection(input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	return c.apigw.PostToConnection(input)
}
