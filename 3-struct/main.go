package main

import (
	"demo/bins/bins"
	"fmt"
)

func main() {
	binList := bins.NewBinList()
	binList.Add(false, "bob")

	fmt.Printf("%+v", binList)
}
