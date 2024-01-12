package flowcontext

import (
	"github.com/SCR-NETWORK/SCR_Network/infrastructure/network/addressmanager"
)

// AddressManager returns the address manager associated to the flow context.
func (f *FlowContext) AddressManager() *addressmanager.AddressManager {
	return f.addressManager
}
