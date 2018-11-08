package routing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"testing"

	"github.com/bitdecaygames/fireport/server/services"
)

// Unmarshal will read all bytes from r and populate the given obj, returning
// any error that may occur
func Unmarshal(r io.ReadCloser, obj interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		return err
	}

	return nil
}

func startTestServer() (int, *services.MasterList) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	port := listener.Addr().(*net.TCPAddr).Port

	svcs := services.NewMasterList()

	go serveInternal(listener, svcs)
	return port, svcs
}

func dumpResponse(t *testing.T, r *http.Response) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Logf("failed to read response: %v", err)
		return
	}

	t.Logf("response body: %v", string(bytes))
}

func get(port int, endpoint string, msg interface{}) ([]byte, error) {
	return doHTTPReq(http.MethodGet, port, endpoint, msg)
}

func post(port int, endpoint string, msg interface{}) ([]byte, error) {
	return doHTTPReq(http.MethodPost, port, endpoint, msg)
}

func put(port int, endpoint string, msg interface{}) ([]byte, error) {
	return doHTTPReq(http.MethodPut, port, endpoint, msg)
}

func doHTTPReq(method string, port int, endpoint string, msg interface{}) ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return []byte{}, err
	}

	fmt.Println("raw response data: ", string(data))

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("http://127.0.0.1:%v%v", port, endpoint),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return []byte{}, err
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte{}, err
	}

	if r.Status != "200 OK" {
		return body, fmt.Errorf(r.Status)
	}

	return body, nil
}
