package main

import (
  "fmt"
)

const (

)

var (

)

// NewCLI creates a new command line interface with the given streams.
func RunScript(path string, command []string) error {
  fmt.Println(path)
  for _, x := range command {
    fmt.Println(x)
  }

  return nil
}
