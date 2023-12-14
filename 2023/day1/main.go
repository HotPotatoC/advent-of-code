package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/HotPotatoC/advent-of-code/2023/utility"
)

func toInt(c byte) int {
	return int(c - '0')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func PartOne(input []string) int {
	sum := 0
	for _, calibration := range input {
		current := 0
		for i := 0; i < len(calibration); i++ {
			if isDigit(calibration[i]) {
				current += toInt(calibration[i]) * 10
				break
			}
		}

		for i := len(calibration) - 1; i >= 0; i-- {
			if isDigit(calibration[i]) {
				current += toInt(calibration[i])
				break
			}
		}

		sum += current
	}

	return sum
}

var mapLetterToInt map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func PartTwo(input []string) int {
	sum := 0
	for _, calibration := range input {
		firstDigit := 0
		firstDigitIDX := math.MaxInt

		lastDigit := 0
		lastDigitIDX := math.MinInt

		for letter := range mapLetterToInt {
			if strings.Contains(calibration, letter) {
				letterIdx := strings.Index(calibration, letter)
				if firstDigitIDX > letterIdx {
					firstDigitIDX = min(firstDigitIDX, letterIdx)
					firstDigit = mapLetterToInt[letter]
				}
				lastLetterIdx := strings.LastIndex(calibration, letter)
				if lastDigitIDX < lastLetterIdx {
					lastDigitIDX = max(lastDigitIDX, lastLetterIdx)
					lastDigit = mapLetterToInt[letter]
				}
			}
		}

		for i := 0; i < len(calibration); i++ {
			if isDigit(calibration[i]) && i < firstDigitIDX {
				firstDigit = toInt(calibration[i])
				firstDigitIDX = i
				break
			}
		}

		for i := len(calibration) - 1; i >= 0; i-- {
			if isDigit(calibration[i]) && i > lastDigitIDX {
				lastDigit = toInt(calibration[i])
				lastDigitIDX = i
				break
			}
		}

		sum += firstDigit*10 + lastDigit
	}

	return sum
}

func main() {
	input := utility.LoadInputFile("input.in")

	partOneElapse := time.Now()
	fmt.Printf("Part one: %d\n", PartOne(input))
	fmt.Printf("Part one completed in %s\n", time.Since(partOneElapse))

	partTwoElapse := time.Now()
	fmt.Printf("Part two: %d\n", PartTwo(input))
	fmt.Printf("Part two completed in %s\n", time.Since(partTwoElapse))
}
