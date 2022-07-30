package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

type Regexp struct{}

// CheckLength 检查字符长度是否符合要求
func (s *Regexp) CheckLength(string string, minLen, maxLen int) bool {
	return utf8.RuneCountInString(string) >= minLen && utf8.RuneCountInString(string) <= maxLen
}

// CheckUserName 数字+字母 1~10 位
func (s *Regexp) CheckUserName(string string) bool {
	if !(s.CheckLength(string, 1, 10)) {
		return false
	}
	stringSplit := strings.Split(string, "")
	for _, str := range stringSplit {
		result, err := regexp.MatchString("^[0-9a-zA-Z]$", str)
		if !result || err != nil {
			return false
		}
	}
	return true
}

// CheckPassword 数字+字母+特殊符号 4-32 位
func (s *Regexp) CheckPassword(string string) bool {
	if !(s.CheckLength(string, 4, 32)) {
		return false
	}
	stringSplit := strings.Split(string, "")
	for _, str := range stringSplit {
		result, err := regexp.MatchString("^[0-9a-zA-Z~!@#$%^&*()_+={}|<>,.?;:]$", str)
		if !result || err != nil {
			return false
		}
	}
	return true
}

// CheckScore 小数 小数点前一位 + 小数点后一位
func (s *Regexp) CheckScore(string string) bool {
	stringSplit := strings.Split(string, "")
	for _, str := range stringSplit {
		result, err := regexp.MatchString("^\\d\\.\\d$", str)
		if !result || err != nil {
			return false
		}
	}
	return true
}
