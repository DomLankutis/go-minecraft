package utils

import (
	"bufio"
	"os"
)

func ReadFile(path string) (string, error) {
	src := ""
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		src +="\n" + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	src += "\x00"
	return src, nil
}