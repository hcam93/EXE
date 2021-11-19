package runners_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hcam93/exe/runners"
)

var (
	testTmpDir string
)

func TestMain(m *testing.M) {
	err := getTmpDir(m)
	if err != nil {
		fmt.Fprint(os.Stderr, "Runners tests could not create temp dir for test files")
		os.Exit(1)
	}
	code := m.Run()
	err = os.Remove(testTmpDir)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to remove temp directory after tests for runners package.\n")
		fmt.Fprint(os.Stderr, "Did you properly cleanup your files in your tests?\n")
		// Don't fial here, just warn the user.
	}
	os.Exit(code)
}

func TestCreateRunnerPython(t *testing.T) {
	t.Parallel()
	_, err := runners.CreateRunner("python")
	if err != nil {
		t.Errorf("CreateRunner returned a non-nil error with a valid input.\n")
	}
}

func TestCreateRunnerJava(t *testing.T) {
	t.Parallel()
	_, err := runners.CreateRunner("java")
	if err != nil {
		t.Errorf("CreateRunner returned a non-nil error with a valid input.\n")
	}
}

func TestCreateFilePython(t *testing.T) {
	// Anything that writes to the FS should not be parrall just so as not to conflic
	// with any other potential tests past/present/future.
	run, _ := runners.CreateRunner("python")
	command, classname, err := run.CreateFile("print(\"Hello\")", testTmpDir)
	if err != nil {
		t.Errorf("CreateFile failed with valid input.")
	}
	if command != "python3" {
		t.Errorf("returned command to run ths code was incorrect.")
	}
	if classname != "PythonRunner.py" {
		t.Error("Returned file name was incorrect.")
	}
	absFilePath := filepath.Join(testTmpDir, classname)
	dat, err := os.ReadFile(absFilePath)
	if err != nil {
		t.Error("The runner python file was not created.")
	}
	if strings.Compare(string(dat), "import numpy \nprint(\"Hello\")") != 0 {
		t.Error("The created python file was not formed correctly.")
	}
	os.Remove(absFilePath)
}

func getTmpDir(m *testing.M) error {
	tmpDir := os.TempDir()
	if tmpDir == "" {
		return errors.New("Failed to get a temporary directory.")
	}
	dir, err := ioutil.TempDir(tmpDir, "runnersTests")
	if err != nil {
		return errors.New("Failed to create temporary directory for tests")
	}
	testTmpDir = dir
	return nil
}
