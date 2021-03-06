package modules

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goatcms/goatcli/cliapp/common/config"
	"github.com/goatcms/goatcli/cliapp/services"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/mockupapp"
)

const (
	testModulesDefJSON = `[{"srcClone":"srcCloneValue", "srcRev":"srcRevValue", "srcDir":"srcDirValue", "testClone":"testCloneValue", "testRev":"testRevValue", "testDir":"testDirValue"}]`
)

func TestModulesFromFile(t *testing.T) {
	var err error
	t.Parallel()
	// prepare mockup application & data
	output := new(bytes.Buffer)
	mapp, err := mockupapp.NewApp(mockupapp.MockupOptions{
		Input:  gio.NewInput(strings.NewReader("my_insert_value\n")),
		Output: gio.NewOutput(output),
	})
	if err != nil {
		t.Error(err)
		return
	}
	if err = mapp.RootFilespace().WriteFile(ModulesDefPath, []byte(testModulesDefJSON), 0766); err != nil {
		t.Error(err)
		return
	}
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		Modules services.ModulesService `dependency:"ModulesService"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	var modules []*config.Module
	if modules, err = deps.Modules.ReadDefFromFS(mapp.RootFilespace()); err != nil {
		t.Error(err)
		return
	}
	if len(modules) != 1 {
		t.Errorf("expected one module and take %d", len(modules))
		return
	}
}

func TestModulesDefaultEmpty(t *testing.T) {
	var err error
	t.Parallel()
	// prepare mockup application & data
	output := new(bytes.Buffer)
	mapp, err := mockupapp.NewApp(mockupapp.MockupOptions{
		Input:  gio.NewInput(strings.NewReader("")),
		Output: gio.NewOutput(output),
	})
	if err != nil {
		t.Error(err)
		return
	}
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		Modules services.ModulesService `dependency:"ModulesService"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	var modules []*config.Module
	if modules, err = deps.Modules.ReadDefFromFS(mapp.RootFilespace()); err != nil {
		t.Error(err)
		return
	}
	if len(modules) != 0 {
		t.Errorf("expected no modules and take %d", len(modules))
		return
	}
}
