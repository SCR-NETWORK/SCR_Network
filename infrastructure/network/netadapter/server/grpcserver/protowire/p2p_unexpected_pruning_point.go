package protowire

import "github.com/SCR-NETWORK/SCR_Network/app/appmessage"

func (x *SCRpadMessage_UnexpectedPruningPoint) toAppMessage() (appmessage.Message, error) {
	return &appmessage.MsgUnexpectedPruningPoint{}, nil
}

func (x *SCRpadMessage_UnexpectedPruningPoint) fromAppMessage(_ *appmessage.MsgUnexpectedPruningPoint) error {
	return nil
}
