package rules

import (
	"testing"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
	"github.com/stretchr/testify/require"
)

var (
	walletAddr = "0xeb92E75eCb3fcB7e6Ef99008a126A3EB51Fbe512"
	sender1    = "0xdFFf55db7904095FfE5F5e850A8EaF072661Fa50"
	sender2    = "0x64D806699127ee59F0AD1578f8098BcD6e095305"
	sender3    = "0x1429430ca11986674aAB130525E2aec42AD5DC97"

	senderGroup1 = &proto.SenderGroup{
		Addresses: []string{sender1, sender2},
		Name:      "TEST_SENDER_1",
	}

	senderGroup2 = &proto.SenderGroup{
		Addresses: []string{sender2, sender3},
		Name:      "TEST_SENDER_2",
	}

	rules = []*proto.AccessControlRule{
		{
			WalletAddr:       walletAddr,
			TokenAddress:     "TOKEN_ADDR_1",
			RecipientAddress: "TO_ADDR_1",
			SenderGroup:      senderGroup1,
		}, {
			// Value Range Clause for sender - To anyone
			// Value Range should contribute to Spending Caps??
			WalletAddr:   walletAddr,
			SenderGroup:  senderGroup1,
			TokenAddress: "TOKEN_ADDR_2",
			ValueRangeClause: &proto.ValueRangeClause{
				MinVal: 100,
				MaxVal: 1000,
			},
		}, {
			// No clauses - to a specific addr
			WalletAddr:       walletAddr,
			SenderGroup:      senderGroup1,
			TokenAddress:     "TOKEN_ADDR_2",
			RecipientAddress: "TO_ADDR_2",
		},
	}
)

var (
	tx1 = &proto.SolaceTx{
		Sender:     &proto.Sender{Addr: sender1, Nonce: 0},
		ToAddr:     "TO_ADDR_1",
		TokenAddr:  "TOKEN_ADDR_1",
		Value:      101,
		WalletAddr: walletAddr,
	}
	ethSenderAddr, err = common.NewEthWalletAddressString(sender1)
)

func Test_RecipientClause(t *testing.T) {
	// Test Rule 1
	require.Nil(t, err)

	err = ValidateTx(tx1, ethSenderAddr, rules)
	require.Nil(t, err)

	// Test Rule 2
	tx1.ToAddr = walletAddr

	err = ValidateTx(tx1, ethSenderAddr, rules)
	t.Log(err)
	require.NotNil(t, err)
}

func Test_Wildcard(t *testing.T) {
	tx1.TokenAddr = "TOKEN_ADDR_2"
	tx1.ToAddr = "TO_ADDR_2"
	err := ValidateTx(tx1, ethSenderAddr, rules)
	require.Nil(t, err)
}
