package main

import (
	"encoding/xml"
)

type Info struct {
	XMLName xml.Name `xml:"info"`
	Entry   Entry    `xml:"entry"`
}

type Entry struct {
	Commit Commit `xml:"commit"`
}

type Commit struct {
	Revision string `xml:"revision,attr"`
	Date     string `xml:"date"`
}
