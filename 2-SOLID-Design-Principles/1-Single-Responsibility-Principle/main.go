package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++

	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)

	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	//TODO
}

// BREAKING SINGLE RESPONSIBILITY PATTERN
func (j *Journal) SaveFile(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// ***************************************

// separation of concerns
var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

// ***********************

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	j.AddEntry("I saw a toad")
	fmt.Print(j.String())

	//
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journalPersistence.txt")
}
