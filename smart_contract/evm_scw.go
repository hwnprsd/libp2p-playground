package smartcontract

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/solace-labs/skeyn/smart_contract/solace_account"
)

const (
	errInvalidRpcUrl        = "Invalid RPC Url // %d"
	errDialingEthClient     = "Error Dialing Eth Client // %e"
	errConnectingToScw      = "Error connecting to SolaceAccount // %e"
	errFetchingOwnerFromScw = "Error fetching owner from SolaceAccount // %e"
)

type EvmScw struct {
	chainID    int
	walletAddr common.Addr
}

func NewEvmScw(chainId int, walletAddr common.Addr) EvmScw {
	return EvmScw{chainId, walletAddr}
}

func (e EvmScw) validateSetup(sig []byte) (bool, error) {
	return true, nil
}

func (e EvmScw) validateRuleDeletion(rule *proto.Rule, sig []byte, sender common.Addr) (bool, error) {
	return true, nil
}

func (e EvmScw) validateRuleAddition(rule *proto.Rule, sig []byte, sender common.Addr) (bool, error) {
	rpcUrl, exists := RPCUrls[e.chainID]

	if !exists {
		return false, fmt.Errorf(errInvalidRpcUrl, e.chainID)
	}
	// RPC
	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		return false, fmt.Errorf(errInvalidRpcUrl, err)
	}

	// Client
	solaceAcc, err := solaceaccount.NewSolaceAccount(ethcommon.BytesToAddress(e.walletAddr.Bytes()), client)
	if err != nil {
		return false, fmt.Errorf(errConnectingToScw, err)
	}

	ownerAddr, err := solaceAcc.Owner(&bind.CallOpts{})
	if err != nil {
		return false, fmt.Errorf(errFetchingOwnerFromScw, err)
	}

	// ECRecover the sig and verify that the owner signed it

	return res == 0, nil
}
