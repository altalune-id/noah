package awscli

import (
	"github.com/altalune-id/noah/config"
	awssdk "github.com/aws/aws-sdk-go/aws"
	awssession "github.com/aws/aws-sdk-go/aws/session"
)

var Session *awssession.Session

func init() {
	cfg, _ := config.LoadConfig("config.yaml")
	Session = awssession.Must(awssession.NewSession(&awssdk.Config{
		Region: awssdk.String(cfg.AWS.Region),
	}))
}
