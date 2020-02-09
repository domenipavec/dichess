import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../bluetooth_state_cn.dart';

class FindDevicesScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Consumer<BluetoothStateCN>(
      builder: (context, bluetoothStateCN, child) =>
          Scaffold(
            appBar: AppBar(
              title: bluetoothStateCN.isDiscovering ? Text('Looking for devices') : Text('Devices'),
            ),
            body: RefreshIndicator(
              onRefresh: () {
                bluetoothStateCN.startDiscovery();
                return Future.delayed(Duration(seconds: 1));
              },
              child: ListView.builder(
                physics: const AlwaysScrollableScrollPhysics(),
                itemCount: bluetoothStateCN.devices.length,
                itemBuilder: (context, index) {
                  if (bluetoothStateCN.devices[index].name != null) {
                    return ListTile(
                        title: Text(bluetoothStateCN.devices[index].name),
                        onTap: () {
                          bluetoothStateCN.connect(bluetoothStateCN.devices[index]);
                        }
                    );
                  } else {
                    return ListTile(
                      title: Text(bluetoothStateCN.devices[index].address),
                    );
                  }
                },
              ),
            ),
          ),
    );
  }
}