package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parseSignaturesFile(signatureFilePath string) (map[string]string, error) {
	fileToReadFrom, err := os.Open(signatureFilePath)
	if err != nil {
		return nil, err
	}
	defer fileToReadFrom.Close()

	signaturesVsFilePaths := make(map[string]string)
	scanner := bufio.NewScanner(fileToReadFrom)

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		currentLineText := scanner.Text()
		currentLineTextFields := strings.Fields(currentLineText)
		if len(currentLineTextFields) != len([]string{"signature", "path"}) {
			errInfo := fmt.Errorf("Error parsing line %d of file %s\n", lineNumber, signatureFilePath)
			return nil, errInfo
		}
		signature := currentLineTextFields[0]
		filePath := currentLineTextFields[1]
		signaturesVsFilePaths[signature] = filePath
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return signaturesVsFilePaths, nil
}

func calculateMD5SignatureForFile(path string) (string, error) {
	fileOpened, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fileOpened.Close()

	fileHashContainer := md5.New()
	if _, err := io.Copy(fileHashContainer, fileOpened); err != nil {
		return "", err
	}

	fileHashSum := fileHashContainer.Sum(nil)
	return fmt.Sprintf("%x", fileHashSum), nil
}

type hashCheckResult struct {
	path  string
	match bool
	err   error
}

func MD5Worker(filePath string, signatureClaimed string, output chan *hashCheckResult) {
	resultObject := &hashCheckResult{path: filePath}
	calculatedSignatue, err := calculateMD5SignatureForFile(filePath)
	if err != nil {
		resultObject.err = err
		output <- resultObject
		return
	}

	resultObject.match = (calculatedSignatue == signatureClaimed)
	output <- resultObject
}

func main() {
	signaturesVsFilePaths, err := parseSignaturesFile("data/md5sum.txt")
	if err != nil {
		log.Fatalf("Failed to parse signatures file: %s", err)
	}

	output := make(chan *hashCheckResult)

	for signatureClaimed, filePath := range signaturesVsFilePaths {
		go MD5Worker(filePath, signatureClaimed, output)
	}

	hashCheckSuccess := true
	for range signaturesVsFilePaths {
		checkResult := <-output
		switch {
		case checkResult.err != nil:
			fmt.Printf("%s - error: %s\n", checkResult.path, checkResult.err)
			hashCheckSuccess = false
		case !checkResult.match:
			fmt.Printf("%s - signature mismatch\n", checkResult.path)
			hashCheckSuccess = false
		}
	}

	if !hashCheckSuccess {
		os.Exit(1)
	}
}
