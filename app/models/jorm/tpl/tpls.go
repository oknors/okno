package tpl

import (
	"fmt"
	"io/ioutil"
	"text/template"

	"git.parallelcoin.io/dev/utl"
)

// TemplateHandler reads in templates
func TemplateHandler() *template.Template {
	var str string
	fls := utl.GPFiles("./tpl/fls")
	for _, fl := range fls {
		ff := utl.GPFiles("./tpl/fls/" + fl)
		for _, f := range ff {
			v, err := ioutil.ReadFile("./tpl/fls/" + fl + "/" + f + "." + fl)
			if err != nil {
				fmt.Print(err)
			}
			s := `{{define "` + f + "_" + fl + `"}}` + string(v) + `{{end}}`
			str = str + s
		}
	}
	// fmt.Println("sssssssss", str)
	t := template.Must(template.New("").Parse(str))
	return t
}
