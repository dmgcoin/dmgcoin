package protowire

import (
	"github.com/dmgcoin/dmgcoin/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DmgcoinMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DmgcoinMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *DmgcoinMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
