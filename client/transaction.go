package client

import "fmt"

type GetTransactionInfoByIdReq struct {
	Value string `json:"value"`
}

type Transaction struct {
	Ret        []TransactionRet   `json:"ret"`
	Signature  []string           `json:"signature"`
	TxID       string             `json:"txID"`
	RawData    TransactionRawData `json:"raw_data"`
	RawDataHex string             `json:"raw_data_hex"`
}

type TransactionRawData struct {
	Contract      []Contract `json:"contract"`
	RefBlockBytes string     `json:"ref_block_bytes"`
	RefBlockHash  string     `json:"ref_block_hash"`
	Expiration    int64      `json:"expiration"`
	Timestamp     int64      `json:"timestamp,omitempty"`
	FeeLimit      int        `json:"fee_limit,omitempty"`
}

type TransactionRet struct {
	ContractRet string `json:"contractRet"`
	Ret         string `json:"ret,omitempty"`
}

type TransactionInfo struct {
	Log                  []Log                 `json:"log,omitempty"`
	BlockNumber          uint64                `json:"blockNumber"`
	ContractResult       []string              `json:"contractResult"`
	BlockTimeStamp       int64                 `json:"blockTimeStamp"`
	Receipt              Receipt               `json:"receipt"`
	Id                   string                `json:"id"`
	ContractAddress      string                `json:"contract_address"`
	Fee                  uint64                `json:"fee,omitempty"`
	InternalTransactions []InternalTransaction `json:"internal_transactions,omitempty"`
}

type InternalTransaction struct {
	CallerAddress     string      `json:"caller_address"`
	Note              string      `json:"note"`
	TransferToAddress string      `json:"transferTo_address"`
	CallValueInfo     []CallValue `json:"callValueInfo"`
	Hash              string      `json:"hash"`
}

type CallValue struct {
	CallValue int64 `json:"callValue,omitempty"`
}

func (t *TronClient) GetTransactionInfoByID(value string) (*TransactionInfo, error) {
	var transactionInfo *TransactionInfo
	resp, err := t.c.R().SetBody(GetTransactionInfoByIdReq{
		Value: value,
	}).SetResult(&transactionInfo).Post("/wallet/gettransactioninfobyid")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get transaction by id, resp: %v", resp)
	}

	return transactionInfo, nil
}
