import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../bluetooth_connection_cn.dart';
import '../bluetoothpb/bluetoothpb.pb.dart';

class WifiScreen extends StatelessWidget {

  Future<String> _showPasswordDialog(BuildContext context) async {
    String password = "";
    return showDialog<String>(
      context: context,
      barrierDismissible: false,
      builder: (context) => AlertDialog(
        title: Text('Enter password:'),
        content: TextField(
          autofocus: true,
          decoration: InputDecoration(
            labelText: 'Password',
          ),
          onChanged: (value) {
            password = value;
          },
        ),
        actions: <Widget>[
          FlatButton(
            child: Text('Ok'),
            onPressed: () {
              Navigator.of(context).pop(password);
            },
          ),
        ],
      ),
    );
  }

  Future<bool> _showForgetDialog(BuildContext context) async {
    return showDialog<bool>(
      context: context,
      barrierDismissible: false,
      builder: (context) => AlertDialog(
        title: Text('Forget this network?'),
        actions: <Widget>[
          FlatButton(
            child: Text('Yes'),
            onPressed: () {
              Navigator.of(context).pop(true);
            },
          ),
          FlatButton(
            child: Text('No'),
            onPressed: () {
              Navigator.of(context).pop(false);
            },
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    Provider.of<BluetoothConnectionCN>(context, listen: false).startWifiScan();

    return Consumer<BluetoothConnectionCN>(
      builder: (context, bluetoothConnectionCN, child) => WillPopScope(
        onWillPop: () async {
          bluetoothConnectionCN.stopWifiScan();
          return true;
        },
        child: Scaffold(
          appBar: AppBar(
            title: Text("Wifi Settings"),
          ),
          body: RefreshIndicator(
            onRefresh: () {
              bluetoothConnectionCN.startWifiScan();
              return Future.delayed(Duration(seconds: 1));
            },
            child: ListView.builder(
              physics: const AlwaysScrollableScrollPhysics(),
              itemCount: bluetoothConnectionCN.latestResponse.networks.length,
              itemBuilder: (context, index) {
                var network = bluetoothConnectionCN.latestResponse.networks[index];

                Widget leading = null;
                if (network.connected) {
                  leading = Icon(Icons.network_wifi);
                } else if (network.available) {
                  leading = Icon(Icons.wifi);
                }

                Widget subtitle = null;
                if (network.connected) {
                  subtitle = Text("Connected");
                } else if (network.failed) {
                  subtitle = Text("Connection failed");
                } else if (network.connecting) {
                  subtitle = Text("Connecting");
                } else if (network.saved) {
                  subtitle = Text("Saved");
                }

                return ListTile(
                  leading: leading,
                  title: Text(network.ssid),
                  subtitle: subtitle,
                  selected: network.connected || network.connecting,
                  onTap: () async {
                    if (network.available && (!network.saved || network.failed)) {
                      var password = await _showPasswordDialog(context);
                      bluetoothConnectionCN.configureWifi(network.ssid, password);
                    } else if (network.available && !network.connected && network.saved) {
                      bluetoothConnectionCN.connectWifi(network.ssid);
                    }
                  },
                  onLongPress: () async {
                    if (network.saved) {
                      var forget = await _showForgetDialog(context);
                      if (forget) {
                        bluetoothConnectionCN.forgetWifi(network.ssid);
                      }
                    }
                  },
                );
              },
            ),
          ),
        ),
      ),
    );
  }
}