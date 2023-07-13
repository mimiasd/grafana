package registry

import (
	"testing"

	"github.com/grafana/dskit/services"
	"github.com/stretchr/testify/require"

	"github.com/grafana/grafana/pkg/api"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/modules"
	"github.com/grafana/grafana/pkg/server/backgroundsvcs"
)

func TestProvideRegistry(t *testing.T) {
	var registeredInvisibleModules []string
	var registeredModules []string

	moduleManager := &modules.MockModuleManager{
		RegisterModuleFunc: func(name string, initFn func() (services.Service, error)) {
			registeredModules = append(registeredModules, name)
		},
		RegisterInvisibleModuleFunc: func(name string, initFn func() (services.Service, error)) {
			registeredInvisibleModules = append(registeredInvisibleModules, name)
		},
	}

	svcRegistry := backgroundsvcs.NewBackgroundServiceRegistry()
	svcRunner := backgroundsvcs.ProvideBackgroundServiceRunner(svcRegistry)
	// the bare minimum apiServer for this test
	apiServer := &api.HTTPServer{}
	apiServer.NamedService = services.NewBasicService(nil, nil, nil).WithName(modules.HTTPServer)

	r := ProvideRegistry(moduleManager, svcRunner, apiServer)
	require.NotNil(t, r)
	require.Equal(t, []string{modules.BackgroundServices, modules.HTTPServer}, registeredInvisibleModules)
	require.Equal(t, []string{modules.All}, registeredModules)
}

func TestNewRegistry(t *testing.T) {
	var registeredInvisibleModules []string
	var registeredModules []string

	moduleManager := &modules.MockModuleManager{
		RegisterModuleFunc: func(name string, initFn func() (services.Service, error)) {
			registeredModules = append(registeredModules, name)
		},
		RegisterInvisibleModuleFunc: func(name string, initFn func() (services.Service, error)) {
			registeredInvisibleModules = append(registeredInvisibleModules, name)
		},
	}

	mockSvcName := "test-registry"
	mockSvc := modules.NewMockNamedService(mockSvcName)

	r := newRegistry(log.New("modules.registry"), moduleManager, mockSvc)
	require.NotNil(t, r)
	require.Equal(t, []string{mockSvcName}, registeredInvisibleModules)
	require.Equal(t, []string{modules.All}, registeredModules)
}
