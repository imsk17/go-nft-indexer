# go-nft-indexer

This application listens for `Transfer` events in any **evm** compatible blockchain to index NFTs.

### Setup

How to run this thing

* Make a new .env file and copy the contents from .env.example to this file
* Fill the variables
* Then run, from the project root, `go run main.go`

### Goals

What I would want to achieve from this thing -

* [ ] Persistence in Mongo/SQL
* [ ] Caching
* [ ] Support for Events Other Than Transfers