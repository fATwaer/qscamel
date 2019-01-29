package checker

import (
	"testing"

	"github.com/yunify/qscamel/tests/executer"
	"github.com/yunify/qscamel/tests/generater"
)

func TestDefaultRunCopy(t *testing.T) {

	fileMap, err := generater.CreateTestDefaultFile(
		"copy", "fs", "fs", nil, nil, true)
	defer generater.CleanTestTempFile(fileMap)
	err = generater.CreateLocalSrcTestRandDirFile(
		fileMap, 4, 1, generater.MB*2, 2)
	err = generater.CreateLocalDstDir(fileMap)
	if err != nil {
		t.Fatal(err)
	}

	if err = executer.Execute(fileMap, "run"); err != nil {
		t.Fatal(err)
	}

	// check running ouput
	if err := executer.CheckOutput(fileMap,
		"Start copying single object [A-Z0-9]*/TESTFILE\\d+.camel", 8, true); err != nil {
		t.Fatal(err)
	}
	if err := executer.CheckOutput(fileMap,
		"Task [a-z0-9]* has been finished", 1, true); err != nil {
		t.Fatal(err)
	}

}

func TestDefaultDelete(t *testing.T) {
	fileMap, err := generater.CreateTestDefaultFile(
		"copy", "fs", "fs", nil, nil, true)
	defer generater.CleanTestTempFile(fileMap)
	err = generater.CreateLocalSrcTestRandDirFile(
		fileMap, 4, 1, generater.MB*4, 2)
	err = generater.CreateLocalDstDir(fileMap)
	if err != nil {
		t.Fatal(err)
	}

	// run command
	if err = executer.Execute(fileMap, "run"); err != nil {
		t.Fatal(err)
	}
	if err := executer.CheckOutput(fileMap,
		"Task [a-z0-9]* has been finished", 1, true); err != nil {
		t.Fatal(err)
	}
	// delete command
	(*fileMap)["delname"] = (*fileMap)["name"]
	if err := executer.Execute(fileMap, "delete"); err != nil {
		t.Fatal(err)
	}
	// check delete output
	if err := executer.CheckOutput(fileMap,
		"Task [a-z0-9]* has been deleted", 1, true); err != nil {
		t.Fatal(err)
	}
}

func TestDefalutStatus(t *testing.T) {
	// env set
	fileMap, err := generater.CreateTestDefaultFile(
		"copy", "fs", "fs", nil, nil, true)
	defer generater.CleanTestTempFile(fileMap)
	err = generater.CreateLocalSrcTestRandDirFile(
		fileMap, 4, 1, generater.MB*4, 2)
	err = generater.CreateLocalDstDir(fileMap)
	if err != nil {
		t.Fatal(err)
	}
	// run command
	if err = executer.Execute(fileMap, "run"); err != nil {
		t.Fatal(err)
	}
	// status command
	if err = executer.Execute(fileMap, "status"); err != nil {
		t.Fatal(err)
	}
	// check status output
	if err := executer.CheckOutput(fileMap,
		"Show status started", 1, true); err != nil {
		t.Fatal(err)
	}
	if err := executer.CheckOutput(fileMap,
		"There are 1 tasks totally", 1, true); err != nil {
		t.Fatal(err)
	}

}

func TestDefaultClean(t *testing.T) {
	// env set
	fileMap, err := generater.CreateTestDefaultFile(
		"copy", "fs", "fs", nil, nil, true)
	defer generater.CleanTestTempFile(fileMap)
	err = generater.CreateLocalSrcTestRandDirFile(
		fileMap, 4, 1, generater.MB*4, 2)
	err = generater.CreateLocalDstDir(fileMap)
	if err != nil {
		t.Fatal(err)
	}
	// run command
	if err = executer.Execute(fileMap, "run"); err != nil {
		t.Fatal(err)
	}
	// clean command
	if err = executer.Execute(fileMap, "clean"); err != nil {
		t.Fatal(err)
	}
	// check clean output
	if err := executer.CheckOutput(fileMap,
		"Clean started", 1, true); err != nil {
		t.Fatal(err)
	}
	if err := executer.CheckOutput(fileMap,
		"Task [a-z0-9]* has been cleaned", 1, true); err != nil {
		t.Fatal(err)
	}
}
