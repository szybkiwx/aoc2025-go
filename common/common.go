package common

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)


func resolvePath(fileName string) (string, error) {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		return "", fmt.Errorf("could not determine caller location")
	}

	callerDir := filepath.Dir(file)
	return filepath.Join(callerDir, fileName), nil
}


func ReadWholeFile(fileName string) ([]byte, error) {
	fullPath, err := resolvePath(fileName)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(fullPath)
}


func ReadLines(fileName string) ([]string, error) {
	fullPath, err := resolvePath(fileName)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
