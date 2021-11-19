package runners

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

type runnerGen struct {
	langCommand string
	className   string
	headerCode  string
	footerCode  string
}

// Setups up a runner file for the provided runner.
// Returns the command to run this code, the fileName of the code, and an error
// if the error != nil the other return values will be ""
func (runner *runnerGen) CreateFile(code string, destination string) (string, string, error) {
	outFileName := filepath.Join(destination, runner.className)
	var permissionCode fs.FileMode = 0755
	var codeBuffer bytes.Buffer

	codeBuffer.WriteString(runner.headerCode)
	codeBuffer.WriteString(code)
	codeBuffer.WriteString(runner.footerCode)

	err := ioutil.WriteFile(outFileName, codeBuffer.Bytes(), permissionCode)
	if err != nil {
		return "", "", err
	}
	return runner.langCommand, runner.className, nil
}
