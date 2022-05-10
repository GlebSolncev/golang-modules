package file

import (
	"encoding/json"
	"fmt"
	"golang-modules/internal/parser/http"
	"golang-modules/pkg/path"
	"io/ioutil"
	"os"
)

var pathFile = path.GetBasePath("storage/parser/out.json")

func AddLink(req http.Request) string {
	jsonFile, _ := os.Open(pathFile)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var links []http.Request
	json.Unmarshal(byteValue, &links)

	if CanAddToCollect(req.Url) {
		links = append(links, req)
		file, _ := json.MarshalIndent(links, "", " ")
		ioutil.WriteFile(pathFile, file, 0644)
	}

	return "OK"
}

func CanAddToCollect(url string) bool {
	var status = true
	var links []http.Request

	jsonFile, _ := os.Open(pathFile)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &links)

	for _, item := range links {
		if item.Url == url {
			status = false
		} else {
			fmt.Println("[NF]" + url + item.Url)
		}
	}

	return status
}
