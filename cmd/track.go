package cmd

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:  "track <DESCRIPTION>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(path.Join(cfg.Directory, fmt.Sprintf("%s.out", time.Now().Format("2006-01-02"))), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		defer file.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if _, err := file.Write([]byte(fmt.Sprintf(cfg.Format, time.Now().Format("15:04"), args[0]))); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}
