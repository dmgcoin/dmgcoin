package protowire

import (
	"github.com/dmgcoin/dmgcoin/app/appmessage"
	"github.com/pkg/errors"
)

func (x *DmgcoinMessage_RequestPruningPointAndItsAnticone) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "DmgcoinMessage_RequestPruningPointAndItsAnticone is nil")
	}
	return &appmessage.MsgRequestPruningPointAndItsAnticone{}, nil
}

func (x *DmgcoinMessage_RequestPruningPointAndItsAnticone) fromAppMessage(_ *appmessage.MsgRequestPruningPointAndItsAnticone) error {
	return nil
}
