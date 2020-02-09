///
//  Generated code. Do not modify.
//  source: bluetoothpb.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const Response$json = const {
  '1': 'Response',
  '2': const [
    const {'1': 'chess_board', '3': 1, '4': 1, '5': 11, '6': '.bluetoothpb.Response.ChessBoard', '10': 'chessBoard'},
    const {'1': 'networks', '3': 2, '4': 3, '5': 11, '6': '.bluetoothpb.Response.WifiNetwork', '10': 'networks'},
  ],
  '3': const [Response_ChessBoard$json, Response_WifiNetwork$json],
};

const Response_ChessBoard$json = const {
  '1': 'ChessBoard',
  '2': const [
    const {'1': 'fen', '3': 1, '4': 1, '5': 9, '10': 'fen'},
  ],
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

const Request$json = const {
  '1': 'Request',
  '2': const [
    const {'1': 'type', '3': 1, '4': 1, '5': 14, '6': '.bluetoothpb.Request.Type', '10': 'type'},
    const {'1': 'wifi_ssid', '3': 2, '4': 1, '5': 9, '10': 'wifiSsid'},
    const {'1': 'wifi_psk', '3': 3, '4': 1, '5': 9, '10': 'wifiPsk'},
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
  ],
};

