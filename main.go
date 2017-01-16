package main

import (
	"os"
	"time"
	"strings"
	"path"
	"path/filepath"
	"github.com/jmoiron/sqlx"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
)



var db *sqlx.DB
const DbName = "gohistgrep.db"


func InitDB(filepath string) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", path.Join(os.Getenv("HOME"), ".history", DbName))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func FilesUpdate(filename string, filedate time.Time) {
	files := []File{}
	err := db.Select(&files, "SELECT * FROM files WHERE filename=$1", filename)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO files (filename, date) VALUES ($1, $2)", filename, filedate)
		tx.Commit()
		log.Info("File added")
	} else {
		log.Info("File skiped")
	}
}

func getAllFiles() {
	searchDir := path.Join(os.Getenv("HOME"), ".history")
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileList {
		fi, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		if !fi.IsDir() {
			log.Info("Found file " + file)

			// FIXME:
			sdata := strings.Split(file, "/")
			filename := sdata[len(sdata) - 1]
			if filename == DbName {
				continue
			}
			month := sdata[len(sdata) - 2]
			year := sdata[len(sdata) - 3]
			sfilename := strings.Split(filename, ".")
			second := strings.Split(sfilename[len(sfilename) - 1],"_")[0]
			minute := sfilename[len(sfilename) - 2]
			hour := sfilename[len(sfilename) - 3]
			day := sfilename[len(sfilename) - 4]
			layout:= "02-01-2006 15:04:05 MST"
			real := day + "-" + month + "-" + year + " " + hour + ":" + minute + ":" + second + " MSK"
			t, _ := time.Parse(layout, real)
			FilesUpdate(file, t)
		}
	}
}

func main() {
	db = InitDB("")
	db.MustExec(Schema)
	if len(os.Args) == 2 {
		SearchCommand(os.Args[1])
	} else {
		getAllFiles()
		ParseFiles()
	}
}
