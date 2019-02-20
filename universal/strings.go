package dev_tempalte_universal

import "strings"

func ShorterStr(s string, sep string, count int) string {
	ss := strings.Split(s, sep)
	ss = ss[:len(ss)-count]
	return strings.Join(ss, sep)
}

func ExtractFromEnd(s string, sep string, count int) string {
	ss := strings.Split(s, sep)
	ss = ss[len(ss)-count:]
	return strings.Join(ss, sep)
}