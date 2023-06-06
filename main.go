package main

import (
	container "github.com/Horiodino/terrago/monitoring/container_metrices"
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
	container.ContainerInfo()

}
