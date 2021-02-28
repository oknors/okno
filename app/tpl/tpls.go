package tpl

import (
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/oknors/okno/pkg/utl"
)

// TemplateHandler reads in templates
func TemplateHandler(path string) *template.Template {
	var str string
	fls := utl.GPFiles(path + "/tpl/")
	for _, fl := range fls {
		ff := utl.GPFiles(path + "/tpl/" + fl)
		for _, f := range ff {
			v, err := ioutil.ReadFile(path + "/tpl/" + fl + "/" + f + "." + fl)
			if err != nil {
				fmt.Print(err)
			}
			s := `{{define "` + f + "_" + fl + `"}}` + string(v) + `{{end}}`
			str = str + s
		}
	}
	// fmt.Println("sssssssss", str)
	temp, err := template.New("").Parse(str)
	utl.ErrorLog(err)

	t := template.Must(temp, nil)
	return t
}
