package gateway

import (
	"github.com/nbs-go/nlogger/v2"
	logOption "github.com/nbs-go/nlogger/v2/option"
)

var log nlogger.Logger

func init() {
	log = nlogger.Get().NewChild(logOption.WithNamespace("ngrpc/gateway"))
}
