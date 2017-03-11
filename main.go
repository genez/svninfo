package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"os/exec"
)

var (
	golang = kingpin.Flag("golang", "Generate Go struct with version info.").Short('g').Default("true").Bool()
	pkg    = kingpin.Flag("package", "Package for the Go struct").Short('p').Default("main").String()
)

const structTemplate string = `//go:generate svninfo golang %s
package %s

var (
	Commit_Revision  string = "%s"
	Commit_TimeStamp string = "%s"
)
`

func main() {
	kingpin.Parse()
	if *golang {
		f, err := os.Create("version.go")
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		f.WriteString(fmt.Sprintf(structTemplate, *pkg, *pkg, runSvnVersion(), runSvnInfo()))
	} else {

	}
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
