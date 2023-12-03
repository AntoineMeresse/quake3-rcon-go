# Quake3 Rcon Go

This project is an implementation of quake3 rcon using golang.

Developped mainly for Urban Terror by Antoine Méresse (IG: Flirow)

## Usage

## Download

```shell
go install github.com/AntoineMeresse/quake3-rcon-go@latest
```

### Main Usage

```golang

import (
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

func main() {

    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ~~~~~~~~~~~~~ Setup ~~~~~~~~~~~~~~
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    rcon := quake3_rcon.Rcon{ServerIp: "localhost", ServerPort: 27960, Password: "toreplace", Connection: nil}

    rcon.Connect()
    defer rcon.CloseConnection()
    
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    // ~~~~~~~~~~~~ Commands ~~~~~~~~~~~~
    // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    
    // Example of command which doesn't require to handle response:
	res := rcon.RconCommand("bigtext Hello")

	// Example of command which might require to handle response:
	res = rcon.RconCommand("sv_fps")
	responseType, datas := quake3_rcon.SplitReadInfos(res)
}
```

### Other useful code sample

If you need to do some processing base on response datas:

```golang
responseType, datas := quake3_rcon.SplitReadInfos(res)
fmt.Printf("[Response] Type: %s, datas: %v", responseType, datas) 

// Will produce:

// [Response] Type: print, datas: ["sv_fps" is:"125^7" default:"20^7"]
```

[Debugging] Shorter command with some nice printing. /!\ Does not return anything /!\

```golang
quake3_rcon.PrintSplitReadInfos(res)
    
// Will produce:

// ~~~~~~~~~~ Print Read Infos ~~~~~~~~~~
// Type: print
// Line: 1
//    |---->  1) "sv_fps" is:"125^7" default:"20^7"
```

## Licence

This project is released under MIT Licence.

## Contribute

If you find any bugs or want to improve this software :
- Feel free to contact me on discord: fliro 
- Open & submit a PR