package main

import (
	"strings"
)

type dict map[string]string

var dictionary dict

func init() {
	dictionary = dict{
		"おおたわ":  "おおたわです",
		"OTYM":  "おたやま",
		"test":  "test desu",
		"コンパイル": "難しい",
	}
}

func getAutoResponses(s string) []string {
	var replyMessages []string
	for key, value := range dictionary {
		if strings.Contains(s, key) {
			replyMessages = append(replyMessages, value)
		}
	}
	return replyMessages
}

func getResponceDict() dict {
	return dictionary
}
