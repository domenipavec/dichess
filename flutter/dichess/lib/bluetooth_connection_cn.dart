import 'dart:async';
import 'dart:typed_data';

import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';

import 'bluetooth_state_cn.dart';
import 'bluetoothpb/bluetoothpb.pb.dart';
import 'util/length_decoder.dart';

class BluetoothConnectionCN extends ChangeNotifier {
  BluetoothConnection _bluetoothConnection;
  BluetoothStateCN _bluetoothStateCN;

  bool isConnecting = true;
  Response latestResponse = Response();
  List<Response_WifiNetwork> networks = [];


  BluetoothConnectionCN(this._bluetoothStateCN) : super() {

    BluetoothConnection.toAddress(_bluetoothStateCN.connectedDevice.address).timeout(Duration(seconds: 10)).then((connection) {
      isConnecting = false;
      _bluetoothConnection = connection;
      notifyListeners();

      var _streamSubscrition = connection.input.transform(StreamTransformer.fromBind(lengthDecoder)).listen((r) {
        var response = Response.fromBuffer(r);
        switch (response.type) {
          case Response_Type.WIFI_UPDATE:
            networks = response.networks;
            break;
          case Response_Type.GAME_UPDATE:
            latestResponse = response;
            break;
        }
        notifyListeners();

//        if (response.hasChessBoard()) {
//          setState(() {
//            print(_chessController.game.load(response.chessBoard.fen));
//            print(_chessController.refreshBoard());
//          });
//        }
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
}