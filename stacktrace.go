package corlog

import (
	"os"
	"runtime"
	"strings"
)

// getFileName ...
func getFileName() string {
	_, file, _, _ := runtime.Caller(2)
	p, _ := os.Getwd()

	return strings.TrimPrefix(file, p)
	// filename := fullpath[strings.LastIndex(fullpath, `/`)+1:]
	// return fullpath, filename
}
