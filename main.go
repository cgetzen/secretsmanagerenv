package main

import (
	"os"
  "log"
  "github.com/urfave/cli"
  "errors"
  "fmt"
)

func main() {
  // var secret_path string
  // var region string
  var config Config
  var secrets *cli.StringSlice

	app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringSliceFlag{
      Name:        "secret, s",
      Usage:       "Secrets Manager entry",
      Value:       secrets,
    },
    cli.StringFlag{
      Name:        "region, r",
      Usage:       "AWS region",
      Destination: &config.region,
    },
  }

  fmt.Println(secrets)

  app.Action = func(c *cli.Context) error {
    if len(config.secret_paths) == 0 {
      return errors.New("Must specify secret with `-s`")
    }
    // c.Args() contains [script, to, run]
    RunScript(config, c.Args())
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
