import 'dart:async';
import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';
import 'package:flutter_chess_board/flutter_chess_board.dart';

import 'package:dichess/bluetoothpb/bluetoothpb.pb.dart';
import 'package:dichess/util/length_decoder.dart';
import 'package:provider/provider.dart';

import '../bluetooth_connection_cn.dart';

class DeviceScreen extends StatefulWidget {
  @override
  _DeviceState createState() => _DeviceState();
}

class _DeviceState extends State<DeviceScreen> {
  ChessBoardController _chessController;
  ChessBoard _chessBoard;

  @override
  void initState() {
    super.initState();

    _chessController = ChessBoardController();
    _chessBoard = ChessBoard(
      enableUserMoves: false,
      chessBoardController: _chessController,
    );
  }


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
          body: _chessBoard,
        );
      },
    );
  }
}