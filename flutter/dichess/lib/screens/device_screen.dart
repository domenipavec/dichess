import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_chess_board/flutter_chess_board.dart';
import 'package:provider/provider.dart';

import '../bluetooth_connection_cn.dart';
import '../bluetoothpb/bluetoothpb.pb.dart';

class DeviceScreen extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    return Consumer<BluetoothConnectionCN>(
      builder: (context, bluetoothConnectionCN, child)
      {
        if (bluetoothConnectionCN.isConnecting) {
          return Scaffold(
            backgroundColor: Colors.lightBlue,
            body: Center(
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: <Widget>[
                  Icon(
                    Icons.bluetooth_searching,
                    size: 200.0,
                    color: Colors.white54,
                  ),
                ],
              ),
            ),
          );
        }
        Widget body;
        if (bluetoothConnectionCN.latestResponse.gameInProgress) {
          var moves = bluetoothConnectionCN.latestResponse.moves;
          body = Column(
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
        } else {
          var settings = bluetoothConnectionCN.latestResponse.settings;
          body = ListView(
            children: <Widget>[
              SwitchListTile(
                title: Text("Random white player selection"),
                value: settings.randomBw,
                onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.randomBw = value; }); },
              ),
              ListTile(
                title: Text(settings.randomBw ? "Player 1" : "White"),
                trailing: DropdownButton(
                  value: settings.player1,
                  items: Settings_PlayerType.values.map((value) => DropdownMenuItem(
                    value: value,
                    child: Text(value.toString()),
                  )).toList(),
                  onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.player1 = value; }); },
                ),
              ),
              ListTile(
                title: Text(settings.randomBw ? "Player 2" : "Black"),
                trailing: DropdownButton(
                  value: settings.player2,
                  items: Settings_PlayerType.values.map((value) => DropdownMenuItem(
                    value: value,
                    child: Text(value.toString()),
                  )).toList(),
                  onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.player2 = value; }); },
                ),
              ),
              SwitchListTile(
                title: Text("Sound output"),
                value: settings.sound,
                onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.sound = value; }); },
              ),
              SwitchListTile(
                title: Text("Voice recognition"),
                value: settings.voiceRecognition,
                onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.voiceRecognition = value; }); },
              ),
              SwitchListTile(
                title: Text("Auto move"),
                value: settings.autoMove,
                onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.autoMove = value; }); },
              ),
              ListTile(
                title: Text("Computer time limit"),
                trailing: Container(
                  width: 100,
                  child: TextField(
                    keyboardType: TextInputType.number,
                    inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                    textAlign: TextAlign.right,
                    controller: TextEditingController(text: settings.computerSettings.timeLimitMs.toString()),
                    onSubmitted: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.computerSettings.timeLimitMs = int.parse(value); }); },
                  ),
                ),
              ),
              SwitchListTile(
                title: Text("Computer limit strength"),
                value: settings.computerSettings.limitStrength,
                onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.computerSettings.limitStrength = value; }); },
              ),
              ListTile(
                enabled: settings.computerSettings.limitStrength,
                title: Text("Computer Elo"),
                trailing: Container(
                  width: 100,
                  child: TextField(
                    enabled: settings.computerSettings.limitStrength,
                    keyboardType: TextInputType.number,
                    inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                    textAlign: TextAlign.right,
                    controller: TextEditingController(text: settings.computerSettings.elo.toString()),
                    onSubmitted: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.computerSettings.elo = int.parse(value); }); },
                  ),
                ),
              ),
              ListTile(
                title: Text(bluetoothConnectionCN.state),
              ),
            ],
          );
        }
        return Scaffold(
          appBar: AppBar(
              title: Text("dichess"),
              actions: [
                IconButton(
                    icon: Icon(Icons.settings),
                    onPressed: () {
                      Navigator.pushNamed(context, "/settings");
                    }
                )
              ]
          ),
          body: body,
        );
      },
    );
  }
}