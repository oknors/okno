package utl

import (
	"strconv"
)

func InsertFloat(i interface{}) (f float64) {
	switch t := i.(type) {
	case string:
		if s, err := strconv.ParseFloat(t, 64); err == nil {
			f = s
		}
	case float64:
		f = t
	}
	return
}


func InsertString(s string, i interface{}) string {
	if s == "" {
		if i != nil {
			s = i.(string)
		}
	}
	return s
}

func InsertStringInSlice(s []string, i interface{}) []string {
	if i != nil {
		iSingle := i.(string)
		if len(s) > 0 {
			for _, sSingle := range s {
				if iSingle != "" {
					if iSingle != sSingle {
						s = append(s, iSingle)
					}
				}
			}
		} else {
			s = append(s, iSingle)
		}
	}
	return s
}

func InsertStringSlice(s []string, i interface{}) []string {
	if i != nil {
		switch v := i.(type) {
		case []interface{}:
			for _, iSingle := range v {
				s = InsertStringInSlice(s, iSingle)
			}
		case interface{}:
			s = InsertStringInSlice(s, v)
		}
	}
	return s
}
