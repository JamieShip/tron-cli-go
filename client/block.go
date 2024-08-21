package client

import (
	"fmt"
)

type GetBlockByNumberReq struct {
	Num uint64 `json:"num"`
}

type GetBlockReq struct {
	Detail  bool   `json:"detail,omitempty"`
	IdOrNum string `json:"id_or_num,omitempty"`
}

type Block struct {
	BlockID     string      `json:"blockID"`
	BlockHeader BlockHeader `json:"block_header"`

	Transactions []Transaction `json:"transactions"`
}

type BlockHeader struct {
	RawData          BlockHeaderRawdata `json:"raw_data"`
	WitnessSignature string             `json:"witness_signature"`
}

type BlockHeaderRawdata struct {
	Number         int    `json:"number,omitempty"`
	TxTrieRoot     string `json:"txTrieRoot,omitempty"`
	WitnessAddress string `json:"witness_address,omitempty"`
	ParentHash     string `json:"parentHash,omitempty"`
	Version        int    `json:"version,omitempty"`
	Timestamp      int64  `json:"timestamp,omitempty"`
}

func (t *TronClient) GetNowBlock() (*Block, error) {
	var block *Block
	resp, err := t.c.R().SetResult(&block).Post("/wallet/getnowblock")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to request getnowblock, resp: %v", resp)
	}

	return block, nil
}

func (t *TronClient) GetBlock(number uint64, detail bool) (*Block, error) {
	var block *Block
	resp, err := t.c.R().SetBody(GetBlockReq{
		IdOrNum: fmt.Sprintf("%d", number),
		Detail:  detail,
	}).SetResult(&block).Post("/wallet/getblock")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to request getblockbynum, resp: %v", resp)
	}

	return block, nil
}

func (t *TronClient) GetBlockByNum(number uint64) (*Block, error) {
	var block *Block
	resp, err := t.c.R().SetBody(GetBlockByNumberReq{
		Num: number,
	}).SetResult(&block).Post("/wallet/getblockbynum")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to request getblockbynum, resp: %v", resp)
	}

	return block, nil
}
