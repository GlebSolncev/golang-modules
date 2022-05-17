package structirung

import (
	"flag"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func parseStructs(pkg *packages.Package) Struct {
	scope := pkg.Types.Scope()
	myStruct := Struct{}
	nameStruct := getNameStruct()
	for _, name := range scope.Names() {
		if name == nameStruct {
			st, ok := scope.Lookup(name).Type().Underlying().(*types.Struct)
			if !ok {
				continue
			}

			myStruct = Struct{
				Name:   name,
				Fields: parseFields(st),
			}
		}
	}

	return myStruct
}

func parseFields(st *types.Struct) []Field {
	fields := make([]Field, st.NumFields())
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		fields[i] = Field{
			Name: field.Name(),
			Type: field.Type(),
		}
	}
	return fields
}

func getNameStruct() string {
	pName := flag.String("name", "", "Name struct")
	flag.Parse()

	return *pName
}
