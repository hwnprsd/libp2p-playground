package smartcontract

import (
	"testing"

	"github.com/solace-labs/skeyn/common"
	"github.com/stretchr/testify/require"
)

func Test_FetchOwnerEVM(t *testing.T) {
	ownerAddr, err := TestEvmScw.GetOwner()
	if err != nil {
		t.Error(err)
		return
	}
	expectedAddr, _ := common.NewEthWalletAddressString("0x5F4Aef7d8AcaA89140aB928539183985958699F2")
	require.Equal(t, ownerAddr, expectedAddr)
}

// TODO: Add more test cases
