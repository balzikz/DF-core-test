package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world/healing"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.InfoLevel

	conf := server.DefaultConfig()
	conf.Network.Address = ":19132"

	srv := conf.New()

	cmd.Register(cmd.New("heal", "Восстанавливает полное здоровье", nil, Heal{}))

	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}

	for srv.Accept(func(p *player.Player) {
		p.Message("§eДобро пожаловать на DF-core-test! Код обновлен и работает.")
		p.Handle(NewPlayerHandler(p))
	}) {
	}
}

type Heal struct{}

func (Heal) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*player.Player); ok {
		p.Heal(p.MaxHealth(), healing.SourceCommand{})
		output.Printf("§aВы были полностью исцелены!")
	} else {
		output.Errorf("Эту команду может использовать только игрок.")
	}
}

type PlayerHandler struct {
	player.NopHandler
	p *player.Player
}

func NewPlayerHandler(p *player.Player) *PlayerHandler {
	return &PlayerHandler{p: p}
}

func (h *PlayerHandler) HandleQuit() {
	fmt.Printf("Игрок %s покинул сервер.\n", h.p.Name())
}
