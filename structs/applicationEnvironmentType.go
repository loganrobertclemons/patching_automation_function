package structs

import log "github.com/sirupsen/logrus"

type ApplicationEnvironmentType struct {
	PatchingProcedure string                              `json:"patching_procedure"`
	Components        map[string]ApplicationComponentType `json:"components"`
}

func (env *ApplicationEnvironmentType) Init() {
	env.Components = make(map[string]ApplicationComponentType)
}

func (env *ApplicationEnvironmentType) Component(name string) *ApplicationComponentType {
	if component, ok := env.Components[name]; ok {
		return &component
	}
	component := new(ApplicationComponentType)
	component.Init()
	env.Components[name] = *component
	return component
}

func (env *ApplicationEnvironmentType) GetComponent(name string) *ApplicationComponentType {
	if component, ok := env.Components[name]; ok {
		return &component
	}
	log.Error("Invalid component: %v", name)
	return nil
}
