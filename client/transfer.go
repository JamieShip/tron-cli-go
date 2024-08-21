package client

type TRC20TransferModel struct {
	BlockNumber uint64 `json:"block_number"`
	TxId        string `json:"tx_id"`
	LogIndex    uint64 `json:"log_index"`
	Timestamp   int64  `json:"timestamp"`
	Address     string `json:"address"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
	Value       string `json:"value"`
}
