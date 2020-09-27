package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of journal",
	Long:  `Show version of journal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Journal Command Line Interface by Edward Jobs <diswares@outlook.com>")
		fmt.Println("version 1.0.0 alpha built at 2020/09/25")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
