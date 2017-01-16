package main

import (
	"time"
)

var Schema = `
CREATE TABLE IF NOT EXISTS files (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  filename TEXT NOT NULL UNIQUE,
  date DATETIME NOT NULL,
  process_date DATETIME NULL
);

CREATE TABLE IF NOT EXISTS log (
  id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  command TEXT NOT NULL,
  filename TEXT NOT NULL
);
`

type File struct {
	Id int `db:"id"`
	Filename string `db:"filename"`
	Date *time.Time `db:"date"`
	ProcessDate *time.Time `db:"process_date"`
}

type Log struct {
	Id int `db:"id"`
	Command string `db:"command"`
	Filename string `db:"filename"`
}
