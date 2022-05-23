package structirung

import "golang.org/x/tools/go/packages"

var importBody []string
var settersFunc []string
var gettersFunc []string
var argAll bool

func Execute(path string, pkg *packages.Package) {
	// Parser structs
	myStruct := parseStructs(pkg)

	// create file and insert code
	generateNewFile(path, pkg.Name, myStruct)
}
