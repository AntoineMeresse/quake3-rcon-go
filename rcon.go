package quake3_rcon

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

const BufferSize = 8192
var PacketPrefix = []byte{'\xff', '\xff', '\xff' ,'\xff'}

type Rcon struct {
	ServerIp string
	ServerPort string
	Password string
	Connection net.Conn
}

func (rcon *Rcon) Connect() {
	serverAddress := fmt.Sprintf("%s:%s", rcon.ServerIp, rcon.ServerPort)
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
	
	fullCommandBytes := append(PacketPrefix, commandBytes...)
	_, sendErr := rcon.Connection.Write(fullCommandBytes)
	
	if sendErr != nil {
		fmt.Printf("Error while sending command (%s): %v", command, sendErr) 
	}
}

func (rcon Rcon) Read() (response string){
	buffer := make([]byte, BufferSize)
    bytesRead, err := rcon.Connection.Read(buffer)
    if err != nil {
        fmt.Printf("Read err %v\n", err)
        os.Exit(-1)
    }

	if bytesRead >= 4 {
		infos := string(buffer[4:bytesRead])
		return infos
	} else {
		return ""
	}
}

func (rcon Rcon) RconCommand(command string) (res string) {
	if rcon.Connection != nil {
		rcon.Send(command)
		return rcon.Read()
	}
	return ""
}

func (rcon Rcon) RconCommandExtractValue(command string) string {
	fields := rcon.RconCommand(command)
	fields = strings.Replace(fields, "\n", " ", -1);
	tmpSplit := cleanEmptyLines(strings.Split(fields, " "))

	for _, elem := range tmpSplit {
		if strings.Contains(elem, "is:") {
			decolor_regex := regexp.MustCompile(`\^\d`)
			decolor := decolor_regex.ReplaceAllString(strings.Split(elem, "is:")[1], "")
			
			sign_regex := regexp.MustCompile(`"|,`)
			return sign_regex.ReplaceAllString(decolor, "")
		}
	}

	return ""
}

func (rcon *Rcon) CloseConnection() {
	fmt.Println("\nClosing connection ...")
	err := rcon.Connection.Close()

	if (err != nil) {
		fmt.Println("Error when closing connection. That's too bad !")
	} else {
		fmt.Println("Successfully closed connection.")
		rcon.Connection = nil
	}
}

func SplitReadInfos(readstr string) (responseType string, datas []string) {
	lines := cleanEmptyLines(strings.Split(readstr, "\n"))
	return lines[0], lines[1:]
}

func cleanEmptyLines(datas []string) []string { 
	var res []string
	for _, value := range(datas) {
		if value != "" {
			res = append(res, value)
		}
	}
	return res
}

func PrintSplitReadInfos(infos string) {
	fmt.Printf("\n~~~~~~~~~~ Print Read Infos ~~~~~~~~~~")
	cmd, datas := SplitReadInfos(infos)
	fmt.Printf("\nType: %s", cmd)
	datasLength := len(datas)
	if datasLength > 1 {
		fmt.Printf("\nLines: %d\n", len(datas))
	} else {
		fmt.Printf("\nLine: %d\n", len(datas))
	}
	for i, l := range(datas) {
		fmt.Printf("   |----> %2d) %s\n", i+1, l)
	}
}