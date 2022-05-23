package structuring

import (
	"flag"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
	"unicode"
)

// parseStructs Parsing from package and get data about struct. Working with Body
func (b *Body) parseStructure(pkg *packages.Package) {
	scope := pkg.Types.Scope()
	b.FlagInfo.setNameStruct()
	for _, name := range scope.Names() {
		if name == b.FlagInfo.NameStruct {
			st, ok := scope.Lookup(name).Type().Underlying().(*types.Struct)
			if !ok {
				continue
			}

			b.FileInfo.NameStructure = name
			b.FileInfo.PkgName = pkg.Name
			b.SetInfo(strings.Replace(pkg.Name, "_", "-", 1), st, name)
		}
	}
}

// SetInfo parser fields structure and fill MethodsInfo for Body
func (b *Body) SetInfo(pkgName string, st *types.Struct, nameStruct string) {
	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		name := field.Name()
		fieldType := b.syncImportByType(pkgName, field.Type())

		if unicode.IsLower(rune(name[0])) || b.FlagInfo.AllFields == true {
			b.MethodsInfo = append(b.MethodsInfo, MethodInfo{
				NameMethod: toUpFirstletter(name),
				ParamName:  name,
				ParamType:  fieldType,
				VarStruct:  string(nameStruct[0]),
				NameStruct: nameStruct,
			})
		}
	}
}

// syncImportByType parser Types and create imports
// return type to field
func (b *Body) syncImportByType(pkgName string, myType types.Type) string {
	fieldType := myType.String()
	arrPKGs := strings.Split(fieldType, "/")

	name := arrPKGs[len(arrPKGs)-1]
	arrPKGs = strings.Split(name, ".")
	if len(arrPKGs) == 1 {
		return fieldType
	}

	if arrPKGs[0] != pkgName && canAddToAppend(strings.Split(fieldType, ".")[0], b.Imports) {
		b.Imports = append(b.Imports, strings.Split(fieldType, ".")[0])
		return name
	}

	return arrPKGs[1]
}

// setNameStruct Parser flags
func (f *FlagInfo) setNameStruct() {
	argName := flag.String("name", "", "Name struct")
	argsAll := flag.Bool("all", false, "If inp this arg, i get all fields for struct")
	flag.Parse()

	f.NameStruct = *argName
	f.AllFields = *argsAll
}
