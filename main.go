package main

import (
	metrices "github.com/Horiodino/terrago/monitoring"
)

func main() {
	metrices.Getinfo()
	metrices.Display()
	// cpu, err := metrices.Cpu()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cpu)
	metrices.Cpu()

}
