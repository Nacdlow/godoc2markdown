package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	titleExp = regexp.MustCompile("^[A-Z]+$")
	funcExp  = regexp.MustCompile("^func ")
	descExp  = regexp.MustCompile("^    ")
)

func main() {
	dat, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	codeBlock := ""

	for _, line := range lines {
		if len(line) == 0 {
			if len(codeBlock) == 0 {
				fmt.Println()
			} else {
				fmt.Printf("```go%s\n```\n", codeBlock)
				codeBlock = ""
			}
		} else {
			if titleExp.Match([]byte(line)) {
				fmt.Printf("## %s\n", line)
			} else if descExp.Match([]byte(line)) {

				fmt.Printf("%s\n", descExp.ReplaceAll([]byte(line), []byte("")))
			} else if funcExp.Match([]byte(line)) {
				name := strings.Split(strings.Split(line, " ")[1], "(")[0]
				fmt.Printf("## func %s\n```go\n%s\n```\n", name, line)
			} else {
				codeBlock = fmt.Sprintf("%s\n%s", codeBlock, line)
			}
		}
	}
}
