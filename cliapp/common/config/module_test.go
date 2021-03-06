package config

import "testing"

const (
	moduleTestDataEmptyReplaces = `{"srcClone":"srcCloneValue", "srcRev":"srcRevValue", "srcBranch":"srcBranchValue", "srcDir":"srcDirValue", "testClone":"testCloneValue", "testRev":"testRevValue", "testBranch":"testBranchValue", "testDir":"testDirValue"}`
	moduleTestDataSrcCloneFail  = `{"srcClone":["srcCloneValue"]}`
	moduleTestDataSrcRevFail    = `{"srcRev":["srcRevValue"]}`
	moduleTestDataSrcDirFail    = `{"srcDir":["srcDirValue"]}`
	moduleTestDataTestCloneFail = `{"testClone":["testCloneValue"]}`
	moduleTestDataTestRevFail   = `{"testRev":["testRevValue"]}`
	moduleTestDataTestDirFail   = `{"testDir":["testDirValue"]}`
	moduleTestDataReplaces      = `{"replaces":[{"from":"fromvalue","to":"tovalue"}]}`
	moduleTestDataModules       = `[{"srcClone":"srcCloneValue1"},{"srcClone":"srcCloneValue2"}]`
)

func TestModules(t *testing.T) {
	t.Parallel()
	c, err := NewModules([]byte(moduleTestDataModules))
	if err != nil {
		t.Error(err)
		return
	}
	if len(c) != 2 {
		t.Errorf("modules array should contains 2 elements (and it have %d)", len(c))
		return
	}
	if c[0].SourceURL != "srcCloneValue1" {
		t.Errorf("wrong import first module")
		return
	}
	if c[1].SourceURL != "srcCloneValue2" {
		t.Errorf("wrong import second module")
		return
	}
}

func TestModuleSuffixAsString(t *testing.T) {
	t.Parallel()
	c, err := NewModule([]byte(moduleTestDataEmptyReplaces))
	if err != nil {
		t.Error(err)
		return
	}
	if c.SourceURL != "srcCloneValue" {
		t.Errorf("incorrect SourceURL value parsing (expected srcCloneValue and take %s)", c.SourceURL)
		return
	}
	if c.SourceRev != "srcRevValue" {
		t.Errorf("incorrect SourceRev value parsing (expected srcRevValue and take %s)", c.SourceRev)
		return
	}
	if c.SourceBranch != "srcBranchValue" {
		t.Errorf("incorrect SourceBranch value parsing (expected srcBranchValue and take %s)", c.SourceRev)
		return
	}
	if c.SourceDir != "srcDirValue" {
		t.Errorf("incorrect SourceDir value parsing (expected srcDirValue and take %s)", c.SourceDir)
		return
	}
	if c.TestURL != "testCloneValue" {
		t.Errorf("incorrect TestURL value parsing (expected testCloneValue and take %s)", c.TestURL)
		return
	}
	if c.TestRev != "testRevValue" {
		t.Errorf("incorrect TestRev value parsing (expected testRevValue and take %s)", c.TestRev)
		return
	}
	if c.TesteBranch != "testBranchValue" {
		t.Errorf("incorrect TesteBranch value parsing (expected testBranchValue and take %s)", c.TestRev)
		return
	}
	if c.TestDir != "testDirValue" {
		t.Errorf("incorrect TestDir value parsing (expected testDirValue and take %s)", c.TestDir)
		return
	}
}

func TestModuleSrcCloneFail(t *testing.T) {
	t.Parallel()
	_, err := NewModule([]byte(moduleTestDataSrcCloneFail))
	if err.Error() != "expected string and take [\"srcCloneValue\"]" {
		t.Error(err)
		return
	}
}

func TestModuleSrcRevFail(t *testing.T) {
	t.Parallel()
	_, err := NewModule([]byte(moduleTestDataSrcRevFail))
	if err.Error() != "expected string and take [\"srcRevValue\"]" {
		t.Error(err)
		return
	}
}

func TestModuleSrcDirFail(t *testing.T) {
	t.Parallel()
	_, err := NewModule([]byte(moduleTestDataSrcDirFail))
	if err.Error() != "expected string and take [\"srcDirValue\"]" {
		t.Error(err)
		return
	}
}

func TestModuleTestCloneFail(t *testing.T) {
	t.Parallel()
	_, err := NewModule([]byte(moduleTestDataTestCloneFail))
	if err.Error() != "expected string and take [\"testCloneValue\"]" {
		t.Error(err)
		return
	}
}

func TestModuleTestRevFail(t *testing.T) {
	t.Parallel()
	_, err := NewModule([]byte(moduleTestDataTestRevFail))
	if err.Error() != "expected string and take [\"testRevValue\"]" {
		t.Error(err)
		return
	}
}

func TestModuleTestDirFail(t *testing.T) {
	t.Parallel()
	_, err := NewModule([]byte(moduleTestDataTestDirFail))
	if err.Error() != "expected string and take [\"testDirValue\"]" {
		t.Error(err)
		return
	}
}
