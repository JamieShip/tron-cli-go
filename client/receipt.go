package client

type Receipt struct {
	Result           string `json:"result"`
	EnergyUsage      int    `json:"energy_usage,omitempty"`
	EnergyUsageTotal int    `json:"energy_usage_total"`
	NetUsage         int    `json:"net_usage,omitempty"`
	EnergyFee        int    `json:"energy_fee,omitempty"`
	NetFee           int    `json:"net_fee,omitempty"`
}
