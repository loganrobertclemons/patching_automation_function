package structs

import (
	log "github.com/sirupsen/logrus"
)

type ApplicationListType struct {
	Applications map[string]ApplicationType `json:"applications"`
}

func (list *ApplicationListType) Init() {
	//newList := make(map[string]applicationType)
	list.Applications = make(map[string]ApplicationType)
}

func (list *ApplicationListType) Application(name string) *ApplicationType {
	if app, ok := list.Applications[name]; ok {
		return &app
	}
	app := new(ApplicationType)
	app.Init()
	list.Applications[name] = *app
	return app
}

func (list *ApplicationListType) GetApplication(name string) *ApplicationType {
	if app, ok := list.Applications[name]; ok {
		return &app
	}
	log.Error("Invalid environment: %v", name)
	return nil
}
