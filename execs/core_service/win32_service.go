package main

import (
	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	// Hier wird der Dienst gestartet
	return nil
}

func (p *program) Stop(s service.Service) error {
	// Hier wird der Dienst gestoppt
	return nil
}
