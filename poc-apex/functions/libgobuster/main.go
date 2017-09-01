package main

import (
	"encoding/json"

	// "fmt"
	// "os"

	"github.com/OJ/gobuster/libgobuster"
	"github.com/apex/go-apex"
	"github.com/hashicorp/go-multierror"
)

func ConfigureGobuster(url string) (*libgobuster.State, *multierror.Error) {
	var extensions string
	var codes string
	var proxy string

	s := libgobuster.InitState()

	s.Threads = 10
	s.Mode = "dir"
	s.Wordlist = "words.txt"
	s.Url = url

	extensions = ""
	codes = "200,204,301,302,307"
	proxy = ""

	if err := libgobuster.ValidateState(&s, extensions, codes, proxy); err.ErrorOrNil() != nil {
		return nil, err
	} else {
		return &s, nil
	}
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var out string

		state, err := ConfigureGobuster("http://devalias.net/")
		if err.ErrorOrNil() != nil {
			return nil, err
		}

		libgobuster.Process(state)

		return out, nil
	})
}
