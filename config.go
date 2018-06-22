package main

import (
  "os"
  "errors"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/secretsmanager"
)

type Config struct {
  region string
  secret_path string
}

func (c *Config) Region() (string, error) {
  if len(c.region) > 0 {
    return c.region, nil
  }

  if envVar := os.Getenv("AWS_DEFAULT_REGION"); len(envVar) > 0 {
    return envVar, nil
  }

  return "", errors.New("Cannot find region")
}

func (c *Config) SmInput() *secretsmanager.GetSecretValueInput {
  return &secretsmanager.GetSecretValueInput{
      SecretId:     aws.String(c.secret_path),
  }
}

func (c *Config) SmSession() *session.Session {
  if region, err := c.Region(); err == nil {
    return session.Must(session.NewSession(&aws.Config{
      Region: aws.String(region),
    }))
  }
  return session.Must(session.NewSession())
}
