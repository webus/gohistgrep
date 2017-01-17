package main

import (
	"os"
	"strconv"
	"time"
	"bufio"
	"strings"
	log "github.com/Sirupsen/logrus"
)

func LogUpdateDb(command string, filename string) {
	if strings.Contains(command, "gohistgrep ") {
		return
	}
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM log WHERE command=$1", command)
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		db.MustExec("INSERT INTO log (command, filename) VALUES ($1, $2)", command, filename)
	} else {
		db.MustExec("UPDATE log SET popularity=popularity+1 WHERE command=$1", command)
	}
}

func LogUpdate(filename string) int {
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
	return len(lines)
}


func ParseFiles() {
	files := []File{}
	err := db.Select(&files, "SELECT * FROM files WHERE process_date IS NULL ORDER BY date ASC")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Info("Processing file " + file.Filename)
		linesCount := LogUpdate(file.Filename)
		db.MustExec("UPDATE files SET process_date=$1 WHERE filename=$2", time.Now(), file.Filename)
		log.Info("Processed " + strconv.Itoa(linesCount) + " lines")
	}
}
