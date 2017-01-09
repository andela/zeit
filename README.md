# Zeit

## Overview
Zeit enables users log billable hours via the CLI

## Installation
Ensure `$GOPATH` and `$GOROOT` is already configured before running the command below. See [guide](https://golang.org/doc/install) for more details
```sh
    go get github.com/andela/zeit
```

## Usage

### Getting Help
At any point, just use the `--help` flag. E.g:
```
zeit --help 
zeit history --help
```

### Available Commands
```
zeit start - Start logging time
zeit stop - Stop current timer
zeit login - Login with google account
zeit history - View all previously logged entries
zeit entries - View entries matching a specified range
zeit log - Log billable hours for a specified day
```
