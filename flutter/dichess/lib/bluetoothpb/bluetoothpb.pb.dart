///
//  Generated code. Do not modify.
//  source: bluetoothpb.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'bluetoothpb.pbenum.dart';

export 'bluetoothpb.pbenum.dart';

class Response_ChessBoard extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response.ChessBoard', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
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
  @$core.pragma('dart2js:noInline')
  static Response_ChessBoard getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Response_ChessBoard>(create);
  static Response_ChessBoard _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get fen => $_getSZ(0);
  @$pb.TagNumber(1)
  set fen($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFen() => $_has(0);
  @$pb.TagNumber(1)
  void clearFen() => clearField(1);
}

class Response_WifiNetwork extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response.WifiNetwork', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..aOS(1, 'ssid')
    ..aOB(2, 'connected')
    ..aOB(3, 'available')
    ..aOB(4, 'saved')
    ..aOB(5, 'connecting')
    ..aOB(6, 'failed')
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
  @$core.pragma('dart2js:noInline')
  static Response_WifiNetwork getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Response_WifiNetwork>(create);
  static Response_WifiNetwork _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get ssid => $_getSZ(0);
  @$pb.TagNumber(1)
  set ssid($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSsid() => $_has(0);
  @$pb.TagNumber(1)
  void clearSsid() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get connected => $_getBF(1);
  @$pb.TagNumber(2)
  set connected($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasConnected() => $_has(1);
  @$pb.TagNumber(2)
  void clearConnected() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get available => $_getBF(2);
  @$pb.TagNumber(3)
  set available($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAvailable() => $_has(2);
  @$pb.TagNumber(3)
  void clearAvailable() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get saved => $_getBF(3);
  @$pb.TagNumber(4)
  set saved($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasSaved() => $_has(3);
  @$pb.TagNumber(4)
  void clearSaved() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get connecting => $_getBF(4);
  @$pb.TagNumber(5)
  set connecting($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasConnecting() => $_has(4);
  @$pb.TagNumber(5)
  void clearConnecting() => clearField(5);

  @$pb.TagNumber(6)
  $core.bool get failed => $_getBF(5);
  @$pb.TagNumber(6)
  set failed($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasFailed() => $_has(5);
  @$pb.TagNumber(6)
  void clearFailed() => clearField(6);
}

class Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..aOM<Response_ChessBoard>(1, 'chessBoard', subBuilder: Response_ChessBoard.create)
    ..pc<Response_WifiNetwork>(2, 'networks', $pb.PbFieldType.PM, subBuilder: Response_WifiNetwork.create)
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
  @$core.pragma('dart2js:noInline')
  static Response getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Response>(create);
  static Response _defaultInstance;

  @$pb.TagNumber(1)
  Response_ChessBoard get chessBoard => $_getN(0);
  @$pb.TagNumber(1)
  set chessBoard(Response_ChessBoard v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasChessBoard() => $_has(0);
  @$pb.TagNumber(1)
  void clearChessBoard() => clearField(1);
  @$pb.TagNumber(1)
  Response_ChessBoard ensureChessBoard() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.List<Response_WifiNetwork> get networks => $_getList(1);
}

class Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Request', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..e<Request_Type>(1, 'type', $pb.PbFieldType.OE, defaultOrMaker: Request_Type.NOOP, valueOf: Request_Type.valueOf, enumValues: Request_Type.values)
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
  @$core.pragma('dart2js:noInline')
  static Request getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Request>(create);
  static Request _defaultInstance;

  @$pb.TagNumber(1)
  Request_Type get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(Request_Type v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get wifiSsid => $_getSZ(1);
  @$pb.TagNumber(2)
  set wifiSsid($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasWifiSsid() => $_has(1);
  @$pb.TagNumber(2)
  void clearWifiSsid() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get wifiPsk => $_getSZ(2);
  @$pb.TagNumber(3)
  set wifiPsk($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasWifiPsk() => $_has(2);
  @$pb.TagNumber(3)
  void clearWifiPsk() => clearField(3);
}

