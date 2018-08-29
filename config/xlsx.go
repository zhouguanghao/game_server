package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var front1 [2048]int
	var front2 [2048]int
	var front3 [2048]int
	var front4 [2048]int
	var front5 [2048]int
	var late1 [2048]int
	var late2 [2048]int
	xlsx, err := excelize.OpenFile("./data.xlsx")
	if err != nil {
		panic(err)
	}
	rows := xlsx.GetRows("data")
	for index, row := range rows {
		if index < 2 {
			continue
		}

		dat1, _ := strconv.Atoi(row[1])
		front1[index-2] = dat1
		dat2, _ := strconv.Atoi(row[2])
		front2[index-2] = dat2
		dat3, _ := strconv.Atoi(row[3])
		front3[index-2] = dat3
		dat4, _ := strconv.Atoi(row[4])
		front4[index-2] = dat4
		dat5, _ := strconv.Atoi(row[5])
		front5[index-2] = dat5

		dat6, _ := strconv.Atoi(row[6])
		late1[index-2] = dat6
		dat7, _ := strconv.Atoi(row[7])
		late2[index-2] = dat7
	}
	fmt.Println(front1)
}
