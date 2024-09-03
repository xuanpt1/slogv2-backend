package utils

import "regexp"

const (
	//RegexEmail    = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	RegexPureText = `[^a-zA-Z0-9\\u4e00-\\u9fa5]`
)

func GetPureEmailRegex(str string) string {
	reg := regexp.MustCompile(RegexPureText)
	return reg.ReplaceAllString(str, "")
}
