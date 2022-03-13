package fileUtils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)


func removePidFile(path string) {
	if err := os.Remove(path); err != nil {
		log.Printf("Error: %s\n", err)
	}
}

func readPidFile(path string) ([]byte, error) {
	fmt.Println("Entering readPidFile func. About to read file.")

	fileContents, err := ioutil.ReadFile(path) // cauting with RAM
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read file")
	}

	fmt.Println("About to remove pid-file.")
	removePidFile(path)

	fmt.Println("Exiting readPidFile func. About to return file contents.")
	return fileContents, nil
}

func GetServerPid(path string) (int, error) {
	fmt.Println("Entering ReadPidFile func. About to call readPidFile.")

	pifFileContents, err := readPidFile(path)
	if err != nil {
		return 0, err
	}
	fmt.Println("Got pid-file contents bytes.")

	pidFileContentsStr := string(pifFileContents)
	pidFileContentsTrimmedStr := strings.TrimSpace(pidFileContentsStr)
	fmt.Println("Converted pid-file contents to string. About to get PID")

	serverPid, err := strconv.Atoi(pidFileContentsTrimmedStr)
	if err != nil {
		return 0, errors.Wrap(err, "Faild to convert pid-file contents to integer.")
	}

	fmt.Println("Got PID integer. Exiting GetFilePid, return server's PID")
	return serverPid, nil
}
