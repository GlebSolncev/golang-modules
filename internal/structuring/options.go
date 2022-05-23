package structuring

// Body for template
type Body struct {
	FileInfo
	MethodsInfo
	FlagInfo
}

type FlagInfo struct {
	NameStruct string
	AllFields  bool
}

// FileInfo info about struct - package name, import list
type FileInfo struct {
	NameStructure string
	PkgName       string   // package name
	Imports       []string // list imports
}

// MethodsInfo scope methods
type MethodsInfo []MethodInfo

// MethodInfo Info about methods
// NameMethod entity for func Set/Get{NameMethod}
// Set/Get{NameMethod}(ParamName ParamType) ParamType {...}
// For setter VarStruct. ParamName = ParamName
// For getter return VarStruct. ParamName
type MethodInfo struct {
	NameMethod string
	ParamName  string
	ParamType  string
	NameStruct string
	VarStruct  string
}
