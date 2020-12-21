package support

import (
	"strconv"
	"strings"
)

type formatter interface {
	Format(logs GitLog)
}

type GitLogDailyFormatter struct{}

func (GitLogDailyFormatter) Format(logs []GitLog) string {
	builder := make([]string, 0)
	for index, log := range logs {
		builder = append(builder, strings.Join([]string{"  ", strconv.Itoa(index + 1), "、", log.Message}, ""))
	}
	return strings.Join(builder, "\n")
}
