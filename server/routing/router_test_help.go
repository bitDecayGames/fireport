package routing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/bitdecaygames/fireport/server/services"
)

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

func get(port int, endpoint string, data []byte) ([]byte, error) {
	return doHTTPReq(http.MethodGet, port, endpoint, data)
}

func post(port int, endpoint string, data []byte) ([]byte, error) {
	return doHTTPReq(http.MethodPost, port, endpoint, data)
}

func put(port int, endpoint string, data []byte) ([]byte, error) {
	return doHTTPReq(http.MethodPut, port, endpoint, data)
}

func doHTTPReq(method string, port int, endpoint string, data []byte) ([]byte, error) {
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
