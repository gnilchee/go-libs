## TinyDS

### A Tiny K/V Datastore wrapper for the V4 [BadgerDB](https://dgraph.io/docs/badger/) database.

##### This wrapper exposes and simplifies common basic actions.

#### Installing
Install using Go 1.21.

```
$ go get github.com/gnilchee/go-libs/tinyds
```

#### API
- `Open` opens the BadgerDB database
    - takes `path` as an argument
    - returns TinyDS object
    - will exit if issue with opening or creating database
- `Set` sets a key/value pair
    - takes `key` and `value` as arguments
    - returns err (if any)
> `key` and `value` are of type string
- `SetwithTTL` sets a key/value pair with TTL (secs)
    - takes `key`, `value`, `ttl` (in secs) as arguments
    - returns err (if any)
- `Get` returns value of a key
    - takes `key` as an argument
    - returns string and error (if any)
- `Delete` deltes a key and its value
    - takes `key` as an argument
    - returns error (if any)
- `Close` closes the BadgerDB database
    - takes no arguments
    - returns an error (if any)

#### Usage
Creating or opening a k/v store
```
opts := tinyds.DefaultOptions("data")
// if you want to disable BadgerDB startup/shutdown logging
opts.Logger = nil

kv := tinyds.Open(opts)
defer kv.Close()
```
Basic set, set with TTL, get, and delete
```
kv.Set("key", "value")
kv.SetwithTTL("key", "value", 3600) // 60 min TTL
kv.Get("key")
kv.Delete("key")
```