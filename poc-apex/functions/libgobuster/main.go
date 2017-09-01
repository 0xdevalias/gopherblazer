package main

import (
	"encoding/json"

	"fmt"
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
	s.Quiet = true

	extensions = ""
	codes = "200,204,301,302,307"
	proxy = ""

	if err := libgobuster.ValidateState(&s, extensions, codes, proxy); err.ErrorOrNil() != nil {
		return nil, err
	} else {
		return &s, nil
	}
}

// type ResultStruct struct {
// 	url    string
// 	status int
// }

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var results []string
		// var results []ResultStruct

		// Defining this here so we can close around the 'results' slice, inside the handler function
		var captureDirResultForApex = func(s *libgobuster.State, r *libgobuster.Result) {
			output := ""

			if s.StatusCodes.Contains(r.Status) || s.Verbose {
				if s.Expanded {
					output += s.Url
				} else {
					output += "/"
				}
				output += r.Entity

				if !s.NoStatus {
					output += fmt.Sprintf(" (Status: %d)", r.Status)
				}

				if r.Size != nil {
					output += fmt.Sprintf(" [Size: %d]", *r.Size)
				}
				output += "\n"

				// r := ResultStruct{
				// 	url:    r.Entity,
				// 	status: r.Status,
				// }

				// fmt.Fprintln(os.Stderr, output)
				results = append(results, output)
				// results = append(results, r)
			}
		}

		// Configure libgobuster
		state, err := ConfigureGobuster("http://devalias.net/")
		if err.ErrorOrNil() != nil {
			return nil, err
		}

		// State is setup/considered valid. Now we can override things..
		state.Printer = captureDirResultForApex

		// Run libgobuster
		libgobuster.Process(state)

		return results, nil
	})
}
