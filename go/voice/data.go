package voice

import "github.com/notnil/chess"

var translations = map[string]map[string]string{
	"sl-SI": map[string]string{
		"to":        "na",
		"from":      "iz",
		"ambiguous": "ni določen. Premik iz",
	},
	"en-US": map[string]string{
		"to":        "to",
		"from":      "from",
		"ambiguous": "is ambiguous. Move from",
	},
}
var piecesStrings = map[string]map[chess.Piece][]string{
	"sl-SI": map[chess.Piece][]string{
		chess.WhiteKing:   []string{"kralj"},
		chess.WhiteQueen:  []string{"kraljica", "dama"},
		chess.WhiteRook:   []string{"top", "trdnjava"},
		chess.WhiteBishop: []string{"lovec", "tekač", "laufer"},
		chess.WhiteKnight: []string{"skakač", "konj"},
		chess.WhitePawn:   []string{"kmet"},
	},
	"en-US": map[chess.Piece][]string{
		chess.WhiteKing:   []string{"king"},
		chess.WhiteQueen:  []string{"queen"},
		chess.WhiteRook:   []string{"rook", "castle"},
		chess.WhiteBishop: []string{"bishop"},
		chess.WhiteKnight: []string{"knight"},
		chess.WhitePawn:   []string{"pawn"},
	},
}
