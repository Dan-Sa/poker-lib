package shared

import (
	"context"

	proto "github.com/Dan-Sa/poker-lib/pb"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "TECHNICAL_CHALLENGE_POKER_BOT",
	MagicCookieValue: "all-in",
}

var PluginMap = map[string]plugin.Plugin{
	"bot": &ConcretePlugin{},
}

// This is the interface that is exposing as a plugin.
type Plugin interface {
	Register(
		guid string,
		timeout_seconds uint,
	) (string, error)

	Action(
		players []*Player,
		hole_cards, board []*Card,
		pots []*Pot,
		bet_so_far, bet_to_player, min_raise uint,
	) (proto.Action, uint, error)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type ConcretePlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation
	Impl Plugin
}

func (p *ConcretePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPluginServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *ConcretePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewPluginClient(c)}, nil
}
