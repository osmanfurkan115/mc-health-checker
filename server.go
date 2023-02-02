package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Server struct {
	IP string `json:"ip"`

	Players struct {
		Online int `json:"online"`
	} `json:"players"`

	Online bool `json:"online"`
}

func GetServer(ip string) (*Server, error) {
	resp, err := http.Get("https://api.mcsrvstat.us/2/" + ip)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var server Server
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}

	return &server, err
}

func (s *Server) ReportHealth(h bool) {
	fmt.Printf("Server %s is healthy: %t", s.IP, h)
	fmt.Println()
}
