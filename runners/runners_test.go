package runners_test

import (
	"testing"

	"github.com/hcam93/exe/runners"
)

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

// More tests to be added here.....
