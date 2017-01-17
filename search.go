package main

import (
	"fmt"
	"strconv"
	log "github.com/Sirupsen/logrus"
)

func SearchCommand(command string) {
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
