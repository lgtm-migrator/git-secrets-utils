package lib

import (
	"strings"
)

func Value() string {
	return "hello golang"
}

type LineErrorInformation struct {
	File    string
	Line    string
	Content string
}

func MatchLineString(lineStr string) *LineErrorInformation {
	matched := strings.SplitN(lineStr, ":", 3)
	if len(matched) == 3 {
		return &LineErrorInformation{matched[0], matched[1], strings.TrimSpace(matched[2])}
	}
	return nil
}
