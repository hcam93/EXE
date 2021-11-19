package runners

type Runner interface {
	CreateFile(string, string) (string, string, error)
}

var supportedLanguages = map[string]Runner{
	"python": &runnerGen{
		langCommand: "python3",
		className:   "PythonRunner.py",
		headerCode:  "numpy.py \n",
		footerCode:  "",
	},
	"java": &runnerGen{
		langCommand: "java",
		className:   "javaRunner.java",
		headerCode:  "import java.util.*;\nimport java.lang.*;\npublic class JavaRunner{\n",
		footerCode:  "\n}",
	},
}
