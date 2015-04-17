# Go-riak [![GoDoc](https://godoc.org/github.com/3XX0/go-riak?status.png)](http://godoc.org/github.com/3XX0/go-riak)

Go-riak is a Riak client for Go.  
It is implemented as a small abstraction on top of [Riak protocol buffer API](http://docs.basho.com/riak/latest/dev/references/protocol-buffers/).  
Go-riak also provides a driver for [pooly](https://github.com/3XX0/pooly) in order to handle multiple connections to Riak nodes.

## Supported Operations
#### Riak 1.4

Operation | Functions
----------|---------------------------------------------
Server    | ServerInfo, Ping
Key-value | Get, Put, Del
Bucket    | GetBucket, SetBucket, ListBuckets, ListKeys
Query     | Index, MapRed, SearchQuery, GetMany

#### Riak 2.0

Operation | Functions
----------|---------------------------------------------
Server    | Authenticate
Bucket    | ResetBucket, GetBucketType, SetBucketType
Data type | DtFetch, DtUpdate
Yokozuna  | YokozunaIndexGet, YokozunaIndexPut, YokozunaIndexDelete, YokozunaSchemaGet, YokozunaSchemaPut

## Build

```sh
go get github.com/3XX0/go-riak
````

## Example

```go
package main

import "github.com/3XX0/go-riak"
import "github.com/3XX0/pooly"

func main() {
    conf := new(pooly.ServiceConfig)
    conf.Driver = riak.NewDriver()

    s := pooly.NewService("riak", conf)
    defer s.Close()

    s.Add("10.0.0.254:8087")

    c, err := s.GetConn()
    if err != nil {
        panic(err)
    }
    info, err := riak.Client(c).ServerInfo()
    if err != nil {
        panic(err)
    }

    println(info.String())

    if err := c.Release(err, pooly.HostUp); err != nil {
        panic(err)
    }
}
```
