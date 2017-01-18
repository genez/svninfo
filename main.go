package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln(`wrong number of arguments: "revision" or "timestamp" argument must be specified`)
	}

	var result string

	if os.Args[1] == "revision" {
		result = runSvnVersion()
	} else if os.Args[1] == "timestamp" {
		result = runSvnInfo()
	} else {
		log.Fatalln(`error: "revision" or "timestamp" argument must be specified`)
	}

	fmt.Print(result)
}

func runSvnVersion() string {
	cmd := exec.Command("svnversion", "-q")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Scan()

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return scanner.Text()
}

func runSvnInfo() string {
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

	return info.Entry.Commit.Date
}
