package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"os"
)
var INC string = "inc"
var VAR string = "var"
var VAR_MAP map[string]string

func setVarMap(maps map[string ]string) {
	VAR_MAP = maps
}
func main()  {

	vars := map[string]string {
		"test":"Hello World",
	}
	setVarMap(vars)
	fmt.Println(evaluate("var_test"))
}
func evaluate(str string) string {

	codes :=  strings.Split(str,"_")
	var func_ string = codes[0]
	var code_ string = codes[1]
	if func_ == INC {
		b, err := ioutil.ReadFile (code_+".gojub") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		content := string(b)
		return content
	}
	if func_ == VAR {
		if VAR_MAP == nil {
			println("VAR_MAP is nil")
			os.Exit(0)
		}
		return VAR_MAP[code_]
	}

	return ""
}
