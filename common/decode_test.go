package common_test

import (
	"testing"

	"github.com/JamieShip/tron-cli-go/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestOnDecode(t *testing.T) {
	// Transaction Detail:
	// https://tronscan.org/#/transaction/589646c3d412db88e37be52124b3032baaa74cfda5cedfaba844c2a13ed5c926/event-logs
	t.Run("test on decode tron address encode", func(t *testing.T) {
		addr, err := common.FromHexAddress("41f6d82beb987baae2f0cdc7bdfa652be660dadac8")
		require.NoError(t, err)
		require.Equal(t, "TYUQC7VrLR4vZ6pe6y1sJMEx91aMpLKJP1", addr)

		usdtAddr, err := common.FromHexAddress("41a614f803b6fd780986a42c78ec9c7f77e6ded13c")
		require.NoError(t, err)
		require.Equal(t, "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", usdtAddr)
	})

	t.Run("test on decode tron address decode", func(t *testing.T) {
		hex := common.ToHexAddress("TQwWvvNF224otczuts1zHTTmcpEnNr7dYC")

		// 000000000000000000000000a438055afd8e6e6b531d46c15cca40968b64b390
		require.Equal(t, "41a438055afd8e6e6b531d46c15cca40968b64b390", hex)
	})

	t.Run("test on decode tron event log data from ", func(t *testing.T) {
		decodeNum, err := hexutil.DecodeBig("0x15a0")
		require.NoError(t, err)
		require.Equal(t, "5536", decodeNum.String())
	})
}
