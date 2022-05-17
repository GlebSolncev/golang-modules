package path

import (
	"path/filepath"
	"runtime"
)

type HelloWorld struct {
	Id int
}

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetBasePath(filename string) string {
	var (
		p  = "/../.."
		fc = filename[0:1]
	)
	if fc != "/" {
		p += "/"
	}

	return basePath + p + filename
}
