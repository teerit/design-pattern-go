package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// seperation of concerns
// In this example is persistence point
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0664)
}

func (j *Journal) Load(filename string) {
	// TODO: implement
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// TODO: implement
}

var LineSeparator = "\n"

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0664)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
