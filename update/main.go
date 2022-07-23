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
	feed, err := fp.ParseURL("https://rssproxy.gr.wtf/api/feed?url=https%3A%2F%2Fme.xn--qck4cud2cb.com%2Fposts%2F&pContext=%2F%2Fbody%2Fmain%5B1%5D%2Fdiv%5B1%5D%2Farticle&pLink=.%2Fh2%5B1%5D%2Fa%5B1%5D&o=RSS&x=s")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the freshest item
	blogItem := feed.Items[0]

	rand.Seed(time.Now().UnixNano())
	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	header := "![](https://i.imgur.com/qiTHeFR.png) \n \n <p align='center'><img src='https://github-readme-stats.vercel.app/api/wakatime?username=gruber&api_domain=wakapi.gruber.dev.br&custom_title=Weekly%20Metrics&langs_count=4' alt='center'> </p> \n \n [<p align='center'> <img src='https://raw.githubusercontent.com/gruberdev/gruberdev/main/icons/left.svg' alt='left'> <img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/git.svg' alt='dribbble' height='25'>](https://sr.ht/~tsyklon/)&emsp;[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='25'>](https://github.com/gruberdev)&emsp;[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/dev-dot-to.svg' alt='dev' height='25'>](https://dev.to/cloudgruber)&emsp;<!-- markdown-link-check-disable -->[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='25'>](https://www.linkedin.com/in/rodrigo-g/)<!-- markdown-link-check-enable -->&emsp;[<img src='https://api.iconify.design/simple-icons:codesandbox.svg?height=24' alt='codesandbox' height='25'>](https://codesandbox.io/u/cloudgruber)&emsp;[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/stackoverflow.svg' alt='stackoverflow' height='25'>](https://stackexchange.com/users/8975552/rodrigo-gruber)&emsp;[<img src='https://api.iconify.design/fa-brands:free-code-camp.svg?height=24' alt='Reddit' height='25'>](https://codestats.net/users/gruber)&emsp;[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/gitlab.svg' alt='gitlab' height='25'>](https://gitlab.com/gruberx) <img src='https://raw.githubusercontent.com/gruberdev/gruberdev/main/icons/right.svg' alt='right'> \n \n "
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
