package guerillaradio

import "testing"
import "bufio"
import "net"
import "encoding/json"

func TestCanJsonifyLibraryRequest(t *testing.T) {

	library := Library{}
	err := library.AddSourceDocument("fixtures/oneline.txt")

	if err != nil {
		t.Errorf("Error while loading library: %v", err)
	}

	json, err := ServiceRequest(&library, 1)

	if err != nil {
		t.Errorf("Error while converting json: %v", err)
	}

	expected_json := "[\"first line\\n\"]"

	if expected_json != string(json) {
		t.Errorf("Converted Document: %v did not match %v", string(json), expected_json)
	}
}

func TestCanConvertFromRequestToInt(t *testing.T) {
	num_lines, err := convertRequestToInt("5\n")
	if err != nil {
		t.Errorf("Could not convert string to int: %v", err)
	}

	if 5 != num_lines {
		t.Errorf("Could not convert string to 5")
	}
}

func TestCanServeDocument(t *testing.T) {
	library := Library{}
	library.AddDirectory("fixtures")
	network_interface := ":9876"
	go Listen(&library, network_interface)

	conn, err := net.Dial("tcp", network_interface)
	if err != nil {
		t.Errorf("Error connecting to %v: %v\n", network_interface, err)
	}

	_, err = conn.Write([]byte("2\n"))
	if err !=nil {
		t.Errorf("Error writing to %v: %v", network_interface, err)
	}

	reader := bufio.NewReader(conn)
	json_bytes, err := reader.ReadString('\n')

	if err != nil {
		t.Errorf("Error while reading from server. json_bytes= %v, err= %v",json_bytes,err)
	}

	lines := make([]string,2)

	err = json.Unmarshal([]byte(json_bytes), &lines)
	if err != nil {
		t.Errorf("Error Decoding JSON: %v, %v",json_bytes, err)
	}
	if 2 != len(lines) {
		t.Errorf("Length of Returned Array was %v not 2",len(lines))
	}
}
