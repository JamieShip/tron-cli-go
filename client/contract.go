package client

type Contract struct {
	Parameter    Parameter `json:"parameter"`
	Type         string    `json:"type"`
	PermissionId int       `json:"Permission_id,omitempty"`
}

type Parameter struct {
	Value   Value  `json:"value"`
	TypeUrl string `json:"type_url"`
}

type Value struct {
	Amount          int64  `json:"amount,omitempty"`
	OwnerAddress    string `json:"owner_address"`
	ToAddress       string `json:"to_address,omitempty"`
	Data            string `json:"data,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	Balance         int64  `json:"balance,omitempty"`
	Resource        string `json:"resource,omitempty"`
	ReceiverAddress string `json:"receiver_address,omitempty"`
	FrozenBalance   int    `json:"frozen_balance,omitempty"`
	AssetName       string `json:"asset_name,omitempty"`
	AccountAddress  string `json:"account_address,omitempty"`
}
