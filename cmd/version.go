package cmd

import (
	"fmt"
	"github.com/4everlynn/journal/config"
	"github.com/spf13/cobra"
)

var JournalVersion = config.Version{
	Master: 1,
	Mirror: 0,
	Patch:  5,
}

func buildVersion() string {
	return fmt.Sprintf("%d.%d.%d", JournalVersion.Master, JournalVersion.Mirror, JournalVersion.Patch)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version of journal",
	Long:  `Show version of journal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Journal Command Line Interface by Edward Jobs <diswares@outlook.com>")
		fmt.Printf("version %s built at 2021/07/09\n", buildVersion())
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
