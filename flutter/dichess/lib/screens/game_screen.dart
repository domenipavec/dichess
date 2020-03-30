import 'package:flutter/material.dart';

import 'package:flutter_chess_board/flutter_chess_board.dart';

import '../bluetooth_connection_cn.dart';

class GameScreen extends StatelessWidget {
  const GameScreen({
    Key key,
    this.bluetoothConnectionCN,
  }) : super(key: key);

  final BluetoothConnectionCN bluetoothConnectionCN;

  @override
  Widget build(BuildContext context) {
    var moves = bluetoothConnectionCN.latestResponse.moves;
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: [
        Center(
          child: ChessBoard(
            size: 200,
            chessBoardController: bluetoothConnectionCN.chessBoardController,
            enableUserMoves: bluetoothConnectionCN.canMakeMove,
            whiteSideTowardsUser: !bluetoothConnectionCN.rotateBoard,
            onMove: (_) {
              var move = bluetoothConnectionCN.chessBoardController.game.getHistory()[0];
              bluetoothConnectionCN.makeMove(move);
            },
          ),
        ),
        Text(bluetoothConnectionCN.state),
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: EdgeInsets.all(3),
              child: RaisedButton(
                onPressed: () { bluetoothConnectionCN.undoMove(); },
                child: Text('Undo move'),
              ),
            ),
            Padding(
              padding: EdgeInsets.all(3),
              child: RaisedButton(
                onPressed: () { bluetoothConnectionCN.newGame(); },
                child: Text('New game'),
              ),
            ),
          ],
        ),
        Expanded(
          child: ListView.builder(
            reverse: true,
            itemCount: (moves.length/2).ceil(),
            itemBuilder: (context, idx) {
              idx = (moves.length/2).ceil() - idx - 1;
              var move1 = moves[2*idx];
              var move2 = "";
              if (moves.length > 2*idx+1) {
                move2 = moves[2*idx+1];
              }
              return Padding(
                padding: EdgeInsets.symmetric(vertical: 4, horizontal: 4),
                child: Text("${idx+1}. ${move1} ${move2}"),
              );
            },
          ),
        )
      ],
    );
  }
}
