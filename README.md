# bitbalance

Shows BTC balance of specified wallet address(es).

## Installation

    go get github.com/theanine/bitbalance

## Running

    go run $GOPATH/src/github.com/theanine/bitbalance/bitbalance.go <addr>...

Or from source:

    go build && ./bitbalance <addr>...

`bitbalance` also accepts a text file as input (with each line as an address):

    go build && ./bitbalance <file>...

## Configuration

The `conf.toml` file allows you to specify a local currency of your choice.
