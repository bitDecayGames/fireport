package files

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GetLogFile returns a file in the log directory for the current OS
func GetLogFile(logName string) (*os.File, error) {
	datestamp := time.Now().Format("2006-01-02_15.04.05")
	fileName := strings.Join([]string{datestamp, logName}, "_")
	if !strings.HasSuffix(fileName, ".log") {
		fileName = fmt.Sprintf("%v.log", fileName)
	}
	var path string
	if runtime.GOOS == "windows" {
		path = filepath.Join(".", fileName)
	} else {
		path = filepath.Join("var", "log", fileName)
	}

	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}
