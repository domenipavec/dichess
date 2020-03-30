package voice

import (
	"github.com/matematik7/dichess/go/bluetoothpb"
	"github.com/notnil/chess"
)

var translations = map[bluetoothpb.Settings_Language]map[string]string{
	bluetoothpb.Settings_SLOVENIAN: map[string]string{
		"to":        "na",
		"from":      "iz",
		"ambiguous": "ni določen. Premik iz",
	},
	bluetoothpb.Settings_ENGLISH: map[string]string{
		"to":        "to",
		"from":      "from",
		"ambiguous": "is ambiguous. Move from",
	},
}
var piecesStrings = map[bluetoothpb.Settings_Language]map[chess.Piece][]string{
	bluetoothpb.Settings_SLOVENIAN: map[chess.Piece][]string{
		chess.WhiteKing:   []string{"kralj"},
		chess.WhiteQueen:  []string{"kraljica", "dama"},
		chess.WhiteRook:   []string{"top", "trdnjava"},
		chess.WhiteBishop: []string{"lovec", "tekač", "laufer"},
		chess.WhiteKnight: []string{"skakač", "konj"},
		chess.WhitePawn:   []string{"kmet"},
	},
	bluetoothpb.Settings_ENGLISH: map[chess.Piece][]string{
		chess.WhiteKing:   []string{"king"},
		chess.WhiteQueen:  []string{"queen"},
		chess.WhiteRook:   []string{"rook", "castle"},
		chess.WhiteBishop: []string{"bishop"},
		chess.WhiteKnight: []string{"knight"},
		chess.WhitePawn:   []string{"pawn"},
	},
}

var languages = map[bluetoothpb.Settings_Language]string{
	bluetoothpb.Settings_SLOVENIAN: "sl-SI",
	bluetoothpb.Settings_ENGLISH:   "en-US",
}
