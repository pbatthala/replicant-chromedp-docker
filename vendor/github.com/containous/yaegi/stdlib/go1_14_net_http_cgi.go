// Code generated by 'github.com/containous/yaegi/extract net/http/cgi'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"net/http/cgi"
	"reflect"
)

func init() {
	Symbols["net/http/cgi"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Request":        reflect.ValueOf(cgi.Request),
		"RequestFromMap": reflect.ValueOf(cgi.RequestFromMap),
		"Serve":          reflect.ValueOf(cgi.Serve),

		// type definitions
		"Handler": reflect.ValueOf((*cgi.Handler)(nil)),
	}
}
