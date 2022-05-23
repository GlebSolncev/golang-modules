package structuring

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var pathToTemplate = "internal/structuring/template.tmpl"

/// generateNewFile create new File with getters and setters. From Body info
func (b Body) generateNewFile(path string) {
	// parse the template
	tmpl, _ := template.ParseFiles(pathToTemplate)

	// Create a new file
	filename := fmt.Sprintf("%s/%s_structuring.go", path, strings.ToLower(b.FileInfo.NameStructure))
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Apply the template to the Body map and write the result to file.
	err = tmpl.Execute(file, b)
	if err != nil {
		panic(err)
	}
}
