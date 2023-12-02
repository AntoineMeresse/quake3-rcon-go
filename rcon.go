package rcon

import (
	"fmt"
	"net"
	"os"
)

const BufferSize = 8192

type Rcon struct {
	ServerIp string
	ServerPort int
	Password string
	Connection net.Conn
}

func (rcon *Rcon) Connect() {
	serverAddress := fmt.Sprintf("%s:%d", rcon.ServerIp, rcon.ServerPort)

	fmt.Println(serverAddress)
	conn, err := net.Dial("udp", serverAddress)

	if err != nil {
		fmt.Printf("Error trying to connect to (%s): %v", serverAddress, err)
		os.Exit(-1)
	}
	
	rcon.Connection = conn
}

func (rcon Rcon) Send(cmd string) {
	command := fmt.Sprintf("rcon %s %s", rcon.Password, cmd)
	commandBytes := []byte(command)
	prefix := []byte{'\xff', '\xff', '\xff' ,'\xff'}
	fullCommandBytes := append(prefix, commandBytes...)

	fmt.Printf("\nSend: %s", fullCommandBytes[4:])

	_, sendErr := rcon.Connection.Write(fullCommandBytes)
	
	if sendErr != nil {
		fmt.Printf("Error when sending command (%s): %v", command, sendErr) 
	}
}

func (rcon Rcon) Read() {
	buffer := make([]byte, BufferSize)
    bytesRead, err := rcon.Connection.Read(buffer)
    if err != nil {
        fmt.Printf("Read err %v\n", err)
        os.Exit(-1)
    }

	if bytesRead >= 4 {
		infos := string(buffer[4:bytesRead])
		fmt.Printf("\nRead (bytesRead: %d): %v\n", bytesRead, infos)
	}
}

func (rcon Rcon) RconCommand(command string) {
	if rcon.Connection != nil {
		rcon.Send(command)
		rcon.Read()
	}
}

func (rcon *Rcon) CloseConnection() {
	fmt.Println("Closing connection...")
	err := rcon.Connection.Close()

	if (err != nil) {
		fmt.Println("Error when closing connection. That's too bad !")
	} else {
		rcon.Connection = nil
	}
}

// Usage: 
// func main() {
// 	rcon := rcon.Rcon{ServerIp: "localhost", ServerPort: 27960, Password: "toreplace", Connection: nil}

// 	rcon.Connect()
// 	defer rcon.CloseConnection()

// 	rcon.RconCommand("bigtext Hello")
// 	rcon.RconCommand("status")
// }