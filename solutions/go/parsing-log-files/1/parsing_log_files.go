package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	var regex = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	return regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	var regex = regexp.MustCompile(`<[~*=-]*>`)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var regex = regexp.MustCompile(`(?i)".*password.*"`)
	var count = 0
	for _, line := range lines {
		if regex.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	var regex = regexp.MustCompile(`end-of-line\d*`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var regex = regexp.MustCompile(`User\s+(\w+)`)
	for index, line := range lines {
		var submatches = regex.FindStringSubmatch(line)
		if len(submatches) > 0 {
			lines[index] = "[USR] " + submatches[1] + " " + line
		}
	}
	return lines
}
