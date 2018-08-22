// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/design
// --out=$(GOPATH)/src/github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// MemosHref returns the resource href.
func MemosHref(id interface{}) string {
	paramid := strings.TrimLeftFunc(fmt.Sprintf("%v", id), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/memos/%v", paramid)
}
