package cmd

import (
	"diswares.com.journal/config"
	"diswares.com.journal/global"
	"diswares.com.journal/support"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

// stylishCmd represents the stylish command
var stylishCmd = &cobra.Command{
	Use:   "stylish",
	Short: "Generate formatted report content based on configuration",
	Long:  ``,
	Run:   generate,
}

// generate daily report
func generate(cmd *cobra.Command, args []string) {
	cfg := GetConfig()
	println(time.Now().Format("20060102") + " 日报")
	for service, git := range cfg.Git {
		all := make([]support.GitLog, 0)
		repo := strings.Join([]string{git.Path, ".git", "logs", "HEAD"}, global.Separator)
		git.Path = repo
		_, err := os.Stat(repo)
		if err == nil {
			timeStr := time.Now().Format("2006-01-02")
			start, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
			all = append(all, ReportFromGit(git, start.Unix(), start.AddDate(0, 0, 1).Unix()-1)...)
			if len(all) > 0 {
				// inject real path
				git.Path = repo
				println(git.Name)
				formatter := new(support.GitLogDailyFormatter)
				println(formatter.Format(all))
			}
		} else if os.IsNotExist(err) {
			color.Danger.Printf(".git is missing in repository %s\n", service)
		}
	}
}

func ReportFromGit(git config.Git, start int64, end int64) []support.GitLog {
	parser := new(support.GitLogParser)
	return parser.Parse(git.Path, start, end)
}

func init() {
	rootCmd.AddCommand(stylishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stylishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	stylishCmd.Flags().BoolP("day", "d", true, "output as work report by day")
}
