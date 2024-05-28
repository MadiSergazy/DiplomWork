package private

import (
	"diplom/foundation/blockchain/state"
	"diplom/foundation/events"
	"diplom/foundation/nameservice"
	"diplom/foundation/web"
	"net/http"

	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log   *zap.SugaredLogger
	State *state.State
	NS    *nameservice.NameService
	Evts  *events.Events
}

// Routes binds all the private routes.
func Routes(app *web.App, cfg Config) {
	prv := Handlers{
		Log:   cfg.Log,
		State: cfg.State,
		NS:    cfg.NS,
	}

	const version = "v1"

	app.Handle(http.MethodPost, version, "/node/peers", prv.SubmitPeer)
	app.Handle(http.MethodGet, version, "/node/status", prv.Status)
	app.Handle(http.MethodGet, version, "/node/block/list/:from/:to", prv.BlocksByNumber)
	app.Handle(http.MethodPost, version, "/node/block/propose", prv.ProposeBlock)
	app.Handle(http.MethodPost, version, "/node/tx/submit", prv.SubmitNodeTransaction)
	app.Handle(http.MethodGet, version, "/node/tx/list", prv.Mempool)
}
