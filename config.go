package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const defaultRegion = "us-east-1"

// awsConfig creates *aws.Config object from the fields.
func awsConfig() *aws.Config {
	cred := awsCredentials()
	awsConf := &aws.Config{
		Credentials: cred,
		Region:      stringPtr(getRegion()),
	}

	ep := getEndpoint()
	if ep != "" {
		awsConf.Endpoint = &ep
	}

	return awsConf
}

func awsCredentials() *credentials.Credentials {
	// from env
	cred := credentials.NewEnvCredentials()
	_, err := cred.Get()
	if err == nil {
		return cred
	}

	// from local file
	return credentials.NewSharedCredentials("", "")
}

func getRegion() string {
	reg := envRegion()
	if reg != "" {
		return reg
	}
	return defaultRegion
}

func getEndpoint() string {
	ep := envEndpoint()
	if ep != "" {
		return ep
	}
	return ""
}

// envRegion get aws region from env params
func envRegion() string {
	return os.Getenv("AWS_REGION")
}

// envEndpoint get aws endpoint from env params
func envEndpoint() string {
	return os.Getenv("AWS_SQS_ENDPOINT")
}
