package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

func makeReadme(filename string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://rssproxy.migor.org/api/w2f?v=0.1&url=https%3A%2F%2Fme.xn--qck4cud2cb.com%2Fposts%2F&link=.%2Fh2%5B1%5D%2Fa%5B1%5D&context=%2F%2Fmain%5B1%5D%2Fdiv%5B1%5D%2Farticle&x=pn&re=none&out=atom&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0eXBlIjoiYW5vbnltb3VzIiwiaWF0IjoxNjY5Njg5NjAzfQ.MFT1yhAaddPVif9fAJMsu5QQhJVze_-RYkqTeTRamvw")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	blogItem := feed.Items[0]

	rand.Seed(time.Now().UnixNano())
	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	header := "![](https://i.imgur.com/qiTHeFR.png) \n \n [<p align='center'> <img src='https://raw.githubusercontent.com/gruberdev/gruberdev/main/icons/left.svg' alt='left'>&emsp;<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='25'>](https://github.com/gruberdev)&emsp;[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/dev-dot-to.svg' alt='dev' height='25'>](https://dev.to/cloudgruber)&emsp;<!-- markdown-link-check-disable -->[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='25'>](https://www.linkedin.com/in/rodrigo-gruber/)<!-- markdown-link-check-enable -->&emsp;[<img src='https://api.iconify.design/simple-icons:codesandbox.svg?height=24' alt='codesandbox' height='25'>](https://killercoda.com/gruber)&emsp;[<img src='https://api.iconify.design/fa-brands:free-code-camp.svg?height=24' alt='Reddit' height='25'>](https://codestats.net/users/gruber)&emsp;[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/gitlab.svg' alt='gitlab' height='25'>](https://gitlab.com/gruberx) <img src='https://raw.githubusercontent.com/gruberdev/gruberdev/main/icons/right.svg' alt='right'> \n \n "
	blog := "<p align='center'>The last article I've written: </p> <a align='center' href='" + blogItem.Link + "'> <p align='center'>" + blogItem.Title + "</p> </a. \n \n"
	updated := "<sub> <p align='center'> <sup>This README was last updated on: " + date + ".</sup> </p> </sub> \n \n \n <p align='center'> \n \n </p>"
	data := fmt.Sprintf("%s\n\n%s\n\n%s\n", header, blog, updated)

	// Prepare file with a light coating of os
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Bake at n bytes per second until golden brown
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	makeReadme("../README.md")
}
