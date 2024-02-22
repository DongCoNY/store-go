package appmodule

import (
	"cosmossdk.io/log"
	"github.com/DongCoNY/store-go/core/branch"
	"github.com/DongCoNY/store-go/core/event"
	"github.com/DongCoNY/store-go/core/gas"
	"github.com/DongCoNY/store-go/core/header"
	"github.com/DongCoNY/store-go/core/store"
)

// Environment is used to get all services to their respective module
type Environment struct {
	BranchService   branch.Service
	EventService    event.Service
	GasService      gas.Service
	HeaderService   header.Service
	KVStoreService  store.KVStoreService
	MemStoreService store.MemoryStoreService
	Logger          log.Logger
}
