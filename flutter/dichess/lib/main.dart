
import 'package:flutter/material.dart';

import 'package:provider/provider.dart';

import 'bluetooth_connection_cn.dart';
import 'screens/bluetooth_off.dart';
import 'screens/device_screen.dart';
import 'screens/find_devices.dart';
import 'bluetooth_state_cn.dart';
import 'screens/settings_screen.dart';
import 'screens/wifi_screen.dart';


void main() => runApp(
  ChangeNotifierProvider(
    create: (context) => BluetoothStateCN(),
    child: DIChessApp(),
  ),
);

class DIChessApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Consumer<BluetoothStateCN>(
        builder: (context, bluetoothStateCN, child) {
          if (!bluetoothStateCN.isBluetoothOn) {
            return BluetoothOffScreen();
          }
          if (!bluetoothStateCN.isConnected) {
            return FindDevicesScreen();
          }
          return ChangeNotifierProvider(
            create: (context) => BluetoothConnectionCN(bluetoothStateCN),
            child: MaterialApp(
                routes: {
                  "/": (BuildContext context) => DeviceScreen(),
                  "/wifi": (BuildContext context) => WifiScreen(),
                  "/settings": (BuildContext context) => SettingsScreen(),
                }
            ),
          );
        },
      ),
    );
  }
}