package main

import (
	"github.com/dmgcoin/dmgcoin/infrastructure/logger"
	"github.com/dmgcoin/dmgcoin/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RORG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
