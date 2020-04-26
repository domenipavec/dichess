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
var piecesStrings = map[bluetoothpb.Settings_Language]map[chess.PieceType][]string{
	bluetoothpb.Settings_SLOVENIAN: map[chess.PieceType][]string{
		chess.King:   []string{"kralj"},
		chess.Queen:  []string{"kraljica", "dama"},
		chess.Rook:   []string{"top", "trdnjava"},
		chess.Bishop: []string{"lovec", "tekač", "laufer"},
		chess.Knight: []string{"skakač", "konj"},
		chess.Pawn:   []string{"kmet"},
	},
	bluetoothpb.Settings_ENGLISH: map[chess.PieceType][]string{
		chess.King:   []string{"king"},
		chess.Queen:  []string{"queen"},
		chess.Rook:   []string{"rook", "castle"},
		chess.Bishop: []string{"bishop"},
		chess.Knight: []string{"knight"},
		chess.Pawn:   []string{"pawn"},
	},
}

var languages = map[bluetoothpb.Settings_Language]string{
	bluetoothpb.Settings_SLOVENIAN: "sl-SI",
	bluetoothpb.Settings_ENGLISH:   "en-US",
}
