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

func printOutput(w io.Writer, r io.Reader) error {
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			_, err = w.Write(buf[:n])
		}
    if err == io.EOF {
      return nil
    } else if err != nil {
			return err
		}
	}
	// never reached
	panic(true)
	return nil
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
  cmd.Env = append(os.Environ(), map_to_equal_string(f)...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	go func() {
		 printOutput(os.Stdout, stdout)
	}()

	go func() {
		printOutput(os.Stderr, stderr)
	}()

  err = cmd.Wait()

  if err != nil {
    fmt.Println(err.Error())
    return err
  }

  return nil
}
