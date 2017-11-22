# PoC-CLI

PoC for a local CLI using [spf13/cobra](https://github.com/spf13/cobra)

## Setup

* Set up your cobra configuration file: https://github.com/spf13/cobra#configuring-the-cobra-generator

```
go get -u github.com/spf13/cobra/cobra

# Init new project
cobra init gitlab.com/devalias/gopherblazer/poc-cli/blaze

# Add commands
cobra add foo
cobra add bar
```

## Usage

```
cd ./blaze/
dep ensure
go run main.go
```

# Future Improvements

* Design a way to allow sub-commands (eg `blaze nmap fastLocal`)
* Design a way to allow 'templated' commands, so I can specify which keys get replaced in a programs arguments, etc
* Add different styles of 'command runner' (eg. helper for docker containers, standard binary, OpenFaaS, lambda, etc)
* Remove the need to specify the full path to the executable
* Make the 'dynamicCommands' use the main configuration?
* Load remote configuration and generate dynamic commands from that
