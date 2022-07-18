package services

import (
	"bufio"
	"os"
)

func StoreFile(threads []string, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	datawriter := bufio.NewWriter(file)
	defer datawriter.Flush()
	for _, data := range threads {
		_, _ = datawriter.WriteString(data + "\n")
	}

	return nil
}

func RetrieveFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
