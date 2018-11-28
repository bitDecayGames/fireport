package files

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const gameLogsDir = "fireport_logs"

// GetLogFile returns a file in the log directory for the current OS
func GetLogFile(logName string) (*os.File, error) {
	datestamp := time.Now().Format("2006-01-02_15.04.05")
	fileName := strings.Join([]string{datestamp, logName}, "_")
	if !strings.HasSuffix(fileName, ".log") {
		fileName = fmt.Sprintf("%v.log", fileName)
	}
	var path string
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		usr, err := user.Current()
		if err != nil {
			// fallback in case we can't get the user
			path = filepath.Join(".", fileName)
		} else {
			err = os.MkdirAll(filepath.Join(usr.HomeDir, gameLogsDir), 0755)
			if err != nil {
				fmt.Println("Failed to make log directory: ", err)
			}
			path = filepath.Join(usr.HomeDir, gameLogsDir, fileName)
		}
	} else {
		dirPath := filepath.Join("/var", "log", gameLogsDir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				fmt.Println("Failed to make nix log directory: ", err)
			}
		}
		path = filepath.Join(dirPath, fileName)
	}

	return os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
}
