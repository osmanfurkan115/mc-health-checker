package main

import (
	"time"
)

type HealthCheck struct {
	server  *Server
	at      time.Time
	success bool
}

type ScheduledHealthCheck struct {
	servers []*Server
	checks  []*HealthCheck
	every   time.Duration
}

func NewScheduledHealthCheck(servers []*Server, every time.Duration) *ScheduledHealthCheck {
	return &ScheduledHealthCheck{
		servers: servers,
		checks:  []*HealthCheck{},
		every:   every,
	}
}

func (s *ScheduledHealthCheck) Add(server *Server) {
	s.servers = append(s.servers, server)
}

func (s *ScheduledHealthCheck) Schedule() {
	for {
		for _, server := range s.servers {
			check := NewHealthCheck(server)
			s.checks = append(s.checks, check)

			if !check.IsHealthy() {
				//server.ReportHealth(false)
			} else {
				server.ReportHealth(true)
			}
		}

		time.Sleep(s.every)
	}
}

func NewHealthCheck(server *Server) *HealthCheck {
	checkedServer, err := GetServer(server.IP)
	healthy := err == nil && checkedServer.Online

	return &HealthCheck{
		server:  server,
		at:      time.Now(),
		success: healthy,
	}
}

func (h *HealthCheck) IsHealthy() bool {
	return h.success
}
