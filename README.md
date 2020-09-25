## Imports in Golang

Imports are done relatively to the `GOPATH`:

```golang
import (
  "go-imports/level1"                 
  Cool "go-imports/level1/level2"
  . "go-imports/level1/level2/level3"
)
```

Imports can be done with the dot or can have an alias.

Dot imports do not expose the private variables and functions from the imported packages.

### Deploy

```shell script
git clone https://github.com/peterdee/go-imports
cd ./go-imports
```

### Launch

```shell script
go run *.go
```
