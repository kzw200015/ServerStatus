package types

type Status struct {
	CPU    float64    `json:"cpu,omitempty"`
	Mem    MemStatus  `json:"mem,omitempty"`
	Swap   SwapStatus `json:"swap,omitempty"`
	Net    NetStatus  `json:"net,omitempty"`
	Uptime uint64     `json:"uptime,omitempty"`
}

type MemStatus struct {
	TotalMem uint64  `json:"total_mem,omitempty"`
	UsedMem  uint64  `json:"used_mem,omitempty"`
	Percent  float64 `json:"percent,omitempty"`
}

type SwapStatus struct {
	TotalMem uint64  `json:"total_mem,omitempty"`
	UsedMem  uint64  `json:"used_mem,omitempty"`
	Percent  float64 `json:"percent,omitempty"`
}

type NetStatus struct {
	RX uint64 `json:"rx,omitempty"`
	TX uint64 `json:"tx,omitempty"`
}
