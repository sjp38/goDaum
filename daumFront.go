package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func daumFrontPage() string {
	cmd := "w3m -dump www.daum.net"
	tokens := strings.Fields(cmd)
	out, err := exec.Command(tokens[0], tokens[1:]...).Output()
	if err != nil {
		fmt.Printf("Failed to execute wem: %s\n", err)
		return ""
	}

	return string(out)
}

func sectionFor(page, title string) string {
	pars := strings.Split(page, "\n\n")
	for idx, p := range pars {
		if p == title {
			return pars[idx+1]
		}
	}

	return ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: daumFront <section title>\n")
		os.Exit(1)
	}
	page := daumFrontPage()

	title := strings.Join(os.Args[1:], " ")
	fmt.Printf("%s\n", sectionFor(page, title))
}
