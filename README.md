# Plugin Poker library

To use in a pluign, get the library with:
```
go get github.com/Dan-Sa/poker-lib/shared
```

You may also need some of the constants defined within the protobuf file. These can be used with:
```
go get github.com/Dan-Sa/poker-lib/pb
```

## Points To Note

If a bet is made that exceeds the amount that the bot has to bet, it will be folded out of future rounds of betting and will not be elegable fo any winnings on that hand.

## Examples

### Match all bets
```go
package main

import (
	proto "github.com/Dan-Sa/poker-lib/pb"
	lib "github.com/Dan-Sa/poker-lib/shared"
)

type PokerBot struct{}

func (b *PokerBot) Action(
	players []*lib.Player,
	hole_cards, board []*lib.Card,
	pots []*lib.Pot,
	bet_so_far, bet_to_player_this_round, min_raise uint,
) (proto.Action, uint) {

	if bet_to_player_this_round > bet_so_far {
		return proto.Action_Bet, bet_to_player_this_round - bet_so_far
	}

	return proto.Action_Bet, min_raise
}

func main() {
	lib.HostBot(&lib.BotConfig{BotName: "Match-All"}, &PokerBot{})
}
```

### All In
```go
package main

import (
	proto "github.com/Dan-Sa/poker-lib/pb"
	lib "github.com/Dan-Sa/poker-lib/shared"
)

type PokerBot struct{}

func (b *PokerBot) Action(
	players []*lib.Player,
	hole_cards, board []*lib.Card,
	pots []*lib.Pot,
	bet_so_far, bet_to_player_this_round, min_raise uint,
) (proto.Action, uint) {

    amount := uint(0)
	for _, p := range players {
		if p.Name == "All-In" {
			amount = p.Bank
		}
	}

	return proto.Action_Bet, amount
}

func main() {
	lib.HostBot(&lib.BotConfig{BotName: "All-In"}, &PokerBot{})
}
```