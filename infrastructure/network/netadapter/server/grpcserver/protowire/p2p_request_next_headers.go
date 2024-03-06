package protowire

import (
	"github.com/dmgcoin/dmgcoin/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DmgcoinMessage_RequestNextHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DmgcoinMessage_RequestNextHeaders is nil")
	}
	return &appmessage.MsgRequestNextHeaders{}, nil
}

func (x *DmgcoinMessage_RequestNextHeaders) fromAppMessage(_ *appmessage.MsgRequestNextHeaders) error {
	return nil
}
