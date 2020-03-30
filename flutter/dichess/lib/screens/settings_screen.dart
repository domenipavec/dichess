import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';

import '../bluetooth_connection_cn.dart';
import '../bluetoothpb/bluetoothpb.pbenum.dart';

class SettingsScreen extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    Provider.of<BluetoothConnectionCN>(context, listen: false).getSettings();

    return Consumer<BluetoothConnectionCN>(
      builder: (context, bluetoothConnectionCN, child) {
        var settings = bluetoothConnectionCN.latestResponse.settings;
        return Scaffold(
            appBar: AppBar(
              title: Text("Settings"),
            ),
            body: ListView(
              children: [
                SwitchListTile(
                  title: Text("Sound output"),
                  value: settings.sound,
                  onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.sound = value; }); },
                ),
                SwitchListTile(
                  title: Text("Voice recognition"),
                  value: settings.voiceRecognition,
                  onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.voiceRecognition = value; }); },
                ),
                ListTile(
                  title: Text("Recognition language"),
                  trailing: DropdownButton(
                    value: settings.language,
                    onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.language = value; }); },
                    items: Settings_Language.values.map((value) => DropdownMenuItem(
                      value: value,
                      child: Text(value.toString()),
                    )).toList(),
                  ),
                ),
                SwitchListTile(
                  title: Text("Auto move"),
                  value: settings.autoMove,
                  onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.autoMove = value; }); },
                ),
                ListTile(
                  title: Text("Computer player parameters", style: TextStyle(fontWeight: FontWeight.bold)),
                ),
                ListTile(
                  title: Text("Time limit"),
                  trailing: Container(
                    width: 100,
                    child: TextField(
                      keyboardType: TextInputType.number,
                      inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                      textAlign: TextAlign.right,
                      controller: TextEditingController(text: settings.computerSettings.timeLimitMs.toString()),
                      onSubmitted: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.computerSettings.timeLimitMs = int.parse(value); }); },
                    ),
                  ),
                ),
                ListTile(
                  title: Text("Skill Level"),
                  trailing: Container(
                    width: 100,
                    child: TextField(
                      keyboardType: TextInputType.number,
                      inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                      textAlign: TextAlign.right,
                      controller: TextEditingController(text: settings.computerSettings.skillLevel.toString()),
                      onSubmitted: (value) { bluetoothConnectionCN.updateSettings((settings) {
                        var v = int.parse(value);
                        if (v < 0) {
                          v = 0;
                        } else if (v > 20) {
                          v = 20;
                        }
                        settings.computerSettings.skillLevel = v;
                      }); },
                    ),
                  ),
                ),
                SwitchListTile(
                  title: Text("Limit strength"),
                  value: settings.computerSettings.limitStrength,
                  onChanged: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.computerSettings.limitStrength = value; }); },
                ),
                ListTile(
                  enabled: settings.computerSettings.limitStrength,
                  title: Text("Elo"),
                  trailing: Container(
                    width: 100,
                    child: TextField(
                      enabled: settings.computerSettings.limitStrength,
                      keyboardType: TextInputType.number,
                      inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                      textAlign: TextAlign.right,
                      controller: TextEditingController(text: settings.computerSettings.elo.toString()),
                      onSubmitted: (value) { bluetoothConnectionCN.updateSettings((settings) { settings.computerSettings.elo = int.parse(value); }); },
                    ),
                  ),
                ),
              ],
            )
        );
      },
    );
  }
}