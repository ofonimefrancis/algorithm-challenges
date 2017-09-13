package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	aliceData, _ := reader.ReadString('\n')
	bobsData, _ := reader.ReadString('\n')
	aliceSlice := strings.Fields(aliceData)
	bobsSlice := strings.Fields(bobsData)

	aliceScores := 0
	bobsScores := 0

	for index, val := range aliceSlice {
		bobsInt, _ := strconv.Atoi(bobsSlice[index])
		aliceInt, _ := strconv.Atoi(val)
		if aliceInt > bobsInt {
			aliceScores += 1
		} else if bobsInt > aliceInt {
			bobsScores += 1
		} else if bobsInt == aliceInt {
			bobsScores += 0
			aliceScores += 0
		}
	}
	fmt.Printf("%d %d", aliceScores, bobsScores)
}
