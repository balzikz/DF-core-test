package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.InfoLevel

	conf, err := server.DefaultConfig().Load()
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.New(&conf, log)
	srv.Start()

	for srv.Accept(func(p *player.Player) {
		p.Message("§eДобро пожаловать на наш Dragonfly сервер!")
		p.Handle(NewPlayerHandler(p))
	}) {
	}
}

type PlayerHandler struct {
	player.NopHandler
	p                 *player.Player
}


func NewPlayerHandler(p *player.Player) *PlayerHandler {
	return &PlayerHandler{p: p}
}

func (h *PlayerHandler) HandleCommandExecution(commandLine *string) {
	command := strings.Split(*commandLine, " ")[0]

	if command == "heal" {
		h.p.Heal(1000, "Исцеление командой /heal")
		h.p.Message("§aВы были полностью исцелены!")
		return
 }
	h.p.Messagef("§cКоманда '%s' не найдена.", command)
}
