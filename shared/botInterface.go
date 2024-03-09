package shared

import (
	proto "github.com/Dan-Sa/poker-lib/pb"
	"github.com/hashicorp/go-plugin"
)

type BotConfig struct {
	BotName string
}

type Bot interface {
	Action(
		players []*Player,
		hole_cards, board []*Card,
		pots []*Pot,
		bet_so_far, bet_to_player, min_raise uint,
	) (proto.Action, uint)
}

type BaseBot struct {
	Defined Bot
	Conf    *BotConfig
}

func (b *BaseBot) Register(
	guid string,
	timeout_seconds uint,
) (string, error) {
	return b.Conf.BotName, nil
}

func (b *BaseBot) Action(
	players []*Player,
	hole_cards, board []*Card,
	pots []*Pot,
	bet_so_far, bet_to_player, min_raise uint,
) (proto.Action, uint, error) {
	a, a2 := b.Defined.Action(players, hole_cards, board, pots, bet_so_far, bet_to_player, min_raise)
	return a, a2, nil
}

func HostBot(conf *BotConfig, bot Bot) {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			"bot": &ConcretePlugin{
				Impl: &BaseBot{
					Defined: bot,
					Conf:    conf,
				},
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
