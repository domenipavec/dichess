package hardware

import "github.com/notnil/chess"

var forwardStrength = map[chess.PieceType]uint8{
	chess.Pawn:   70,
	chess.King:   70,
	chess.Queen:  80,
	chess.Bishop: 100,
	chess.Knight: 80,
	chess.Rook:   60,
}

var sidewayStrength = map[chess.PieceType]uint8{
	chess.Pawn:   80,
	chess.King:   130,
	chess.Queen:  140,
	chess.Bishop: 190,
	chess.Knight: 190,
	chess.Rook:   100,
}

var rotateStrength = map[chess.PieceType]uint8{
	chess.Pawn:   80,
	chess.King:   140,
	chess.Queen:  100,
	chess.Bishop: 80,
	chess.Knight: 100,
	chess.Rook:   70,
}
