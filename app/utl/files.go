package utl

import (
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// GPFiles reads in the list of template files
func GPFiles(p string) (s []string) {
	cm, err := ioutil.ReadDir(p)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range cm {
		cn := c.Name()
		n := strings.ToLower(strings.TrimSuffix(cn, path.Ext(cn)))
		s = append(s, n)
	}
	return s
}
