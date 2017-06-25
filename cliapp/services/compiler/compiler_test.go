package compiler

/*
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
	if err = mapp.RootFilespace().WriteFile(modulesDefPath, []byte(testModulesDefJSON), 0766); err != nil {
		t.Error(err)
		return
	}
	if err = RegisterDependencies(mapp.DependencyProvider()); err != nil {
		t.Error(err)
		return
	}
	// test
	var deps struct {
		Modules services.Modules `dependency:"ModulesService"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if err = deps.Modules.Init(); err != nil {
		t.Error(err)
		return
	}
	var modulesConfig []*config.Module
	if modulesConfig, err = deps.Modules.ModulesConfig(); err != nil {
		t.Error(err)
		return
	}
	if len(modulesConfig) != 1 {
		t.Errorf("expected one module and take %d", len(modulesConfig))
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
		Modules services.Modules `dependency:"ModulesService"`
	}
	if err = mapp.DependencyProvider().InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if err = deps.Modules.Init(); err != nil {
		t.Error(err)
		return
	}
	var modulesConfig []*config.Module
	if modulesConfig, err = deps.Modules.ModulesConfig(); err != nil {
		t.Error(err)
		return
	}
	if len(modulesConfig) != 0 {
		t.Errorf("expected no modules and take %d", len(modulesConfig))
		return
	}
}
*/
