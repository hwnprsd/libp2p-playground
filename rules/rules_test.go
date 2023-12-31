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
			WalletAddr:    walletAddr,
			TokenAddr:     "TOKEN_ADDR_1",
			RecipientAddr: "TO_ADDR_1",
			SenderGroup:   senderGroup1,
		}, {
			// Value Range Clause for sender - To anyone
			// Value Range should contribute to Spending Caps??
			WalletAddr:  walletAddr,
			SenderGroup: senderGroup1,
			TokenAddr:   "TOKEN_ADDR_2",
			ValueRangeClause: &proto.ValueRangeClause{
				MinVal: 100,
				MaxVal: 1000,
			},
		}, {
			// No clauses - to a specific addr
			WalletAddr:    walletAddr,
			SenderGroup:   senderGroup1,
			TokenAddr:     "TOKEN_ADDR_2",
			RecipientAddr: "TO_ADDR_2",
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

	_, err = ValidateTx(tx1, ethSenderAddr, rules)
	require.Nil(t, err)

	// Test Rule 2
	tx1.ToAddr = walletAddr

	_, err = ValidateTx(tx1, ethSenderAddr, rules)
	require.NotNil(t, err)
}

func Test_Multirules(t *testing.T) {
	tx1.TokenAddr = "TOKEN_ADDR_2"
	tx1.ToAddr = "TO_ADDR_2"
	tx1.Value = 99999
	_, err := ValidateTx(tx1, ethSenderAddr, rules)
	require.Nil(t, err)

	tx1.ToAddr = walletAddr
	tx1.Value = 120
	_, err = ValidateTx(tx1, ethSenderAddr, rules)
	require.Nil(t, err)

	tx1.Value = 99999
	_, err = ValidateTx(tx1, ethSenderAddr, rules)
	require.NotNil(t, err)
}

func Test_RuleAddition(t *testing.T) {
	newRule := &proto.AccessControlRule{
		WalletAddr:  walletAddr,
		TokenAddr:   "TOKEN_ADDR_1",
		SenderGroup: senderGroup1,
		ValueRangeClause: &proto.ValueRangeClause{
			MinVal: 100,
			MaxVal: 10000,
		},
	}
	err := ValidateRuleAddition(rules, newRule)
	require.Nil(t, err)

	// Should Fail
	newRule.RecipientAddr = "TO_ADDR_1"
	err = ValidateRuleAddition(rules, newRule)
	require.NotNil(t, err)
}
