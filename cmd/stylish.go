package cmd

import (
	"github.com/4everlynn/journal/config"
	"github.com/4everlynn/journal/global"
	"github.com/4everlynn/journal/support"
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

const (
	Day     = 0
	Week    = 1
	Month   = 2
	Quarter = 3
)

// generate daily report
func generate(cmd *cobra.Command, args []string) {
	cfg := GetConfig()
	// get gen type
	genType := ResolveType(cmd)
	// weather to use last period's data
	isLast := isLastPeriod(cmd)
	start, end := GetDateRange(genType, isLast)
	// echo head info
	echoHead(start, end, genType)
	for service, git := range cfg.Git {
		if git.Disable == true {
			continue
		}
		all := make([]support.GitLog, 0)
		repo := strings.Join([]string{git.Path, ".git", "logs", "HEAD"}, global.Separator)
		git.Path = repo
		_, err := os.Stat(repo)
		if err == nil {
			// add report content
			all = append(all, ReportFromGit(git, start.Unix(), end.Unix()-1)...)
			if len(all) > 0 {
				// inject real path
				git.Path = repo
				println(color.FgCyan.Render(git.Name))
				formatter := new(support.GitLogDailyFormatter)
				println(formatter.Format(all))
			}
		} else if os.IsNotExist(err) {
			color.Danger.Printf(".git is missing in repository %s\n", service)
		}
	}
}

func isLastPeriod(cmd *cobra.Command) bool {
	isLast, err := cmd.Flags().GetBool("last")
	if err == nil {
		return isLast
	}
	return false
}

func echoHead(start time.Time, end time.Time, genType int) {
	if genType == Day {
		println(time.Now().Format("20060102") + " Daily")
	} else if genType == Week {
		println(start.Format("20060102") + " ～ " + end.AddDate(0, 0, -1).Format("20060102") + " Weekly")
	} else if genType == Month {
		println(start.Format("200601") + " ～ " + end.Format("200601") + " Monthly")
	} else if genType == Quarter {
		println(start.Format("200601") + " ～ " + end.Format("200601") + " Quarterly")
	}
}

func GetDateRange(genType int, isLast bool) (time.Time, time.Time) {
	timeStr := time.Now().Format("2006-01-02")
	start, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	end, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if genType == Day {
		end = end.AddDate(0, 0, 1)
		if isLast {
			start.AddDate(0, 0, -1)
			end = end.AddDate(0, 0, -1)
		}
	} else if genType == Week {
		start = start.AddDate(0, 0, int(time.Monday-start.Weekday()))
		end = end.AddDate(0, 0, int(time.Friday-end.Weekday())+1)
		if isLast {
			start = start.AddDate(0, 0, -7)
			end = end.AddDate(0, 0, -7)
		}
	} else if genType == Month {
		start = time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
		end = start.AddDate(0, 1, 0)
		if isLast {
			start = start.AddDate(0, -1, 0)
			end = end.AddDate(0, -1, 0)
		}
	} else if genType == Quarter {
		start = time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
		end = start.AddDate(0, 3, 0)
		if isLast {
			start = start.AddDate(0, -3, 0)
			end = end.AddDate(0, -3, 0)
		}
	}
	return start, end
}

// ResolveType  determine the build type
func ResolveType(cmd *cobra.Command) int {

	isQuarter, err := cmd.Flags().GetBool("quarter")
	if err == nil && isQuarter {
		return Quarter
	}

	isMonth, err := cmd.Flags().GetBool("month")
	if err == nil && isMonth {
		return Month
	}

	isWeek, err := cmd.Flags().GetBool("week")
	if err == nil && isWeek {
		return Week
	}

	isDay, err := cmd.Flags().GetBool("day")
	if err == nil && isDay {
		return Day
	}

	return 0
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
	stylishCmd.Flags().BoolP("last", "l", false, "select previous period data")
	stylishCmd.Flags().BoolP("day", "d", true, "output as work report by day")
	stylishCmd.Flags().BoolP("week", "w", false, "output as work report by week")
	stylishCmd.Flags().BoolP("month", "m", false, "output as work report by month")
	stylishCmd.Flags().BoolP("quarter", "q", false, "output as work report by quarter")
}
