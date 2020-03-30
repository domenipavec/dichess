import 'package:flutter/material.dart';

import '../bluetoothpb/bluetoothpb.pb.dart';


class StartScreen extends StatelessWidget {

  const StartScreen({
    Key key,
    this.settings,
    this.state,
    this.update,
  }) : super(key: key);

  final Settings settings;
  final String state;
  final Function update;

  @override
  Widget build(BuildContext context) {
    return ListView(
      children: <Widget>[
        SwitchListTile(
          title: Text("Random white player selection"),
          value: settings.randomBw,
          onChanged: (value) { update((settings) { settings.randomBw = value; }); },
        ),
        ListTile(
          title: Text(settings.randomBw ? "Player 1" : "White"),
          trailing: DropdownButton(
            value: settings.player1,
            items: Settings_PlayerType.values.map((value) => DropdownMenuItem(
              value: value,
              child: Text(value.toString()),
            )).toList(),
            onChanged: (value) { update((settings) { settings.player1 = value; }); },
          ),
        ),
        ListTile(
          title: Text(settings.randomBw ? "Player 2" : "Black"),
          trailing: DropdownButton(
            value: settings.player2,
            items: Settings_PlayerType.values.map((value) => DropdownMenuItem(
              value: value,
              child: Text(value.toString()),
            )).toList(),
            onChanged: (value) { update((settings) { settings.player2 = value; }); },
          ),
        ),
        SwitchListTile(
          title: Text("Intro"),
          value: settings.intro,
          onChanged: (value) { update((settings) { settings.intro = value; }); },
        ),
        ListTile(
          title: Text(state),
        ),
      ],
    );
  }
}
