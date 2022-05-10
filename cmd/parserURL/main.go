package main

import (
	"fmt"
	"golang-modules/internal/parser/file"
	"golang-modules/internal/parser/http"
	"os"
	"strings"
	"sync"
)

func recoverUrls(req http.Request) {
	fmt.Println(req)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input args(links)")
	}

	urls := os.Args[1:]

	for _, url := range urls {
		fmt.Println("START >>> " + url)
		workWithUrl(url)
	}

}

func workWithUrl(url string) {
	var req http.Request
	lock := sync.Mutex{}
	if strings.Index(url, "https:/") == -1 {
		url = "https://" + url
	}

	if file.CanAddToCollect(url) {

		req = http.SendRequest(url).GetLinks()
		if len(req.CollectLinks) > 1 {
			req.Len = len(req.CollectLinks)
			lock.Lock()
			file.AddLink(req)
			lock.Unlock()
		}

		defer recoverUrls(req)
	} else {
		fmt.Println("[NATC] " + url)
	}

	if len(req.CollectLinks) > 2 {
		wg := sync.WaitGroup{}

		//wg.Add(len(req.CollectLinks))
		for _, newUrl := range req.CollectLinks {
			wg.Add(1)
			fmt.Println("NEW SPACE " + newUrl)

			workWithUrl(newUrl)

			//go func(req http.Request) {
			//lock.Lock()

			//lock.Unlock()
			//wg.Done()
			//}(req)
		}
		//wg.Wait()

	}
}
