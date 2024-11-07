package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use: "current <DESCRIPTION>",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(path.Join(cfg.Directory, fmt.Sprintf("%s.out", time.Now().Format("2006-01-02"))), os.O_RDONLY, 0644)
		defer file.Close()

		if err != nil {
			if os.IsNotExist(err) {
				return
			}
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}

		buffer := make([]byte, 32)
		if _, err := file.Seek(max(stat.Size()-int64(len(buffer)), 0), 0); err != nil {
			panic(err)
		}
		if _, err := file.Read(buffer); err != nil {
			if err.Error() != io.EOF.Error() {
				panic(err)
			}
		}

		lines := strings.Split(string(buffer), "\n")
		if len(lines) > 1 {
			fmt.Println(lines[len(lines)-1])
		}
	},
}
