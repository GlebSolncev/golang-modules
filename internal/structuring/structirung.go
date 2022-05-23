package structuring

import "golang.org/x/tools/go/packages"

// getPackageName by path check package and get info about struct and etc.
func getPackageName(path string) *packages.Package {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		Tests: false,
	}, path)

	if err != nil {
		panic(err)
	}

	if len(pkgs) != 1 {
		panic("Package name not found")
	}

	return pkgs[0]
}

func Execute(path string) {
	var myBody = Body{}
	// Parser structs
	myBody.parseStructure(getPackageName(path))

	// create file and insert code
	myBody.generateNewFile(path)
}
