package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/swisstackle/checkers/x/checkers/rules"
	"github.com/swisstackle/checkers/x/checkers/types"
	"strings"
	"testing"
)

const (
	bob   = "cosmos1dauuytp2temg5pdskcws60tt7ns8dquwrffm0k"
	alice = "cosmos1rk0c0gufn50el8stwtp7keq3vj380qn6pvkpm3"
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
		Black: alice,
		Red:   bob,
	}
}

func TestCanGetAddressBlack(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, aliceAddress, black)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1rk0c0gufn50el8stwtp7keq3vj380qn6pvkpm4"
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1rk0c0gufn50el8stwtp7keq3vj380qn6pvkpm4: decoding bech32 failed: invalid checksum (expected pvkpm3 got pvkpm4)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressRed(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	red, err2 := GetStoredGame1().GetRedAddress()
	require.Equal(t, bobAddress, red)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrongRed(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1dauuytp2temg5pdskcws60tt7ns8dquwrffm0b"
	red, err := storedGame.GetRedAddress()
	require.Nil(t, red)
	require.EqualError(t,
		err,
		"red address is invalid: cosmos1dauuytp2temg5pdskcws60tt7ns8dquwrffm0b: decoding bech32 failed: invalid character not part of charset: 98")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameCorrect(t *testing.T) {
	game, err := GetStoredGame1().ParseGame()
	require.EqualValues(t, rules.New().Pieces, game.Pieces)
	require.Nil(t, err)
}
func TestParseGameCanIfChangedOk(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Board = strings.Replace(storedGame.Board, "b", "r", 1)
	game, err := storedGame.ParseGame()
	require.NotEqualValues(t, rules.New().Pieces, game)
	require.Nil(t, err)
}
func TestParseGameWrongPieceColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Board = strings.Replace(storedGame.Board, "b", "w", 1)
	game, err := storedGame.ParseGame()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: invalid board, invalid piece at 1, 0")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestParseGameWrongTurnColor(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Turn = "w"
	game, err := storedGame.ParseGame()
	require.Nil(t, game)
	require.EqualError(t, err, "game cannot be parsed: Turn w")
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestGameValidateOk(t *testing.T) {
	storedGame := GetStoredGame1()
	require.NoError(t, storedGame.Validate())
}
