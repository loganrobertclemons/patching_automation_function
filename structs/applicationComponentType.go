package structs

import (
	"github.com/ChrisHirsch/puppetdb-client-go"
	log "github.com/sirupsen/logrus"
)

type ApplicationComponentType struct {
	PatchingProcedure string                `json:"patching_procedure"`
	Servers           map[string]ServerType `json:"servers"`
}

func (component *ApplicationComponentType) Init() {
	component.Servers = make(map[string]ServerType)
}

// server: return a server object or create a new/blank one
func (component *ApplicationComponentType) server(name string) *ServerType {
	if result, ok := component.Servers[name]; ok {
		return &result
	}
	log.Info("Creating a new empty server object... This is probably not what you wanted!")
	newServer := new(ServerType)
	newServer.Init()
	component.Servers[name] = *newServer
	return newServer
}

func (component *ApplicationComponentType) GetServer(name string) *ServerType {
	if result, ok := component.Servers[name]; ok {
		return &result
	}
	log.Error("Invalid component: %v", name)
	return nil
}

func (Component *ApplicationComponentType) AddServer(Server *puppetdb.Inventory) {
	if _, ok := Component.Servers["test"]; ok {
		panic("Tried to create a new server that already exists")
	}
	newServer := new(ServerType)
	newServer.Init()
	newServer.parseInventoryResult(Server)
	Component.Servers[Server.Certname] = *newServer
}
