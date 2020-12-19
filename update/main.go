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
	feed, err := fp.ParseURL("https://blog.tulpas.dev/feed.xml")
	if err != nil {
		log.Fatalf("error getting feed: %v", err)
	}
	// Get the freshest item
	blogItem := feed.Items[0]


	rand.Seed(time.Now().UnixNano())
	date := time.Now().Format("2 Jan 2006")

	// Whisk together static and dynamic content until stiff peaks form
	header := "![](https://i.imgur.com/qiTHeFR.png) \n \n <h3 align='center'>DevOps</h3> \n \n [<p align='center'> <img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/git.svg' alt='dribbble' height='25'>](https://sr.ht/~tsyklon/) [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='25'>](https://github.com/verifiedgruber)   [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/dev-dot-to.svg' alt='dev' height='25'>](https://dev.to/verifiedgruber) [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='25'>](https://www.linkedin.com/in/rodrigogruber/) [<img src='https://api.iconify.design/simple-icons:codesandbox.svg?height=24' alt='codesandbox' height='25'>](https://codesandbox.io/u/VerifiedGruber) [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/stackoverflow.svg' alt='stackoverflow' height='25'>](https://stackexchange.com/users/8975552/rodrigo-gruber)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/youtube.svg' alt='YouTube' height='25'>](https://www.youtube.com/channel/UC51Kh4tA9IvPsGly15Y8vZA)  [<img src='https://api.iconify.design/fa-brands:free-code-camp.svg?height=24' alt='Reddit' height='25'>](https://codestats.net/users/gruber)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/gmail.svg' alt='gmail' height='25'>](mailto:736d9d13-b432-45b3-a5e5-f012126caca9@gruber.anonaddy.com)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/gitlab.svg' alt='gitlab' height='25'>](https://gitlab.com/gruberx)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/youtubemusic.svg' alt='youtubemusic' height='25'>](https://www.last.fm/user/rpgruber)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/steam.svg' alt='steam' height='25'>](https://steamcommunity.com/id/takerukazuya)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/spotify.svg' alt='spotify' height='25'>](https://open.spotify.com/user/rodrigogruber) \n \n "
	blog := "<p align='center'>This my latest blog post: </p> <a align='center' href='" + blogItem.Link + "'> <p align='center'>" + blogItem.Title + "</p> </a. \n \n"
	updated := "<sub> <p align='center'> Last updated on: " + date + ".</p> </sub> \n \n \n <p align='center'> \n \n <img align='center' src='./githubterm.svg' /> \n \n </p>"
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
