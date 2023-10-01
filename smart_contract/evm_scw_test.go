package smartcontract

import (
	"testing"

	"github.com/solace-labs/skeyn/common"
)

func Test_FetchOwner(t *testing.T) {
	addr, _ := common.NewEthWalletAddressString("0xdcE2a3609308163633bd9F30d6259B1785ed88B6")
	evmScw := NewEvmScw(80001, addr)
	ownerAddr, _ := evmScw.GetOwner()
	t.Log(ownerAddr.String())
}
