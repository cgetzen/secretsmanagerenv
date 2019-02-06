package cmd

import (
	"errors"
	"fmt"
	"github.com/cgetzen/secretsmanagerenv/cmd/handler"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	secrets []string
	region  string
)

var rootCmd = &cobra.Command{
	Use:   "smenv",
	Short: "B",
	Long:  "C",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires at least one arg")
		}
		if len(secrets) == 0 {
			return errors.New("Must specify secret with `-s`")
		}
		return nil
	},
	Run: func(_ *cobra.Command, args []string) {
		if err := handler.RunCommandWithSecret(secrets, region, parse(args)); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parse(args []string) []string {
	var ret []string
	for _, arg := range args {
		for _, a := range strings.Split(arg, " ") {
			ret = append(ret, a)
		}
	}
	return ret
}

func init() {
	rootCmd.PersistentFlags().StringSliceVarP(&secrets, "secret", "s", []string{}, "name of secret")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "region")
}
