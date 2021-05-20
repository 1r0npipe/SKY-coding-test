package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	fileName := flag.String("fileResult", "lottoResults.csv", "provide the file with lotto result")
	pNumbers := flag.String("pNumbers", "", "provide the number of player comma separated")
	date := flag.String("date", "", "provide the date of draw to check")

	flag.Parse()
	dateUser, err := strconv.Atoi(*date)
	if err != nil {
		log.Fatal("can't transform user input to date format")
	}

	userNumbers, err := stringToSlice(*pNumbers)
	if err != nil {
		log.Fatal("I can't read your ouput correctly from pNumbers key")
	}
	lottoResult, err := readFromCSV(*fileName)
	if err != nil {
		log.Fatal("error occurs when read file: ", err)
	}
	guessedSame := make(map[int]int, len(lottoResult))

	for index, lottoLine := range lottoResult {
		if userDate >= drawDate {
			guessedNumbers := getAmountGuessPerLine(userNumbers, lottoLine.res)
			guessedSame[index] = guessedNumbers
		}
	}
	winRes := 0
	max, ind := findMaxMap(guessedSame)
	for inside of map{
		guessedNumbers := getAmountGuessPerLine(userNumbers, lottoResult[ind].res)
		switch{
	case max < 3:
		winRes = sumSlice(lottoResult[ind].res)
	case max >= 3 && max <= 5:
		winRes = guessedNumbers*500 + getMulDifferent(userNumbers, lottoResult[ind].res)
	case max == 6:
		winRes = sumSlice(userNumbers) * 10000
	}
		//fmt.Println(winRes)
	}
	//fmt.Println(winRes)

}

