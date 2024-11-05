package cmd

import (
	"fmt"
	"os"

	"github.com/rhermens/timer/config"
	"github.com/spf13/cobra"
)

var (
	cfg     config.Config
	rootCmd = &cobra.Command{
		Use:   "timer",
		Short: "timer is a CLI tool to track time",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cfg = config.DefaultConfig()
	
	if err := os.MkdirAll(cfg.Directory, 0755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
