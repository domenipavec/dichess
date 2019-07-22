import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';

import 'package:dichess/screens/device_screen.dart';

class FindDevicesScreen extends StatefulWidget {
  const FindDevicesScreen({Key key}) : super(key: key);

  @override
  _FindDevicesState createState() => _FindDevicesState();
}

class _FindDevicesState extends State<FindDevicesScreen> {
  List<BluetoothDevice> devices = [];
  bool isDiscovering = false;

  Future<void> _restartDiscovery() {
    setState(() {
      isDiscovering = true;
      devices.clear();
    });

//    FlutterBluetoothSerial.instance.getBondedDevices().then((bondedDevices) {
//      setState(() {
//        bondedDevices.forEach((bondedDevice) {
//          devices.add(bondedDevice);
//        });
//      });
//    });

    var _streamSubscription = FlutterBluetoothSerial.instance.startDiscovery().listen((r) {
      setState(() {
        devices.add(r.device);
      });
    });

    _streamSubscription.onDone(() {
      setState(() {
        isDiscovering = false;
      });
    });

    return _streamSubscription.asFuture();
  }

  @override
  void initState() {
    super.initState();

    _restartDiscovery();
  }

  @override
  void dispose() {
    FlutterBluetoothSerial.instance.cancelDiscovery();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: isDiscovering ? Text('Looking for devices') : Text('Devices'),
      ),
      body: RefreshIndicator(
        onRefresh: _restartDiscovery,
        child: ListView.builder(
          physics: const AlwaysScrollableScrollPhysics(),
          itemCount: devices.length,
          itemBuilder: (context, index) {
            if (devices[index].name != null) {
              return ListTile(
                title: Text(devices[index].name),
                onTap: () {
                  Navigator.push(context, MaterialPageRoute(builder: (BuildContext context) {
                    return DeviceScreen(device: devices[index]);
                  }));
                }
              );
            } else {
              return ListTile(
                title: Text(devices[index].address),
              );
            }
          },
        ),
      ),
    );
  }
}