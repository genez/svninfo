package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("svn", "info", "--xml")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	var info = &Info{}
	if err := xml.NewDecoder(stdout).Decode(&info); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) != 2 {
		log.Fatalln(`wrong number of arguments: "revision" or "timestamp" argument must be specified`)
	}

	if os.Args[1] == "revision" {
		fmt.Print(info.Entry.Commit.Revision)
	} else if os.Args[1] == "timestamp" {
		fmt.Print(info.Entry.Commit.Date)
	} else {
		log.Fatalln(`error: "revision" or "timestamp" argument must be specified`)
	}
}
