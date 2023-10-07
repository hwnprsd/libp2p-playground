package rules

import (
	"testing"

	"github.com/solace-labs/skeyn/common"
	"github.com/solace-labs/skeyn/proto"
)

var (
	walletAddr = "0xeb92E75eCb3fcB7e6Ef99008a126A3EB51Fbe512"
	sender1    = "0xdFFf55db7904095FfE5F5e850A8EaF072661Fa50"
	sender2    = "0x64D806699127ee59F0AD1578f8098BcD6e095305"

	tx = &proto.SolaceTx{
		Sender:     &proto.Sender{Addr: sender1, Nonce: 0},
		ToAddr:     "TO_ADDR_1",
		TokenAddr:  "TOKEN_ADDR_1",
		Value:      1000,
		WalletAddr: walletAddr,
	}

	rules = []*proto.AccessControlRule{
		{
			WalletAddr: walletAddr,
			SenderGroup: &proto.SenderGroup{
				Addresses: []string{sender1, sender2},
				Name:      "TEST_SENDER_1",
			},
			TokenAddress:     "TOKEN_ADDR_1",
			RecipientAddress: "TO_ADDR_1",
		}, {
			WalletAddr: walletAddr,
			SenderGroup: &proto.SenderGroup{
				Addresses: []string{sender1, sender2},
				Name:      "TEST_SENDER_2",
			},
			TokenAddress:     "TOKEN_ADDR_2",
			RecipientAddress: "TO_ADDR_1",
		},
	}
)

func Test_GetRules(t *testing.T) {
	t.Log(rules[0].Ids())
	ethSenderAddr, err := common.NewEthWalletAddressString(sender1)
	if err != nil {
		t.Error(err)
	}
	t.Log("//", sender1, "//", ethSenderAddr.String())
	r1, r2, r3 := GetRulesForSender(tx, ethSenderAddr, rules)
	t.Log(len(r1), len(r2), len(r3))
}
