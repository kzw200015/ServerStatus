package net

import (
	"log"
	"strings"
	"time"

	"github.com/kzw200015/ServerStatus/types"
	"github.com/shirou/gopsutil/v3/net"
)

var rxSpeed uint64
var txSpeed uint64

func GetNetStatus() (types.NetStatus, error) {
	return types.NetStatus{
		RX: rxSpeed,
		TX: txSpeed,
	}, nil
}

func getRxAndTx() (uint64, uint64, error) {
	stats, err := net.IOCounters(true)
	if err != nil {
		return 0, 0, err
	}
	var rx uint64
	var tx uint64
	for _, stat := range stats {
		if stat.Name == "lo" && strings.HasPrefix(stat.Name, "docker") && strings.HasPrefix(stat.Name, "br-") && strings.HasPrefix(stat.Name, "veth") {
			continue
		}
		rx += stat.BytesRecv
		tx += stat.BytesSent
	}
	return rx, tx, nil
}

func init() {
	go func() {
		for {
			rx1, tx1, err := getRxAndTx()
			if err != nil {
				log.Println(err)
				continue
			}
			time.Sleep(1 * time.Second)

			rx2, tx2, err := getRxAndTx()
			if err != nil {
				log.Println(err)
				continue
			}

			rxSpeed = rx2 - rx1
			txSpeed = tx2 - tx1
		}
	}()
}
