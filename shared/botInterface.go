package shared

import proto "github.com/Dan-Sa/poker-lib/pb"

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

func HostBot(conf *BotConfig, bot Bot) {

}
