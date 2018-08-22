// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application User Types
//
// Command:
// $ goagen
// --design=github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api/design
// --out=$(GOPATH)/src/github.com/akm/gae_go-datastore-goa-goon-viron-react-redux-example/api
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
)

// memoPayload user type.
type memoPayload struct {
	// Content of memo
	Content *string `form:"content,omitempty" json:"content,omitempty" yaml:"content,omitempty" xml:"content,omitempty"`
	// Auther name
	CreatedBy *string `form:"created_by,omitempty" json:"created_by,omitempty" yaml:"created_by,omitempty" xml:"created_by,omitempty"`
	// Shared to public
	Shared *bool `form:"shared,omitempty" json:"shared,omitempty" yaml:"shared,omitempty" xml:"shared,omitempty"`
}

// Validate validates the memoPayload type instance.
func (ut *memoPayload) Validate() (err error) {
	if ut.Content == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "content"))
	}
	return
}

// Publicize creates MemoPayload from memoPayload
func (ut *memoPayload) Publicize() *MemoPayload {
	var pub MemoPayload
	if ut.Content != nil {
		pub.Content = *ut.Content
	}
	if ut.CreatedBy != nil {
		pub.CreatedBy = ut.CreatedBy
	}
	if ut.Shared != nil {
		pub.Shared = ut.Shared
	}
	return &pub
}

// MemoPayload user type.
type MemoPayload struct {
	// Content of memo
	Content string `form:"content" json:"content" yaml:"content" xml:"content"`
	// Auther name
	CreatedBy *string `form:"created_by,omitempty" json:"created_by,omitempty" yaml:"created_by,omitempty" xml:"created_by,omitempty"`
	// Shared to public
	Shared *bool `form:"shared,omitempty" json:"shared,omitempty" yaml:"shared,omitempty" xml:"shared,omitempty"`
}

// Validate validates the MemoPayload type instance.
func (ut *MemoPayload) Validate() (err error) {
	if ut.Content == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "content"))
	}
	return
}
