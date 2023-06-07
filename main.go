package main

import (
	alerts "github.com/Horiodino/terrago/Alerts"
)

func main() {
	// metrices.Getinfo()
	// metrices.Display()
	// cpu, err := metrices.Cpu()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cpu)
	// metrices.Cpu()
	// container.Containermatricesinfo()
	alerts.DeploumentsFailure()

}
