package mem

import (
	"github.com/kzw200015/ServerStatus/types"
	"github.com/shirou/gopsutil/v3/mem"
)

func GetMemStatus() (types.MemStatus, error) {

	m, err := mem.VirtualMemory()
	if err != nil {
		return types.MemStatus{}, err
	}

	return types.MemStatus{
		TotalMem: m.Total,
		UsedMem:  m.Used,
		Percent:  m.UsedPercent,
	}, nil
}

func GetSwapStatus() (types.SwapStatus, error) {

	m, err := mem.SwapMemory()
	if err != nil {
		return types.SwapStatus{}, err
	}

	return types.SwapStatus{
		TotalMem: m.Total,
		UsedMem:  m.Used,
		Percent:  m.UsedPercent,
	}, nil
}
