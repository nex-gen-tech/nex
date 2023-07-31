package nex

import (
	"github.com/nex-gen-tech/nex/context"
	"github.com/nex-gen-tech/nex/router"
)

type (
	Context     = context.Context
	Router      = router.Router
	HandlerFunc = router.HandlerFunc
	RouterGroup = router.RouterGroup
)

// New - Create a new router
func New() *Router {
	return router.NewRouter()
}
