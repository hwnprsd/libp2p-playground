package smartcontract

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/smart_contract/solace_account"
	"github.com/solace-labs/skeyn/utils"
)

const (
	errInvalidRpcUrl        = "Invalid RPC Url // %d"
	errDialingEthClient     = "Error Dialing Eth Client // %e"
	errConnectingToScw      = "Error connecting to SolaceAccount // %e"
	errFetchingOwnerFromScw = "Error fetching owner from SolaceAccount // %e"
	errInvalidSignature     = "Invalid Signature - Not signed by SCW Owner // Owner - %s // %e"
)

var (
	addr, _    = common.NewEthWalletAddressString("0x6eDbBd37699FE0e7F24E7956E95D1630d9DF7971")
	TestEvmScw = NewEvmScw(4337, addr)
)

type EvmScw struct {
	chainID    int
	walletAddr common.Addr
}

func NewEvmScw(chainId int, walletAddr common.Addr) EvmScw {
	return EvmScw{chainId, walletAddr}
}

func (e EvmScw) GetOwner() (common.Addr, error) {
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

	if ownerAddr.Cmp(ethcommon.HexToAddress("0x00")) == 0 {
		return common.ZeroAddr(), fmt.Errorf("Owner not found for smart account")
	}

	return common.NewEthWalletAddress(ownerAddr), nil
}

func (e EvmScw) ValidateSetup(sig []byte) (bool, error) {
	return true, nil
}

func (e EvmScw) ValidateRuleDeletion(rule []byte, sig []byte, sender common.Addr) (bool, error) {
	return true, nil
}

// Verify that the role addition has been authorized by the owner of SCW
func (e EvmScw) ValidateRuleAddition(rule []byte, sig []byte, sender common.Addr) (bool, error) {
	ownerAddr, err := e.GetOwner()
	if err != nil {
		return false, err
	}
	// Verify that the owner has signed the rule
	err = utils.VerifyEthSignature(rule, sig, ownerAddr)
	if err != nil {
		return false, fmt.Errorf(errInvalidSignature, ownerAddr.String(), err)
	}

	return true, nil
}
