package protowire

import (
	"github.com/dmgcoin/dmgcoin/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DmgcoinMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DmgcoinMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *DmgcoinMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
