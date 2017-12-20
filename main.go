package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hey")
	f, err := os.Open("./problems.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}

}
