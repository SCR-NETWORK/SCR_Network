package protowire

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SCRpadMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SCRpadMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *SCRpadMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
