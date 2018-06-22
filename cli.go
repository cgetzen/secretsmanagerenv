package main

import (
  "fmt"
  "os"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/secretsmanager"
  "encoding/json"
  "os/exec"
  "io"
  // "bytes"
  // "reflect"
)

const (

)

var (
  DEBUG = false
)

func map_to_equal_string(m map[string]string) []string {
    var ret []string
    for key, value := range m {
        keyval := fmt.Sprintf("%s=%s", key, value)
        ret = append(ret, keyval)
    }
    return ret
}

func RunScript(config Config, command []string) error {
  var sess *session.Session
  var secrets_json map[string]string

  if region, err := config.Region(); err == nil {
    sess = session.Must(session.NewSession(&aws.Config{
    	Region: aws.String(region),
    }))
  } else {
    sess = session.Must(session.NewSession())
  }

  svc := secretsmanager.New(sess)

  input := &secretsmanager.GetSecretValueInput{
      SecretId:     aws.String(config.secret_path),
  }

  result, err := svc.GetSecretValue(input)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  err = json.Unmarshal([]byte(*result.SecretString), &secrets_json)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  cmd := exec.Command(command[0], command[1:]...)
  cmd.Env = append(os.Environ(), map_to_equal_string(secrets_json)...)

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err = cmd.Run()

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  return nil
}
