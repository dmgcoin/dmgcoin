package standalone

import (
	"github.com/dmgcoin/dmgcoin/infrastructure/logger"
	"github.com/dmgcoin/dmgcoin/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
