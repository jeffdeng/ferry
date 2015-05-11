package main

import (
	"../server"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Packet struct {
	Command string                 `json:"cmd"`
	Params  []string               `json:"params"`
	Sender  string                 `json:"sender"`
	Content map[string]interface{} `json:"content"`
}

type Server struct {
	Cache *server.Cache
}

var BizServer Server

func (s *Server) start() {
	var startTime int64 = time.Now().Unix()
	s.Cache = server.NewCache()
	s.Cache.SetValue("start-time", startTime)

	v, _ := s.Cache.GetIntValue("start-time")
	println(v)

	http.ListenAndServe(":8888", nil)
}

func (s *Server) setValue(packet Packet) {
	params := packet.Params
	if len(params) < 0 {
		return
	}
	/**/
	for k, v := range packet.Content {
		switch value := v.(type) {
		case string:
			print(k, value)
		case float64:
			print(k, int64(value))
		}
	}

	bytes, _ := json.Marshal(packet.Content)
	s.Cache.SetValue(params[0], string(bytes))
}

func (s *Server) getValue(packet Packet) (interface{}, error) {
	params := packet.Params
	if len(params) < 0 {
		return "", errors.New("Wrong number of params")
	}

	return s.Cache.GetValue(params[0])
}

func (s *Server) sendData(response http.ResponseWriter, data interface{}) {
	switch value := data.(type) {
	case string:
		response.Write([]byte(value))

	}
}

func mainHandler(response http.ResponseWriter, request *http.Request) {
	postedMsg, err := ioutil.ReadAll(request.Body)

	if err == nil {
		content := []byte(postedMsg)
		response.Write([]byte("{\"errorCode\": 0}"))
		// println(content)
		var packet Packet
		if err := json.Unmarshal(content, &packet); err == nil {
			if packet.Command == "set-value" {
				BizServer.setValue(packet)
			} else if packet.Command == "get-value" {
				value, _ := BizServer.getValue(packet)
				BizServer.sendData(response, value)
			} else if packet.Command == "get-value" {

			} else if packet.Command == "get-value" {

			}

		} else {
			println("Error")
		}
	}
}

func main() {
	http.HandleFunc("/main", mainHandler)
	BizServer.start()
}
