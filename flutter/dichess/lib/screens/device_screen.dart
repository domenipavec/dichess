import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../bluetooth_connection_cn.dart';
import 'game_screen.dart';
import 'start_screen.dart';

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
          body = GameScreen(
            bluetoothConnectionCN: bluetoothConnectionCN,
          );
        } else {
          body = StartScreen(
            settings: bluetoothConnectionCN.latestResponse.settings,
            state: bluetoothConnectionCN.state,
            update: bluetoothConnectionCN.updateSettings,
          );
        }
        return Scaffold(
          appBar: AppBar(
            title: Text("dichess"),
            actions: [
              PopupMenuButton(
                onSelected: (item) {
                  switch (item) {
                    case "wifi":
                      {
                        Navigator.pushNamed(context, "/wifi");
                      }
                      break;
                    case "settings":
                      {
                        Navigator.pushNamed(context, "/settings");
                      }
                      break;
                  }
                },
                itemBuilder: (context) {
                  return [
                    PopupMenuItem(
                      value: "settings",
                      child: Text("Settings"),
                    ),
                    PopupMenuItem(
                      value: "wifi",
                      child: Text("Wifi config"),
                    ),
                  ];
                },
              ),
            ],
          ),
          body: body,
        );
      },
    );
  }
}