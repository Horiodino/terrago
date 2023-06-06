package monitoring

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"cpu_tem "github.com/Horiodino/terrago/monitoring/temprature"

)

func GetCPUTemperature() (float64, error) {
	data, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	temperatureStr := strings.TrimSpace(string(data))
	temperature, err := strconv.ParseFloat(temperatureStr, 64)
	if err != nil {
		return 0, err
	}

	temperature /= 1000.0 /
	return temperature, nil
}


