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

type Config struct {
	StageName string
	NoahApiID string
}

func (c *Config) StageLowerCase() string {
	return strings.ToLower(c.StageName)
}
