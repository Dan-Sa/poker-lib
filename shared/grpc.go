package shared

import (
	"context"

	proto "github.com/Dan-Sa/poker-lib/pb"
)

type GRPCClient struct{ client proto.PluginClient }

func (m *GRPCClient) Register(
	guid string,
	timeout_seconds uint,
) (string, error) {
	resp, err := m.client.Register(context.Background(), &proto.RegisterRequest{
		PlayerGuid:             guid,
		DecisionTimeoutSeconds: uint64(timeout_seconds),
	})
	if err != nil {
		return "", err
	}

	return resp.BotName, nil
}

func (m *GRPCClient) Action(
	players []*Player,
	hole_cards, board []*Card,
	pots []*Pot,
	bet_so_far, bet_to_player, min_raise uint,
) (proto.Action, uint, error) {
	ar := &proto.ActionRequest{
		Players:     ParsePlayerListToProto(players),
		Pots:        ParsePotListToProto(pots),
		HoleCards:   ParseBoardToProto(hole_cards),
		BetSoFar:    uint64(bet_so_far),
		BetToPlayer: uint64(bet_to_player),
		MinRaise:    uint64(min_raise),
	}

	// Parse Cards
	if len(board) > 0 {
		parsed_board := ParseBoardToProto(board)
		ar.Flop = parsed_board[0:3]
		if len(board) > 3 {
			ar.Turn = parsed_board[3]
		}
		if len(board) > 4 {
			ar.River = parsed_board[4]
		}
	}

	resp, err := m.client.Action(context.Background(), ar)
	if err != nil {
		return proto.Action_UNKNOWN_Action, 0, err
	}

	return resp.Action, uint(resp.Amount), nil
}

type GRPCServer struct {
	proto.UnimplementedPluginServer
	// This is the real implementation
	Impl Plugin
}

func (m *GRPCServer) Register(
	ctx context.Context,
	req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	name, err := m.Impl.Register(req.PlayerGuid, uint(req.DecisionTimeoutSeconds))
	return &proto.RegisterResponse{BotName: name}, err
}

func (m *GRPCServer) Action(
	ctx context.Context,
	req *proto.ActionRequest) (*proto.ActionResponse, error) {
	// TODO: Take items from req and parse them into what is needed by Action: Proto -> struct

	action, amount, err := m.Impl.Action(
		ParsePlayerListToStruc(req.Players),
		ParseBoardToStruct(req.HoleCards),
		append(ParseBoardToStruct(req.Flop), ParseCardToStruct(req.Turn), ParseCardToStruct(req.River)),
		ParsePotListToStruct(req.Pots),
		uint(req.BetSoFar),
		uint(req.BetToPlayer),
		uint(req.MinRaise),
	)
	return &proto.ActionResponse{Action: action, Amount: uint64(amount)}, err
}
