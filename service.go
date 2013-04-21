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
			fmt.Println("Error Reading from connection ", err)
			break
		}

		num_lines, err := convertRequestToInt(request)
		if err != nil {
			fmt.Println("Error parsing Int ", err)
			//TODO: Return Some sort of error back to the socket
			continue
		}

		json_bytes, err := ServiceRequest(library, num_lines)
		if err != nil {
			fmt.Println("Error Receiving JSON ", err)
			//TODO: Return some sort of error back to the socket
			continue
		}

		_, err = conn.Write(json_bytes)
		if err != nil {
			fmt.Println("Error Writing JSON ", err)
		}

	}
	conn.Close()
}

func Listen(library *Library, network_interface string) {
	listener, err := net.Listen("tcp", network_interface)
	if err != nil {
		fmt.Println("Error Listening on ", network_interface)
		os.Exit(1)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting request ", err)
			continue
		}
		go handleConnection(library,conn)
	}
}
