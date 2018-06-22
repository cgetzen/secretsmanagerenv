package main

import (
  "fmt"
  "os"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/secretsmanager"
  "encoding/json"
  "os/exec"
  // "bytes"
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
  path := config.secret_path
  if DEBUG {
    fmt.Println(path)
    for _, x := range command {
      fmt.Println(x)
    }
  }

  sess := session.Must(session.NewSession(&aws.Config{
  	Region: aws.String("us-east-1"),
  }))

  svc := secretsmanager.New(sess)
  input := &secretsmanager.GetSecretValueInput{
      SecretId:     aws.String(path),
  }

  result, err := svc.GetSecretValue(input)

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  secret_bytes := []byte(*result.SecretString)

  if DEBUG {
    fmt.Println(secret_bytes)
  }

  // var f interface{}
  var f map[string]string
  err = json.Unmarshal(secret_bytes, &f)

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  cmd := exec.Command(command[0], command[1:]...)
  cmd.Env = os.Environ()
   cmd.Env = append(os.Environ(), map_to_equal_string(f)...)
  cmd_output, err := cmd.Output()
  if err != nil {
      fmt.Println(err.Error())
     return err
  }
  fmt.Println(string(cmd_output))

  return nil
}
