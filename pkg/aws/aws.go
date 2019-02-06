package aws

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/secretsmanager"
  "encoding/json"
  "os"
  "fmt"
)

func GetSecretData(name, region string) (map[string]string, error) {
  var secrets map[string]string
  // Grab env vars from Secrets Manager
  session := getSession(region)
  svc := secretsmanager.New(session)
  input := &secretsmanager.GetSecretValueInput{
    SecretId: aws.String(name),
  }

  result, err := svc.GetSecretValue(input)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal([]byte(*result.SecretString), &secrets)
  if err != nil {
    return nil, fmt.Errorf("%s is not a key-pair secret", name)
  }

  return secrets, nil
}

func getSession(region string) *session.Session {
  if len(os.Getenv("AWS_SDK_LOAD_CONFIG")) == 0 {
    os.Setenv("AWS_SDK_LOAD_CONFIG", "TRUE")
    defer os.Unsetenv("AWS_SDK_LOAD_CONFIG")
  }

  if len(region) > 0 {
    return session.Must(session.NewSession(&aws.Config{
      Region: aws.String(region),
    }))
  }
  return session.Must(session.NewSession())
}
