package handler

import (
  "github.com/cgetzen/secretsmanagerenv/pkg/aws"
  "os/exec"
  "fmt"
  "os"
)

func RunCommandWithSecret(secrets []string, region string, args []string) error {
  data, err := aws.GetSecretData(secrets[0], region)
  if err != nil {
    return err
  }

  env := mapToEnv(data)

  cmd := exec.Command(args[0], args[1:]...)
  cmd.Env = append(os.Environ(), env...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func mapToEnv(m map[string]string) []string {
    var ret []string
    for key, value := range m {
        keyval := fmt.Sprintf("%s=%s", key, value)
        ret = append(ret, keyval)
    }
    return ret
}
