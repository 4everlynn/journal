package cmd

import (
	"bufio"
	"diswares.com.journal/global"
	"diswares.com.journal/support"
	"github.com/martinlindhe/notify"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start a daemon to observe whether the warehouse code is committed",
	Long:  `Start a daemon to observe whether the warehouse code is committed`,
	Run: func(cmd *cobra.Command, args []string) {
		config := GetConfig()
		for _, git := range config.Git {
			path := strings.Join([]string{git.Path, ".git", "logs", "HEAD"}, global.Separator)
			log.Println(path)
			log.Println(git)
			monitoring(path, func(bytes []byte) {
				gitLog := support.ParseLine(string(bytes))
				if len(gitLog.Message) > 0 {
					builder := make([]string, 0)
					builder = append(builder, "项目 ", git.Name, "提交了代码记得要规范地提交代码，以便于快速按时提交日报")
					notify.Notify("Journal", "监测到代码提交", strings.Join(builder, ""), "")
				}
			})
		}
	},
}

// 文件监控
func monitoring(filePth string, hook func([]byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	_, _ = f.Seek(0, 2)
	for {
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hook(line)
	}

}

func init() {
	rootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
