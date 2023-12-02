package main

import (
	"fmt"
	"net"
	"os"
)

const RconPassword = "todefine"
const BufferSize = 8192

func connect(serverIP string, serverPort string) (net.Conn) {
	serverAddress := fmt.Sprintf("%s:%s", serverIP, serverPort)

	print(serverAddress)
	connection, err := net.Dial("udp", serverAddress)

	if err != nil {
		fmt.Printf("Error trying to connect to (%s): %v", serverAddress, err)
	}
	
	return connection
}

func send(connection net.Conn, cmd string) {
	command := fmt.Sprintf("rcon %s %s", RconPassword, cmd)
	commandBytes := []byte(command)
	prefix := []byte{'\xff', '\xff', '\xff' ,'\xff'}
	fullCommandBytes := append(prefix, commandBytes...)

	fmt.Printf("\nSend: %s", fullCommandBytes[4:])

	_, sendErr := connection.Write(fullCommandBytes)
	
	if sendErr != nil {
		fmt.Printf("Error when sending command (%s): %v", command, sendErr) 
	}
}

func read(connection net.Conn) {
	buffer := make([]byte, BufferSize)
    bytesRead, err := connection.Read(buffer)
    if err != nil {
        fmt.Printf("Read err %v\n", err)
        os.Exit(-1)
    }

	if bytesRead >= 4 {
		infos := string(buffer[4:bytesRead])
		fmt.Printf("\nRead (bytesRead: %d): %v\n", bytesRead, infos)
	}
}

func RconCommand(connection net.Conn, command string) {
	send(connection, command)
	read(connection)
}

func closeConnection(connection net.Conn) {
	fmt.Println("Closing connection...")
	err := connection.Close()

	if (err != nil) {
		fmt.Println("Error when closing connection. That's too bad !")
	}
}

func main() {
	connection := connect("localhost", "27960")
	defer closeConnection(connection)

	RconCommand(connection, "bigtext Hello")
	RconCommand(connection, "status")
}