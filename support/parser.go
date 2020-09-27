package support

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type LogParser interface {
	// parse head file to GitLog
	Parse(headFilePath string, start int64, end int64) []GitLog
}

type GitLogParser struct{}

func (GitLogParser) Parse(headFilePath string, start int64, end int64) []GitLog {
	logs, err := os.Open(headFilePath)
	if err == nil {
		//log.Println(green("[INFO]"), green(headFilePath), green("OPENED SUCCESS"))
		scanner := bufio.NewScanner(logs)
		git := make([]GitLog, 0)
		for scanner.Scan() {
			data := ParseLine(scanner.Text())
			if data.timestamp >= start && data.timestamp <= end {
				if len(data.Message) > 0 {
					git = append(git, data)
				}
			}
		}

		err := logs.Close()
		if err == nil {
			return git
		}
	}
	return nil
}

func ParseLine(text string) GitLog {
	base := strings.Split(text, "<")
	split := strings.Split(base[0], " ")
	timestamp, _ := strconv.ParseInt(strings.Split(base[1], " ")[1], 10, 64)
	messageArr := strings.Split(base[1], "commit:")
	if len(messageArr) < 2 {
		return GitLog{}
	}
	git := GitLog{
		strings.Join(split[2:], " "),
		timestamp,
		strings.Trim(messageArr[1], " "),
	}
	return git
}
