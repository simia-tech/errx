package errx

import (
	"runtime"
	"strings"
)

var prefixSize int
var goPath string

func init() {
	_, file, _, ok := runtime.Caller(0)
	if file == "?" {
		return
	}
	if ok {
		size := len(file)
		suffix := len("code.posteo.de/common/errx/path.go")
		goPath = file[:size-suffix]
		prefixSize = len(goPath)
	}
}

func trimGoPath(filename string) string {
	if strings.HasPrefix(filename, goPath) {
		return filename[prefixSize:]
	}
	return filename
}
