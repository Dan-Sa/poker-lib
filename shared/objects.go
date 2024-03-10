package shared

import (
	proto "github.com/Dan-Sa/poker-lib/pb"
)

type Player struct {
	Guid     string
	Name     string
	State    proto.State
	Bank     uint
	BetSoFar uint
}

func (p *Player) ParseToProto() *proto.Player {
	return &proto.Player{
		Guid:     p.Guid,
		Name:     p.Name,
		State:    p.State,
		Bank:     uint64(p.Bank),
		BetSoFar: uint64(p.BetSoFar),
	}
}

func ParsePlayerToStruct(p *proto.Player) *Player {
	if p == nil {
		return nil
	}
	return &Player{
		Guid:     p.Guid,
		Name:     p.Name,
		State:    p.State,
		Bank:     uint(p.Bank),
		BetSoFar: uint(p.BetSoFar),
	}
}

func ParsePlayerListToProto(in []*Player) []*proto.Player {
	return SliceMap(in, func(v *Player, _ int) *proto.Player { return v.ParseToProto() })
}

func ParsePlayerListToStruc(in []*proto.Player) []*Player {
	return SliceMap(in, func(v *proto.Player, _ int) *Player { return ParsePlayerToStruct(v) })
}

type Card struct {
	Rank proto.Rank
	Suit proto.Suit
}

func (c *Card) String() string {
	ranks := "~~23456789TJQKA"
	suits := "~sdch"
	return ranks[c.Rank:c.Rank+1] + suits[c.Suit:c.Suit+1]
}

func (c *Card) ParseToProto() *proto.Card {
	return &proto.Card{
		Rank: c.Rank,
		Suit: c.Suit,
	}
}

func ParseCardToStruct(c *proto.Card) *Card {
	if c == nil {
		return nil
	}
	if c.Rank == proto.Rank_UNKNOWN_Rank || c.Suit == proto.Suit_UNKNOWN_Suit {
		return nil
	}
	return &Card{
		Rank: c.Rank,
		Suit: c.Suit,
	}
}

func ParseBoardToProto(in []*Card) []*proto.Card {
	return SliceMap(in, func(v *Card, _ int) *proto.Card { return v.ParseToProto() })
}

func ParseBoardToStruct(in []*proto.Card) []*Card {
	return SliceMap(in, func(v *proto.Card, _ int) *Card { return ParseCardToStruct(v) })
}

type Pot struct {
	Size    uint
	Players []string
}

func (p *Pot) ParseToProto() *proto.Pot {
	return &proto.Pot{
		Size:    uint64(p.Size),
		Players: p.Players,
	}
}

func ParsePotToStruct(p *proto.Pot) *Pot {
	if p == nil {
		return nil
	}
	return &Pot{
		Size:    uint(p.Size),
		Players: p.Players,
	}
}

func ParsePotListToProto(in []*Pot) []*proto.Pot {
	return SliceMap(in, func(v *Pot, _ int) *proto.Pot { return v.ParseToProto() })
}

func ParsePotListToStruct(in []*proto.Pot) []*Pot {
	return SliceMap(in, func(v *proto.Pot, _ int) *Pot { return ParsePotToStruct(v) })
}
