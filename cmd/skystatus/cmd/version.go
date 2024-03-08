package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print SkyStatus version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
