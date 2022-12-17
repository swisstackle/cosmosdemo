package keeper_test

import (
	"github.com/stretchr/testify/require"
	"github.com/swisstackle/checkers/x/checkers/types"
	"testing"
)

const (
	alice = "cosmos1h43c9qmv04aj43r5pnts8wens6gc6jtxjfrn74"
	bob   = "cosmos1wl3xze8xvtz6la349kqdt3xuvccjq9xu2v96vw"
	carol = "cosmos1426zv69mav7glynccqwfhlm5we73l2njud2ljn"
)

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "",
	}, *createResponse)
}
