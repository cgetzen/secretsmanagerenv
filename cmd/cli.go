package main

import (
  "fmt"
  "os"
  "os/exec"
  "encoding/json"
  "github.com/aws/aws-sdk-go/service/secretsmanager"
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
  var secrets map[string]string

  // Grab env vars from Secrets Manager
  session := config.SmSession()
  svc := secretsmanager.New(session)
  input := config.SmInput()
  result, err := svc.GetSecretValue(input)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  // Massage resulting string
  err = json.Unmarshal([]byte(*result.SecretString), &secrets)
  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  // Run command in a subprocess
  cmd := exec.Command(command[0], command[1:]...)
  cmd.Env = append(os.Environ(), map_to_equal_string(secrets)...)
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
