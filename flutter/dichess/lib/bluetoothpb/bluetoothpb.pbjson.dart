///
//  Generated code. Do not modify.
//  source: bluetoothpb.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Settings$json = const {
  '1': 'Settings',
  '2': const [
    const {'1': 'sound', '3': 1, '4': 1, '5': 8, '10': 'sound'},
    const {'1': 'language', '3': 2, '4': 1, '5': 14, '6': '.bluetoothpb.Settings.Language', '10': 'language'},
    const {'1': 'voice_recognition', '3': 3, '4': 1, '5': 8, '10': 'voiceRecognition'},
    const {'1': 'auto_move', '3': 4, '4': 1, '5': 8, '10': 'autoMove'},
    const {'1': 'random_bw', '3': 5, '4': 1, '5': 8, '10': 'randomBw'},
    const {'1': 'player1', '3': 6, '4': 1, '5': 14, '6': '.bluetoothpb.Settings.PlayerType', '10': 'player1'},
    const {'1': 'player2', '3': 7, '4': 1, '5': 14, '6': '.bluetoothpb.Settings.PlayerType', '10': 'player2'},
    const {'1': 'computerSettings', '3': 8, '4': 1, '5': 11, '6': '.bluetoothpb.Settings.ComputerSettings', '10': 'computerSettings'},
    const {'1': 'intro', '3': 9, '4': 1, '5': 8, '10': 'intro'},
  ],
  '3': const [Settings_ComputerSettings$json],
  '4': const [Settings_Language$json, Settings_PlayerType$json],
};

const Settings_ComputerSettings$json = const {
  '1': 'ComputerSettings',
  '2': const [
    const {'1': 'time_limit_ms', '3': 1, '4': 1, '5': 5, '10': 'timeLimitMs'},
    const {'1': 'skill_level', '3': 2, '4': 1, '5': 5, '10': 'skillLevel'},
    const {'1': 'limit_strength', '3': 3, '4': 1, '5': 8, '10': 'limitStrength'},
    const {'1': 'elo', '3': 4, '4': 1, '5': 5, '10': 'elo'},
  ],
};

const Settings_Language$json = const {
  '1': 'Language',
  '2': const [
    const {'1': 'ENGLISH', '2': 0},
    const {'1': 'SLOVENIAN', '2': 1},
  ],
};

const Settings_PlayerType$json = const {
  '1': 'PlayerType',
  '2': const [
    const {'1': 'HUMAN', '2': 0},
    const {'1': 'COMPUTER', '2': 1},
  ],
};

const Response$json = const {
  '1': 'Response',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.bluetoothpb.Response.Type', '10': 'type'},
    const {'1': 'networks', '3': 2, '4': 3, '5': 11, '6': '.bluetoothpb.Response.WifiNetwork', '10': 'networks'},
    const {'1': 'settings', '3': 3, '4': 1, '5': 11, '6': '.bluetoothpb.Settings', '10': 'settings'},
    const {'1': 'gameInProgress', '3': 4, '4': 1, '5': 8, '10': 'gameInProgress'},
    const {'1': 'moves', '3': 5, '4': 3, '5': 9, '10': 'moves'},
    const {'1': 'whiteTurn', '3': 6, '4': 1, '5': 8, '10': 'whiteTurn'},
    const {'1': 'state', '3': 7, '4': 1, '5': 9, '10': 'state'},
    const {'1': 'chess_board', '3': 8, '4': 1, '5': 11, '6': '.bluetoothpb.Response.ChessBoard', '10': 'chessBoard'},
  ],
  '3': const [Response_WifiNetwork$json, Response_ChessBoard$json],
  '4': const [Response_Type$json],
};

const Response_WifiNetwork$json = const {
  '1': 'WifiNetwork',
  '2': const [
    const {'1': 'ssid', '3': 1, '4': 1, '5': 9, '10': 'ssid'},
    const {'1': 'connected', '3': 2, '4': 1, '5': 8, '10': 'connected'},
    const {'1': 'available', '3': 3, '4': 1, '5': 8, '10': 'available'},
    const {'1': 'saved', '3': 4, '4': 1, '5': 8, '10': 'saved'},
    const {'1': 'connecting', '3': 5, '4': 1, '5': 8, '10': 'connecting'},
    const {'1': 'failed', '3': 6, '4': 1, '5': 8, '10': 'failed'},
  ],
};

const Response_ChessBoard$json = const {
  '1': 'ChessBoard',
  '2': const [
    const {'1': 'fen', '3': 1, '4': 1, '5': 9, '10': 'fen'},
    const {'1': 'rotate', '3': 2, '4': 1, '5': 8, '10': 'rotate'},
    const {'1': 'canMakeMove', '3': 3, '4': 1, '5': 8, '10': 'canMakeMove'},
  ],
};

const Response_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'NOOP', '2': 0},
    const {'1': 'GAME_UPDATE', '2': 1},
    const {'1': 'WIFI_UPDATE', '2': 2},
    const {'1': 'STATE_UPDATE', '2': 3},
  ],
};

const Request$json = const {
  '1': 'Request',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.bluetoothpb.Request.Type', '10': 'type'},
    const {'1': 'wifi_ssid', '3': 2, '4': 1, '5': 9, '10': 'wifiSsid'},
    const {'1': 'wifi_psk', '3': 3, '4': 1, '5': 9, '10': 'wifiPsk'},
    const {'1': 'settings', '3': 4, '4': 1, '5': 11, '6': '.bluetoothpb.Settings', '10': 'settings'},
    const {'1': 'move', '3': 5, '4': 1, '5': 9, '10': 'move'},
  ],
  '4': const [Request_Type$json],
};

const Request_Type$json = const {
  '1': 'Type',
  '2': const [
    const {'1': 'NOOP', '2': 0},
    const {'1': 'START_WIFI_SCAN', '2': 1},
    const {'1': 'STOP_WIFI_SCAN', '2': 2},
    const {'1': 'CONFIGURE_WIFI', '2': 3},
    const {'1': 'FORGET_WIFI', '2': 4},
    const {'1': 'CONNECT_WIFI', '2': 5},
    const {'1': 'UPDATE_SETTINGS', '2': 6},
    const {'1': 'UNDO_MOVE', '2': 7},
    const {'1': 'MOVE', '2': 8},
    const {'1': 'NEW_GAME', '2': 9},
    const {'1': 'GET_SETTINGS', '2': 10},
    const {'1': 'START_GAME', '2': 11},
  ],
};

