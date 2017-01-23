package main

import (
	"fmt"
	"strings"
	"strconv"
	log "github.com/Sirupsen/logrus"
)


func SearchCommand(command string) {
	if strings.Contains(command, "+") {
		data := strings.Split(command, "+")
		out := ""
		if len(data) > 0 {
			out += "%"
		}
		for _, d := range data {
			if d != "" {
				out += d + "%"
			}
		}
		command = out
	} else {
		if !strings.HasPrefix(command, "%") && !strings.HasSuffix(command,"%") {
			command = "%" + command + "%"
		}
	}
	logs := []Log{}
	err := db.Select(&logs, "SELECT * FROM log WHERE LOWER(command) LIKE $1 ORDER BY popularity DESC", command)
	if err != nil {
		log.Fatal(err)
	}
	if len(logs) > 0 {
		for _, logData := range logs {
			fmt.Println("[" + strconv.Itoa(logData.Popularity) + "] " + logData.Command)
		}
	} else {
		fmt.Println("No result")
	}
}
