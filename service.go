package guerillaradio

import "encoding/json"
import "net"
import "fmt"
import "strconv"
import "os"
import "strings"
import "bufio"

func ServiceRequest(library *Library, num_lines int) (encoded_json []byte, err error) {
	lines := library.RetrieveLines(num_lines)
	encoded_json, err = json.Marshal(lines)
	return
}

func convertRequestToInt(request string) (num_lines int, err error) {
	stripped := strings.TrimSpace(request)
	num_lines_64, err := strconv.ParseInt(stripped, 10, 64)
	num_lines = int(num_lines_64)
	return
}

func handleConnection(library *Library,conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		request, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error Reading from connection %v\n", err)
			break
		}

		num_lines, err := convertRequestToInt(request)
		if err != nil {
			fmt.Printf("Error parsing Int %v\n", err)
			//TODO: Return Some sort of error back to the socket
			continue
		}

		json_bytes, err := ServiceRequest(library, num_lines)
		if err != nil {
			fmt.Printf("Error Receiving JSON %v\n", err)
			//TODO: Return some sort of error back to the socket
			continue
		}
		json_bytes = appendNewLine(json_bytes)
		_, err = conn.Write(json_bytes)
		if err != nil {
			fmt.Printf("Error Writing JSON %v\n", err)
		}
		fmt.Printf("Wrote %v bytes\n", len(json_bytes))

	}
	conn.Close()
}

func appendNewLine(jsonbytes []byte) (newlined []byte) {
	newline := []byte("\n")
	newlined = make([]byte,len(jsonbytes) + len(newline))
	copy(newlined,jsonbytes)
	copy(newlined[len(jsonbytes):],newline)
	return
}


func Listen(library *Library, network_interface string) {
	listener, err := net.Listen("tcp", network_interface)
	if err != nil {
		fmt.Printf("Error Listening on %v: %v\n", network_interface, err)
		os.Exit(1)
	}

	fmt.Printf("Listening on %v\n", network_interface)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting request %v\n", err)
			continue
		}
		fmt.Printf("Accepted Request from: %v\n",conn.RemoteAddr())
		go handleConnection(library,conn)
	}
}
