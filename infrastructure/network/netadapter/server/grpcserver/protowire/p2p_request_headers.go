package protowire

import (
	"github.com/SCR-NETWORK/SCR_Network/app/appmessage"
	"github.com/pkg/errors"
)

func (x *SCRpadMessage_RequestHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "SCRpadMessage_RequestBlockLocator is nil")
	}
	lowHash, err := x.RequestHeaders.LowHash.toDomain()
	if err != nil {
		return nil, err
	}

	highHash, err := x.RequestHeaders.HighHash.toDomain()
	if err != nil {
		return nil, err
	}

	return &appmessage.MsgRequestHeaders{
		LowHash:  lowHash,
		HighHash: highHash,
	}, nil
}
func (x *RequestHeadersMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestHeadersMessage is nil")
	}
	lowHash, err := x.LowHash.toDomain()
	if err != nil {
		return nil, err
	}

	highHash, err := x.HighHash.toDomain()
	if err != nil {
		return nil, err
	}

	return &appmessage.MsgRequestHeaders{
		LowHash:  lowHash,
		HighHash: highHash,
	}, nil

}

func (x *SCRpadMessage_RequestHeaders) fromAppMessage(msgRequestHeaders *appmessage.MsgRequestHeaders) error {
	x.RequestHeaders = &RequestHeadersMessage{
		LowHash:  domainHashToProto(msgRequestHeaders.LowHash),
		HighHash: domainHashToProto(msgRequestHeaders.HighHash),
	}
	return nil
}
