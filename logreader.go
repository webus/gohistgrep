package main

import (
	"os"
	"time"
	"bufio"
	"strings"
	log "github.com/Sirupsen/logrus"
)

func LogUpdateDb(command string, filename string) {
	logs := []Log{}
	if strings.Contains(command, "gohistgrep ") {
		return
	}
	err := db.Select(&logs, "SELECT * FROM log WHERE command=$1", command)
	if err != nil {
		log.Fatal(err)
	}
	if len(logs) == 0 {
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO log (command, filename) VALUES ($1, $2)", command, filename)
		tx.Commit()
	}
}

func LogUpdate(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		LogUpdateDb(line, filename)
	}
}

func ParseFiles() {
	files := []File{}
	err := db.Select(&files, "SELECT * FROM files WHERE process_date IS NULL ORDER BY date ASC")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Info("Processing file " + file.Filename)
		LogUpdate(file.Filename)
		tx := db.MustBegin()
		tx.MustExec("UPDATE files SET process_date=$1 WHERE filename=$2", time.Now(), file.Filename)
		tx.Commit()
		log.Info("Processing file done")
	}
}
