package protowire

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SCRpadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SCRpadMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *SCRpadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
