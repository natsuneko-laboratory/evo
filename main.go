package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/exec"
	"path"
	"time"

	cr "github.com/robfig/cron/v3"
)

func getCacheKey(exec string) string {
	s := sha256.Sum256([]byte(exec))
	return hex.EncodeToString(s[:])
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func restoreFromCacheContent(run string, loc string) {
	if isFileExists(loc) {
		content, err := os.ReadFile(loc)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(content))
	} else {
		execute(run, loc)
	}
}

func execute(run string, loc string) {
	out, err := exec.Command("sh", "-c", run).Output()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(loc, out, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func main() {
	cron := flag.String("cron", "* * * * *", "The cron schedule to use")
	run := flag.String("run", "", "The command to execute")
	store := flag.String("store", "/tmp/", "The path to store the results")
	delay := flag.Int("delay", 0, "The delay in seconds to wait before executing the command")
	random := flag.Bool("random", false, "Add a random delay to the execution")
	flag.Parse()

	parser := cr.NewParser(cr.Minute | cr.Hour | cr.Dom | cr.Month | cr.Dow)
	s, err := parser.Parse(*cron)
	if err != nil {
		log.Fatal(err)
	}

	loc := path.Join(*store, getCacheKey(*run))
	now := time.Now()
	previous := now

	if *random {
		val := rand.N(*delay)
		time.Sleep(time.Duration(val) * time.Second)
	} else {
		time.Sleep(time.Duration(*delay) * time.Second)
	}

	if isFileExists(loc) {
		fi, err := os.Stat(loc)
		if err != nil {
			log.Fatal(err)
		}
		previous = fi.ModTime()
	}

	if s.Next(previous).Unix() <= now.Unix() {
		execute(*run, loc)
	} else {
		restoreFromCacheContent(*run, loc)
	}
}
