import 'dart:async';
import 'dart:typed_data';

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';
import 'package:flutter_chess_board/flutter_chess_board.dart';

import 'bluetooth_state_cn.dart';
import 'bluetoothpb/bluetoothpb.pb.dart';
import 'util/length_decoder.dart';
import 'package:chess/chess.dart' as chess;


class BluetoothConnectionCN extends ChangeNotifier {
  BluetoothConnection _bluetoothConnection;
  BluetoothStateCN _bluetoothStateCN;

  bool isConnecting = true;
  Response latestResponse = Response();
  List<Response_WifiNetwork> networks = [];
  ChessBoardController chessBoardController = ChessBoardController();
  bool rotateBoard = false;
  bool canMakeMove = false;
  String state = "";



  BluetoothConnectionCN(this._bluetoothStateCN) : super() {
    chessBoardController.game = chess.Chess();

    BluetoothConnection.toAddress(_bluetoothStateCN.connectedDevice.address).timeout(Duration(seconds: 10)).then((connection) {
      isConnecting = false;
      _bluetoothConnection = connection;
      notifyListeners();

      var _streamSubscrition = connection.input.transform(StreamTransformer.fromBind(lengthDecoder)).listen((r) {
        var response = Response.fromBuffer(r);
        print(response);
        switch (response.type) {
          case Response_Type.WIFI_UPDATE:
            networks = response.networks;
            break;
          case Response_Type.GAME_UPDATE:
            latestResponse = response;
            if (response.hasChessBoard()) {
              rotateBoard = response.chessBoard.rotate;
              canMakeMove = response.chessBoard.canMakeMove;
              chessBoardController.game.load(response.chessBoard.fen);
            }
            if (response.state != "") {
              state = response.state;
            }
            break;
          case Response_Type.STATE_UPDATE:
            state = response.state;
            break;
        }
        notifyListeners();
      });

      connection.output.done.then((_) {
        _bluetoothStateCN.disconnect();
      });

      // Disconnected
      _streamSubscrition.onDone(() {
        _bluetoothStateCN.disconnect();
      });

    }).catchError((error) {
      _bluetoothStateCN.disconnect();
    });
  }

  void _send(Request request) {
    if (isConnecting) {
      return;
    }
    var data = request.writeToBuffer();

    var lengthData = ByteData.view(Uint8List(8).buffer);
    lengthData.setUint64(0, data.length);

    _bluetoothConnection.output.add(lengthData.buffer.asUint8List());
    _bluetoothConnection.output.add(data);
  }

  void startWifiScan() {
    var request = Request();
    request.type = Request_Type.START_WIFI_SCAN;
    _send(request);
  }

  void stopWifiScan() {
    var request = Request();
    request.type = Request_Type.STOP_WIFI_SCAN;
    _send(request);
  }

  void connectWifi(String ssid) {
    var request = Request();
    request.type = Request_Type.CONNECT_WIFI;
    request.wifiSsid = ssid;
    _send(request);
  }

  void forgetWifi(String ssid) {
    var request = Request();
    request.type = Request_Type.FORGET_WIFI;
    request.wifiSsid = ssid;
    _send(request);
  }

  void configureWifi(String ssid, String psk) {
    var request = Request();
    request.type = Request_Type.CONFIGURE_WIFI;
    request.wifiSsid = ssid;
    request.wifiPsk = psk;
    _send(request);
  }

  void updateSettings(void Function(Settings) update) {
    var request = Request();
    request.type = Request_Type.UPDATE_SETTINGS;
    request.settings = latestResponse.settings.copyWith(update);
    _send(request);
  }

  void makeMove(String move) {
    var request = Request();
    request.type = Request_Type.MOVE;
    request.move = move;
    _send(request);
  }

  void undoMove() {
    var request = Request();
    request.type = Request_Type.UNDO_MOVE;
    _send(request);
  }

  void newGame() {
    var request = Request();
    request.type = Request_Type.NEW_GAME;
    _send(request);
  }
}