package protowire

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SCRpadMessage_RequestNextHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SCRpadMessage_RequestNextHeaders is nil")
	}
	return &appmessage.MsgRequestNextHeaders{}, nil
}

func (x *SCRpadMessage_RequestNextHeaders) fromAppMessage(_ *appmessage.MsgRequestNextHeaders) error {
	return nil
}
