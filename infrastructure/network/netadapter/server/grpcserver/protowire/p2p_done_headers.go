package protowire

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SCRpadMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SCRpadMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *SCRpadMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
