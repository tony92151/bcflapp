package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers bcflapp-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/bcflapp/apply", listApplyHandler(cliCtx, "bcflapp")).Methods("GET")
	r.HandleFunc("/bcflapp/apply", createApplyHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/bcflapp/joblist", listJoblistHandler(cliCtx, "bcflapp")).Methods("GET")
	r.HandleFunc("/bcflapp/joblist", createJoblistHandler(cliCtx)).Methods("POST")
}
