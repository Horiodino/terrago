package main

import (
	"fmt"

	metrices "github.com/Horiodino/terrago/monitoring"
)

func main() {
	metrices.Getinfo()
	metrices.Display()

	datail := metrices.cpu()

	for _, v := range datail {
		fmt.Println(v)
	}

}
