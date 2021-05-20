package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type oneLineResult struct {
	res []int
	date int64
}

func stringToSlice(in string) ([]int, error) {
	var err error
	a := strings.Split(in, ",")
	output := make([]int, len(a))
	for i, val := range a {
		output[i], err = strconv.Atoi(val)
		if err != nil {
			return nil, fmt.Errorf("can't convert string to interger array")
		}
	}
	return output, nil
}

func readFromCSV(fileName string) ([]oneLineResult, error) {
	var err error
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't read CSV file from disk")
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	output := make([]oneLineResult, len(csvLines))
	if err != nil {
		return nil, fmt.Errorf("can't read the CSV file, make sure the format is correct")
	}

	for index, line := range csvLines {
		output[index].res, err = stringToSlice(strings.Join(line, ";"))
	}
	return output, nil
}

func getAmountGuessPerLine(user, lotto []int) int {
	var result = 0
	for _, val := range user {
		if contains(lotto, val) {
			result += 1
		}
	}
	return result
}

func contains(in []int, s int) bool {
	for _, val := range in {
		if s == val {
			return true
		}
	}
	return false
}

func sumSlice(sli []int) int {
	result := 0
	for _, val := range sli {
		result += val
	}
	return result
}

func getMulDifferent(user, lotto []int) int {
	result := 1
	for _, val := range lotto {
		if !contains(user, val) {
			result *= val
		}
	}
	return result
}

func findMaxMap(m map[int]int) (int,int) {
	max := 0
	ind := 0
	for key, value := range m {
		if value >= max {
			max = value
			ind = key
		}
	}
	return max,ind
}
