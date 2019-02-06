package handler

import (
	"fmt"
	"github.com/cgetzen/secretsmanagerenv/pkg/aws"
	"os"
	"os/exec"
)

func RunCommandWithSecret(secrets []string, region string, args []string) error {
	var env []string

	for _, secret := range secrets {
		data, err := aws.GetSecretData(secret, region)
		if err != nil {
			return err
		}
		for _, pair := range mapToEnv(data) {
			env = append(env, pair)
		}
	}

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
