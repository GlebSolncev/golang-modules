package main

//Structuring

import (
	"golang-modules/internal/structirung"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func getPath() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	return dir
}

func getPackageName() *packages.Package {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		Tests: false,
	}, getPath())

	if err != nil {
		panic(err)
	}

	if len(pkgs) != 1 {
		panic(err)
	}

	return pkgs[0]
}

func main() {
	structirung.Execute(getPath(), getPackageName())
}
