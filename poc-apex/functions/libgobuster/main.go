package main

import (
	"encoding/json"

	"fmt"
	"os"

	"github.com/OJ/gobuster/libgobuster"
	"github.com/apex/go-apex"
	"github.com/hashicorp/go-multierror"
)

type Config struct {
	Url      string `json:"url"`
	Wordlist string `json:"wordlist"`
	Threads  int    `json:"threads"`
	Mode     string `json:"mode"`
}

func ConfigureGobuster(config *Config) (*libgobuster.State, *multierror.Error) {
	var extensions string
	var codes string
	var proxy string

	s := libgobuster.InitState()

	s.Threads = config.Threads
	s.Mode = config.Mode
	s.Wordlist = config.Wordlist
	s.Url = config.Url
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

// Get event JSON as string for debug
func event2str(event json.RawMessage) string {
	eventJsonStr := ""

	eventJson, err := json.Marshal(&event)
	if err != nil {
		eventJsonStr = "Couldn't marshal event JSON"
	}
	eventJsonStr = string(eventJson)

	return eventJsonStr
}

func event2config(event json.RawMessage) (*Config, error) {
	config := Config{
		Url:      "http://devalias.net/",
		Wordlist: "words.txt",
		Threads:  10,
		Mode:     "dir",
	}

	err := json.Unmarshal(event, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func Gobuster(event json.RawMessage) (interface{}, error) {
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

	// Get event JSON as string for debug
	results = append(results, fmt.Sprintf("Event JSON: %s\n", event2str(event)))

	// Parse event JSON as config
	config, err1 := event2config(event)
	if err1 != nil {
		return nil, err1
	}
	results = append(results, fmt.Sprintf("Config: %+v\n", config))

	// Configure libgobuster
	state, err := ConfigureGobuster(config)
	if err.ErrorOrNil() != nil {
		return nil, err
	}

	// State is setup/considered valid. Now we can override things..
	state.Printer = captureDirResultForApex

	// Run libgobuster
	libgobuster.Process(state)

	return results, nil
}

func LocalGobuster() {
	fmt.Println("Note: Running locally")

	var json string
	if len(os.Args) > 2 {
		json = os.Args[2]
	} else {
		json = "{}"
	}

	fmt.Println("JSON:", json)

	results, err := Gobuster([]byte(json))
	if err != nil {
		panic(err)
	}
	fmt.Println("Results:", results)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "local" {
		LocalGobuster()
	} else {
		apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
			return Gobuster(event)
		})
	}
}
