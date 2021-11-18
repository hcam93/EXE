package runners

import (
	"bytes"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

//pythonRunner creates the runner file that will be used to run
//the executable.
//It takes in the code that you want to run, and the destination where you
//want to put the runner file and returns the command to run the file and
//the name of the file or any errors if it encounters them during the creation
//process.
func (runner *languageTemplate) CreateFile(code string, destination string) (string, string, error) {
	outFileName := filepath.Join(destination, runner.className)
	var permissionCode fs.FileMode = 0755
	var codeFormatter bytes.Buffer
	runner.writeStringCode(&codeFormatter, "header")
	codeFormatter.WriteString(code)
	runner.writeStringCode(&codeFormatter, "footer")

	err := ioutil.WriteFile(outFileName, codeFormatter.Bytes(), permissionCode)
	if err != nil {
		return "", "", err
	}
	return runner.langCommand, runner.className, nil
}

func (runner *languageTemplate) writeStringCode(codeFormatter *bytes.Buffer, stringLocation string){
	if(stringLocation=="header"){
		for i := range runner.headerCode {	codeFormatter.WriteString(runner.headerCode[i])
		}
	}else if(stringLocation=="footer"){
		for i := range runner.headerCode {	codeFormatter.WriteString(runner.footerCode[i])
		}
	}
	
}
