package main

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.InfoLevel

	conf, err := server.LoadConfig(log)
	if err != nil {
		log.Fatalln(err)
	}
	srv := server.New(&conf, log)

	cmd.Register(cmd.New("heal", "Восстанавливает полное здоровье", nil, Heal{}))

	if err := srv.Start(); err != nil {
		log.Fatalln(err)
	}

	for srv.Accept(func(p *player.Player) {
		p.Message("§aДобро пожаловать! Финальная версия кода работает!")
		p.Handle(NewPlayerHandler(p))
	}) {
	}
}

type Heal struct{}

func (Heal) Run(src cmd.Source, output *cmd.Output, tx *world.Tx) {
	if p, ok := src.(*player.Player); ok {
		p.Heal(p.MaxHealth(), nil)
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

func (h *PlayerHandler) HandleQuit(p *player.Player) {
	fmt.Printf("Игрок %s покинул сервер.\n", p.Name())
}
