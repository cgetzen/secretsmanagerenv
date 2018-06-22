package main

import (
  "os"
  "errors"
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
