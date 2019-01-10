package main

import (
	"regexp"
)

var segExp = regexp.MustCompile("[^\\n.。!?！？]+")

func segment(text string) []string {
	return segExp.FindAllString(text, -1)
}
