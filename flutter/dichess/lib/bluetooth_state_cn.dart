import 'package:async/async.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';

class BluetoothStateCN extends ChangeNotifier {
  bool isBluetoothOn = false;

  List<BluetoothDevice> devices = [];
  bool isDiscovering = false;

  BluetoothDevice connectedDevice;
  bool isConnected = false;

  BluetoothStateCN() : super() {
    StreamGroup.merge([
      FlutterBluetoothSerial.instance.onStateChanged(),
      FlutterBluetoothSerial.instance.state.asStream()
    ]).listen((BluetoothState state) {
      bool isOn = state == BluetoothState.STATE_ON;
      if (isOn != isBluetoothOn) {
        isBluetoothOn = isOn;
        if (isBluetoothOn) {
          startDiscovery();
        }
        notifyListeners();
      }
    });
  }

  void _add(BluetoothDevice d) {
    if (d.name != "dichess") {
      return;
    }
    if (devices.any((existing) => existing.address == d.address)) {
      return;
    }
    devices.add(d);
    notifyListeners();
  }

  void _clear() {
    isDiscovering = true;
    devices.clear();
    notifyListeners();
  }

  void startDiscovery() {
    _clear();

    FlutterBluetoothSerial.instance.cancelDiscovery();
    var _streamSubscription = FlutterBluetoothSerial.instance.startDiscovery().listen((r) {
      _add(r.device);
    });

    _streamSubscription.onDone(() {
      isDiscovering = false;
      notifyListeners();
    });
  }

  void connect(BluetoothDevice device) {
    isConnected = true;
    connectedDevice = device;
    notifyListeners();
  }

  void disconnect() {
    isConnected = false;
    connectedDevice = null;
    notifyListeners();
  }
}