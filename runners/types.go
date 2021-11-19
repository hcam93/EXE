package runners

type Runner interface {
	CreateFile(string, string) (string, string, error)
}

var supportedLanguages = map[string]Runner{
	"python": &languageTemplate{
		langCommand: "python3",
		className:   "PythonRunner.py",
		headerCode:  []string{"numpy.py \n"},
		footerCode:  []string{},
	},
	"java": &languageTemplate{
		langCommand: "java",
		className:   "javaRunner.java",
		headerCode:  []string{"import java.util.*;\n", "import java.lang.*;\n", "public class JavaRunner{\n"},
		footerCode:  []string{"\n}"},
	},
}
