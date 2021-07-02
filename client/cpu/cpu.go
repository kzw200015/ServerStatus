package cpu

import (
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPU() (float64, error) {
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}
