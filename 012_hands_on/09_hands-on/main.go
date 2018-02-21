package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

var data [2]string

func main() {
	r := csv.NewReader(strings.NewReader(`"a", "b"`))
	data[0] = "a"
	fmt.Println(r.Read())
}
