package host

import "github.com/shirou/gopsutil/v3/host"

func GetUptime() (uint64, error) {
	uptime, err := host.Uptime()
	if err != nil {
		return 0, err
	}
	return uptime, nil
}
