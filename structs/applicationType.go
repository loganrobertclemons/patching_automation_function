package structs

import log "github.com/sirupsen/logrus"

type ApplicationType struct {
	PatchingProcedure string                                `json:"patching_procedure"`
	Environments      map[string]ApplicationEnvironmentType `json:"environments"`
}

func (app *ApplicationType) Init() {
	app.Environments = make(map[string]ApplicationEnvironmentType)
}

func (app *ApplicationType) Environment(name string) *ApplicationEnvironmentType {
	if env, ok := app.Environments[name]; ok {
		return &env
	}
	env := new(ApplicationEnvironmentType)
	env.Init()
	app.Environments[name] = *env
	return env
}

func (app *ApplicationType) GetEnvironment(name string) *ApplicationEnvironmentType {
	if env, ok := app.Environments[name]; ok {
		return &env
	}
	log.Error("Invalid environment: %v", name)
	return nil
}

func (app *ApplicationType) GetEnvironmentList() []string {
	keys := make([]string, 0, len(app.Environments))
	for k := range app.Environments {
		keys = append(keys, k)
	}
	return keys
}
