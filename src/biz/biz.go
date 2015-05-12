package main

import (
	"../server"
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
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
	Cache          *server.Cache
	DataConnection *sql.DB
}

var BizServer Server

func (s *Server) start() {
	var startTime int64 = time.Now().Unix()
	s.Cache = server.NewCache()
	s.Cache.SetValue("start-time", startTime)

	v, _ := s.Cache.GetIntValue("start-time")
	println(v)

	// Here
	db, err := sql.Open("mysql", "root:123@/xjy_main?charset=utf8")
	if err == nil {
		s.DataConnection = db
	} else {
		println("database initialize error : ", err.Error())
	}

	http.ListenAndServe(":8888", nil)
}

func (s *Server) setValue(packet Packet) {
	params := packet.Params
	if len(params) < 0 {
		return
	}
	/*
		for k, v := range packet.Content {
			switch value := v.(type) {
			case string:
				//print(k, value)
			case float64:
				//print(k, int64(value))
			}
		}
	*/

	bytes, _ := json.Marshal(packet.Content)
	s.Cache.SetValue(params[0], string(bytes))
}

func (s *Server) getValue(packet Packet) (string, error) {
	params := packet.Params
	if len(params) < 0 {
		return "", errors.New("Wrong number of params")
	}
	arr := []interface{}{}
	for _, v := range params {
		val, _ := s.Cache.GetValue(v)
		arr = append(arr, val)
	}
	print(":Len=")
	println(len(arr))
	bytes, _ := json.Marshal(arr)
	return string(bytes), nil
}

func (s *Server) query(packet Packet) ([]interface{}, error) {

	params := packet.Params
	if len(params) < 0 {
		return nil, errors.New("Wrong number of params")
	}

	queryStr := params[0]
	db := s.DataConnection

	rows, err := db.Query(queryStr)
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		_ = rows.Scan(scanArgs)
		//println(id, name)
	}
	return nil, nil
}

func (s *Server) sendData(response http.ResponseWriter, data string) {
	/*
		switch value := data.(type) {
		case string:

		}*/

	response.Write([]byte(data))
	println("----")
	println(data)
	println("----")
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
			} else if packet.Command == "query" {

			} else if packet.Command == "echo" {

			} else if packet.Command == "" {

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
