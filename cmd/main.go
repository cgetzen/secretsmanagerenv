package main

import (
	"os"
  "log"
  "github.com/urfave/cli"
  "errors"
)

func main() {
  var config Config
	app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name:        "secret, s",
      Usage:       "Secrets Manager entry",
      Destination: &config.secret_path,
    },
    cli.StringFlag{
      Name:        "region, r",
      Usage:       "AWS region",
      Destination: &config.region,
    },
  }

  app.Action = func(c *cli.Context) error {
    if len(config.secret_path) == 0 {
      return errors.New("Must specify secret with `-s`")
    }
    if len(c.Args()) == 0 {
      return errors.New("Must run a process with the secret")
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
