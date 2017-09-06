// Copyright © 2017 Glenn 'devalias' Grant <glenn@devalias.net>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdViper *viper.Viper
var commands []*cobra.Command

type ConfigCommands struct {
	Commands []ConfigCommand
}

type ConfigCommand struct {
	Use   string
	Short string
	Long  string
	Run   string
	Exec  *ConfigExec
}

type ConfigExec struct {
	Command string
	Args    []string
}

func init() {
	cmdViper = viper.New()
	cmdViper.SetConfigName("commands")
	cmdViper.AddConfigPath(".")
	cmdViper.SetConfigFile(cfgFile)

	// Find and read the config file
	err := cmdViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Unmarshal the commands
	var cfg ConfigCommands
	err = cmdViper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct, %v", err))
	}

	// fmt.Printf("%+v\n", cfg)

	// Create dynamic commands
	for _, cfgCmd := range cfg.Commands {
		commands = append(commands, newCommand(cfgCmd))
	}

	RootCmd.AddCommand(commands...)
}

func newCommand(cfgCmd ConfigCommand) *cobra.Command {
	// Create a new command from the config values
	return &cobra.Command{
		Use:   cfgCmd.Use,
		Short: cfgCmd.Short,
		Long:  cfgCmd.Long,
		Run: func(cmd *cobra.Command, args []string) {
			if cfgCmd.Exec != nil {
				execProgram(*cfgCmd.Exec)
			} else {
				// Print some placeholder stuff
				fmt.Println("FOO:" + cfgCmd.Run)
			}
		},
	}
}

func execProgram(cfgExec ConfigExec) {
	// Execute the command
	argv0 := []string{cfgExec.Command}
	argv := append(argv0, cfgExec.Args...)

	err := syscall.Exec(cfgExec.Command, argv, os.Environ())
	if err != nil {
		panic(err)
	}
}
