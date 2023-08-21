package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func p1() {
	var ans string
	var correct int
	var incorrect int
	fd, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	file, err := fileReader.ReadAll()
	for _, k := range file {
		fmt.Println(k[0], "= ?")
		fmt.Scanf("%s", &ans)
		if k[1] == ans {
			ans = ""
			correct += 1
		} else {
			incorrect += 1
		}
	}
	fmt.Println("Total correct answer:", correct)
	fmt.Println("Total questions:", len(file))
}

func main() {
	p1()
}
