package tpl

import (
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/oknors/okno/app/utl"
)

// TemplateHandler reads in templates
func TemplateHandler() *template.Template {
	var str string
	fls := utl.GPFiles("./app/models/jorm/tpl/fls")
	for _, fl := range fls {
		ff := utl.GPFiles("./app/models/jorm/tpl/fls/" + fl)
		for _, f := range ff {
			v, err := ioutil.ReadFile("./app/models/jorm/tpl/fls/" + fl + "/" + f + "." + fl)
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
