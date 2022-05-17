package structirung

import (
	"flag"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
	"unicode"
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
				Fields: parseFields(strings.Replace(pkg.Name, "_", "-", 1), st),
			}
		}
	}

	return myStruct
}

func parseFields(pkgName string, st *types.Struct) []Field {
	fields := make([]Field, st.NumFields())
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		name := field.Name()

		fields[i] = Field{
			Name:    name,
			Type:    field.Type(),
			Private: unicode.IsLower(rune(name[0])),
		}.getImportsForFields(pkgName)
	}

	return fields
}

// parse CLI args(flags)
func getNameStruct() string {
	argName := flag.String("name", "", "Name struct")
	argsAll := flag.Bool("all", false, "If inp this arg, i get all fields for struct")
	flag.Parse()

	argAll = *argsAll
	return *argName
}

func (f Field) getImportsForFields(pkgName string) Field {
	fieldType := f.Type.String()
	arrPKGs := strings.Split(fieldType, "/")

	name := arrPKGs[len(arrPKGs)-1]
	arrPKGs = strings.Split(name, ".")
	if len(arrPKGs) == 1 {
		f.TypeS = fieldType
		return f
	}

	if arrPKGs[0] != pkgName {
		f.TypeS = name
		f.Imports = strings.Split(fieldType, ".")[0]
		return f
	}
	f.TypeS = arrPKGs[1]

	return f
}
