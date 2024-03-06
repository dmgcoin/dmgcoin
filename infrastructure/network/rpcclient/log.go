package rpcclient

import (
	"github.com/dmgcoin/dmgcoin/infrastructure/logger"
	"github.com/dmgcoin/dmgcoin/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
