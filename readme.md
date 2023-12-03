# Quake3 Rcon Go

This project is an implementation of quake3 rcon using go.

## Main usage

```golang

import (
	quake3_rcon "github.com/AntoineMeresse/quake3-rcon-go"
)

func main() {
    // ~~~~~~~~~~~~ Setup ~~~~~~~~~~~~
    rcon := quake3_rcon.Rcon{ServerIp: "localhost", ServerPort: 27960, Password: "toreplace", Connection: nil}

    rcon.Connect()
    defer rcon.CloseConnection()
    
    // ~~~~~~~~~~~~ Commands ~~~~~~~~~~~~
    // Example of command which doesn't require to handle response:
	res := rcon.RconCommand("bigtext Hello")

	// Example of command which might require to handle response:
	res = rcon.RconCommand("sv_fps")
	
    // In 2 steps:
	responseType, datas := quake3_rcon.SplitReadInfos(res)
	fmt.Printf("[Response] Type: %s, datas: %v", responseType, datas) 
    // [Response] Type: print, datas: ["sv_fps" is:"125^7" default:"20^7"]
	
    // In 1 step: (mainly for debugging purpose) Shorter command with some nice printing. Does not return anything.
	quake3_rcon.PrintSplitReadInfos(res)
    //     ==================================== Print Read Infos ====================================
    // Type: print
    // Lines: 1, datas : ["sv_fps" is:"125^7" default:"20^7"]
    //    |---->  1) "sv_fps" is:"125^7" default:"20^7"	
}
```