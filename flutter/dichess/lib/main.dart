import 'package:async/async.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';

import 'package:dichess/screens/bluetooth_off.dart';
import 'package:dichess/screens/find_devices.dart';

void main() => runApp(DIChessApp());

class DIChessApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      color: Colors.lightBlue,
      home: StreamBuilder<BluetoothState>(
          stream: StreamGroup.merge([FlutterBluetoothSerial.instance.onStateChanged(), FlutterBluetoothSerial.instance.state.asStream()]),
          initialData: BluetoothState.UNKNOWN,
          builder: (c, snapshot) {
            final state = snapshot.data;
            if (state == BluetoothState.STATE_ON) {
              return FindDevicesScreen();
            }
            return BluetoothOffScreen(state: state);
          }),
    );
  }
}