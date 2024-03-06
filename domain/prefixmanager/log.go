package prefixmanager

import (
	"github.com/dmgcoin/dmgcoin/infrastructure/logger"
	"github.com/dmgcoin/dmgcoin/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
