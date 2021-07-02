package client

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/kzw200015/ServerStatus/client/config"
	"github.com/kzw200015/ServerStatus/client/cpu"
	"github.com/kzw200015/ServerStatus/client/host"
	"github.com/kzw200015/ServerStatus/client/mem"
	"github.com/kzw200015/ServerStatus/client/net"
	"github.com/kzw200015/ServerStatus/types"
)

var client = resty.New()

func Run(path string) {
	config.Parse(path)

	for {
		percent, err := cpu.GetCPU()
		if err != nil {
			log.Println(err)
			continue
		}

		uptime, err := host.GetUptime()
		if err != nil {
			log.Println(err)
			continue
		}

		memStatus, err := mem.GetMemStatus()
		if err != nil {
			log.Println(err)
			continue
		}

		swapStatus, err := mem.GetSwapStatus()
		if err != nil {
			log.Println(err)
			continue
		}

		netStatus, err := net.GetNetStatus()
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = client.SetBasicAuth(config.Get().Id, config.Get().Token).R().SetBody(types.Status{
			CPU:    percent,
			Mem:    memStatus,
			Swap:   swapStatus,
			Net:    netStatus,
			Uptime: uptime,
		}).Post(config.Get().ServerUrl + "/api/nodes")

		if err != nil {
			log.Println(err)
		}

		time.Sleep(3 * time.Second)
	}
}
