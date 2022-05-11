package main

import (
	"fmt"
	"golang-modules/internal/parser/file"
	"golang-modules/internal/parser/http"
	"os"
	"sync"
	"time"
)

var links []http.Request
var LinksForScan []string
var lock = sync.Mutex{}

func main() {
	url := os.Args[1]
	startData := http.SendRequest(url).GetLinks()

	UpdLinksForScan := make(chan string)
	UpdLinks := make(chan string)
	UpdateFileData := make(chan string)

	go func(data http.Request) {
		// Check var LInks from file or start Data and update var NeedToScanLinks
		for {
			if len(links) == 0 {
				links = file.GetContentFromFile()
				if len(links) == 0 {
					links = append(links, data)
				}
			}

			if len(links) == 0 {
				fmt.Println("Done")
				os.Exit(0)
			}

			for _, link := range links {
				// check from var NeedScanLinks for sync links na uniq
				for _, newLink := range link.CollectLinks {
					var status = true

					for _, iHaveLink := range LinksForScan {
						if newLink == iHaveLink {
							status = false
							break
						}
					}

					if status {
						LinksForScan = append(LinksForScan, newLink)
					}
				}
			}
			UpdLinksForScan <- "[+ NS]: " + fmt.Sprintf("%v", len(LinksForScan))
			time.Sleep(time.Second * 1)
		}
	}(startData)

	go func() {
		// Work with LinksForScan and update Links And clear NeedScanLinks
		for {
			if len(LinksForScan) >= 1 {
				for _, link := range LinksForScan {
					if file.CanAddToCollect(link) { // Check link info in file
						links = append(links, http.SendRequest(link).GetLinks())
					}
				}
			}

			LinksForScan = []string{}
			UpdLinks <- "[+ L]]: " + fmt.Sprintf("%v", len(links))
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		// Add links to File and clear var Links for work retry with new data
		for {
			if len(links) >= 1 {
				for _, link := range links {
					if file.CanAddToCollect(link.Url) {
						if len(link.CollectLinks) >= 1 {
							lock.Lock()
							file.AddLink(link)
							lock.Unlock()
						}
					}
				}
			}

			links = []http.Request{}
			UpdateFileData <- "[- L + F]: " + fmt.Sprintf("%v", len(links))
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-UpdLinksForScan:
				fmt.Println(msg1)

			case msg2 := <-UpdLinks:
				fmt.Println(msg2)

			case msg3 := <-UpdateFileData:
				fmt.Println(msg3)
			}

		}
	}()

	var input string
	fmt.Scanln(&input)
}
