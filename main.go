package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	var (
		file *os.File
		info os.FileInfo
		err  error
	)
	cmdPath := os.Args[1]
	if file, err = os.Open(cmdPath); err != nil {
		log.Fatal(err.Error())
	}
	if info, err = file.Stat(); err != nil {
		log.Fatal(err.Error())
	}
	now := time.Now()
	if isSameDay(now, info.ModTime()) {
		// nothing to do
		os.Exit(0)
	}
	if err = os.Chtimes(cmdPath, now, now); err != nil {
		log.Fatal(err.Error())
	}
	cmd := exec.Command(cmdPath, os.Args[2:]...)
	if err = cmd.Start(); err != nil {
		log.Fatal(err.Error())
	}
	os.Exit(0)
}
func isSameDay(rhs, lhs time.Time) bool {
	return rhs.Year() == lhs.Year() && rhs.Month() == lhs.Month() && rhs.Day() == lhs.Day()
}
