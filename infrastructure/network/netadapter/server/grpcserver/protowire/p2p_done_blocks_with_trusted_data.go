package protowire

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SCRpadMessage_DoneBlocksWithTrustedData) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SCRpadMessage_DoneBlocksWithTrustedData is nil")
	}
	return &appmessage.MsgDoneBlocksWithTrustedData{}, nil
}

func (x *SCRpadMessage_DoneBlocksWithTrustedData) fromAppMessage(_ *appmessage.MsgDoneBlocksWithTrustedData) error {
	return nil
}
