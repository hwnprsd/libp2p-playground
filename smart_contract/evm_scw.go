package smartcontract

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/smart_contract/solace_account"
)

const (
	errInvalidRpcUrl        = "Invalid RPC Url // %d"
	errDialingEthClient     = "Error Dialing Eth Client // %e"
	errConnectingToScw      = "Error connecting to SolaceAccount // %e"
	errFetchingOwnerFromScw = "Error fetching owner from SolaceAccount // %e"
)

var (
	RPCUrls = map[int]string{
		80001: "https://rpc.ankr.com/polygon_mumbai",
	}
)

type EvmScw struct {
	chainID    int
	walletAddr common.WalletAddress
}

func NewEvmScw(chainId int, walletAddr common.WalletAddress) EvmScw {
	return EvmScw{chainId, walletAddr}
}

func (e EvmScw) GetOwner() (common.WalletAddress, error) {
	rpcUrl, exists := RPCUrls[e.chainID]

	if !exists {
		return common.ZeroAddr(), fmt.Errorf(errInvalidRpcUrl, e.chainID)
	}
	// RPC
	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		return common.ZeroAddr(), fmt.Errorf(errInvalidRpcUrl, err)
	}

	// Client
	solaceAcc, err := solaceaccount.NewSolaceAccount(ethcommon.BytesToAddress(e.walletAddr.Bytes()), client)
	if err != nil {
		return common.ZeroAddr(), fmt.Errorf(errConnectingToScw, err)
	}

	ownerAddr, err := solaceAcc.Owner(&bind.CallOpts{})
	if err != nil {
		return common.ZeroAddr(), fmt.Errorf(errFetchingOwnerFromScw, err)
	}

	return common.NewEthWalletAddress(ownerAddr), nil
}
