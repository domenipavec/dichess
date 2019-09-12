///
//  Generated code. Do not modify.
//  source: bluetoothpb.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, pragma, String;

import 'package:protobuf/protobuf.dart' as $pb;

import 'bluetoothpb.pbenum.dart';

export 'bluetoothpb.pbenum.dart';

class Response_ChessBoard extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response.ChessBoard', package: const $pb.PackageName('bluetoothpb'))
    ..aOS(1, 'fen')
    ..hasRequiredFields = false
  ;

  Response_ChessBoard._() : super();
  factory Response_ChessBoard() => create();
  factory Response_ChessBoard.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Response_ChessBoard.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Response_ChessBoard clone() => Response_ChessBoard()..mergeFromMessage(this);
  Response_ChessBoard copyWith(void Function(Response_ChessBoard) updates) => super.copyWith((message) => updates(message as Response_ChessBoard));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Response_ChessBoard create() => Response_ChessBoard._();
  Response_ChessBoard createEmptyInstance() => create();
  static $pb.PbList<Response_ChessBoard> createRepeated() => $pb.PbList<Response_ChessBoard>();
  static Response_ChessBoard getDefault() => _defaultInstance ??= create()..freeze();
  static Response_ChessBoard _defaultInstance;

  $core.String get fen => $_getS(0, '');
  set fen($core.String v) { $_setString(0, v); }
  $core.bool hasFen() => $_has(0);
  void clearFen() => clearField(1);
}

class Response_WifiNetwork extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response.WifiNetwork', package: const $pb.PackageName('bluetoothpb'))
    ..aOS(1, 'ssid')
    ..hasRequiredFields = false
  ;

  Response_WifiNetwork._() : super();
  factory Response_WifiNetwork() => create();
  factory Response_WifiNetwork.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Response_WifiNetwork.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Response_WifiNetwork clone() => Response_WifiNetwork()..mergeFromMessage(this);
  Response_WifiNetwork copyWith(void Function(Response_WifiNetwork) updates) => super.copyWith((message) => updates(message as Response_WifiNetwork));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Response_WifiNetwork create() => Response_WifiNetwork._();
  Response_WifiNetwork createEmptyInstance() => create();
  static $pb.PbList<Response_WifiNetwork> createRepeated() => $pb.PbList<Response_WifiNetwork>();
  static Response_WifiNetwork getDefault() => _defaultInstance ??= create()..freeze();
  static Response_WifiNetwork _defaultInstance;

  $core.String get ssid => $_getS(0, '');
  set ssid($core.String v) { $_setString(0, v); }
  $core.bool hasSsid() => $_has(0);
  void clearSsid() => clearField(1);
}

class Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response', package: const $pb.PackageName('bluetoothpb'))
    ..a<Response_ChessBoard>(1, 'chessBoard', $pb.PbFieldType.OM, Response_ChessBoard.getDefault, Response_ChessBoard.create)
    ..pc<Response_WifiNetwork>(2, 'configuredNetworks', $pb.PbFieldType.PM,Response_WifiNetwork.create)
    ..pc<Response_WifiNetwork>(3, 'discoveredNetworks', $pb.PbFieldType.PM,Response_WifiNetwork.create)
    ..hasRequiredFields = false
  ;

  Response._() : super();
  factory Response() => create();
  factory Response.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Response.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Response clone() => Response()..mergeFromMessage(this);
  Response copyWith(void Function(Response) updates) => super.copyWith((message) => updates(message as Response));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Response create() => Response._();
  Response createEmptyInstance() => create();
  static $pb.PbList<Response> createRepeated() => $pb.PbList<Response>();
  static Response getDefault() => _defaultInstance ??= create()..freeze();
  static Response _defaultInstance;

  Response_ChessBoard get chessBoard => $_getN(0);
  set chessBoard(Response_ChessBoard v) { setField(1, v); }
  $core.bool hasChessBoard() => $_has(0);
  void clearChessBoard() => clearField(1);

  $core.List<Response_WifiNetwork> get configuredNetworks => $_getList(1);

  $core.List<Response_WifiNetwork> get discoveredNetworks => $_getList(2);
}

class Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Request', package: const $pb.PackageName('bluetoothpb'))
    ..e<Request_Type>(1, 'type', $pb.PbFieldType.OE, Request_Type.NOOP, Request_Type.valueOf, Request_Type.values)
    ..aOS(2, 'wifiSsid')
    ..aOS(3, 'wifiPsk')
    ..hasRequiredFields = false
  ;

  Request._() : super();
  factory Request() => create();
  factory Request.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Request.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Request clone() => Request()..mergeFromMessage(this);
  Request copyWith(void Function(Request) updates) => super.copyWith((message) => updates(message as Request));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Request create() => Request._();
  Request createEmptyInstance() => create();
  static $pb.PbList<Request> createRepeated() => $pb.PbList<Request>();
  static Request getDefault() => _defaultInstance ??= create()..freeze();
  static Request _defaultInstance;

  Request_Type get type => $_getN(0);
  set type(Request_Type v) { setField(1, v); }
  $core.bool hasType() => $_has(0);
  void clearType() => clearField(1);

  $core.String get wifiSsid => $_getS(1, '');
  set wifiSsid($core.String v) { $_setString(1, v); }
  $core.bool hasWifiSsid() => $_has(1);
  void clearWifiSsid() => clearField(2);

  $core.String get wifiPsk => $_getS(2, '');
  set wifiPsk($core.String v) { $_setString(2, v); }
  $core.bool hasWifiPsk() => $_has(2);
  void clearWifiPsk() => clearField(3);
}

