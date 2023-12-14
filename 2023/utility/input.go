package utility

import (
	"bufio"
	"log"
	"os"
)

func ByteToInt(b byte) int {
	return int(b - '0')
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func LoadInputFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
