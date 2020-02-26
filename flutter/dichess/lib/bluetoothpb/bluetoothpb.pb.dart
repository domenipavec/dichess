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

class Settings_ComputerSettings extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Settings.ComputerSettings', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..a<$core.int>(1, 'timeLimitMs', $pb.PbFieldType.O3)
    ..a<$core.int>(2, 'skillLevel', $pb.PbFieldType.O3)
    ..aOB(3, 'limitStrength')
    ..a<$core.int>(4, 'elo', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  Settings_ComputerSettings._() : super();
  factory Settings_ComputerSettings() => create();
  factory Settings_ComputerSettings.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Settings_ComputerSettings.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Settings_ComputerSettings clone() => Settings_ComputerSettings()..mergeFromMessage(this);
  Settings_ComputerSettings copyWith(void Function(Settings_ComputerSettings) updates) => super.copyWith((message) => updates(message as Settings_ComputerSettings));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Settings_ComputerSettings create() => Settings_ComputerSettings._();
  Settings_ComputerSettings createEmptyInstance() => create();
  static $pb.PbList<Settings_ComputerSettings> createRepeated() => $pb.PbList<Settings_ComputerSettings>();
  @$core.pragma('dart2js:noInline')
  static Settings_ComputerSettings getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Settings_ComputerSettings>(create);
  static Settings_ComputerSettings _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get timeLimitMs => $_getIZ(0);
  @$pb.TagNumber(1)
  set timeLimitMs($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimeLimitMs() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimeLimitMs() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get skillLevel => $_getIZ(1);
  @$pb.TagNumber(2)
  set skillLevel($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSkillLevel() => $_has(1);
  @$pb.TagNumber(2)
  void clearSkillLevel() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get limitStrength => $_getBF(2);
  @$pb.TagNumber(3)
  set limitStrength($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLimitStrength() => $_has(2);
  @$pb.TagNumber(3)
  void clearLimitStrength() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get elo => $_getIZ(3);
  @$pb.TagNumber(4)
  set elo($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasElo() => $_has(3);
  @$pb.TagNumber(4)
  void clearElo() => clearField(4);
}

class Settings extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Settings', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..aOB(1, 'sound')
    ..e<Settings_Language>(2, 'language', $pb.PbFieldType.OE, defaultOrMaker: Settings_Language.ENGLISH, valueOf: Settings_Language.valueOf, enumValues: Settings_Language.values)
    ..aOB(3, 'voiceRecognition')
    ..aOB(4, 'autoMove')
    ..aOB(5, 'randomBw')
    ..e<Settings_PlayerType>(6, 'player1', $pb.PbFieldType.OE, defaultOrMaker: Settings_PlayerType.HUMAN, valueOf: Settings_PlayerType.valueOf, enumValues: Settings_PlayerType.values)
    ..e<Settings_PlayerType>(7, 'player2', $pb.PbFieldType.OE, defaultOrMaker: Settings_PlayerType.HUMAN, valueOf: Settings_PlayerType.valueOf, enumValues: Settings_PlayerType.values)
    ..aOM<Settings_ComputerSettings>(8, 'computerSettings', protoName: 'computerSettings', subBuilder: Settings_ComputerSettings.create)
    ..hasRequiredFields = false
  ;

  Settings._() : super();
  factory Settings() => create();
  factory Settings.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Settings.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Settings clone() => Settings()..mergeFromMessage(this);
  Settings copyWith(void Function(Settings) updates) => super.copyWith((message) => updates(message as Settings));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Settings create() => Settings._();
  Settings createEmptyInstance() => create();
  static $pb.PbList<Settings> createRepeated() => $pb.PbList<Settings>();
  @$core.pragma('dart2js:noInline')
  static Settings getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Settings>(create);
  static Settings _defaultInstance;

  @$pb.TagNumber(1)
  $core.bool get sound => $_getBF(0);
  @$pb.TagNumber(1)
  set sound($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasSound() => $_has(0);
  @$pb.TagNumber(1)
  void clearSound() => clearField(1);

  @$pb.TagNumber(2)
  Settings_Language get language => $_getN(1);
  @$pb.TagNumber(2)
  set language(Settings_Language v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasLanguage() => $_has(1);
  @$pb.TagNumber(2)
  void clearLanguage() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get voiceRecognition => $_getBF(2);
  @$pb.TagNumber(3)
  set voiceRecognition($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasVoiceRecognition() => $_has(2);
  @$pb.TagNumber(3)
  void clearVoiceRecognition() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get autoMove => $_getBF(3);
  @$pb.TagNumber(4)
  set autoMove($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAutoMove() => $_has(3);
  @$pb.TagNumber(4)
  void clearAutoMove() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get randomBw => $_getBF(4);
  @$pb.TagNumber(5)
  set randomBw($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasRandomBw() => $_has(4);
  @$pb.TagNumber(5)
  void clearRandomBw() => clearField(5);

  @$pb.TagNumber(6)
  Settings_PlayerType get player1 => $_getN(5);
  @$pb.TagNumber(6)
  set player1(Settings_PlayerType v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasPlayer1() => $_has(5);
  @$pb.TagNumber(6)
  void clearPlayer1() => clearField(6);

  @$pb.TagNumber(7)
  Settings_PlayerType get player2 => $_getN(6);
  @$pb.TagNumber(7)
  set player2(Settings_PlayerType v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasPlayer2() => $_has(6);
  @$pb.TagNumber(7)
  void clearPlayer2() => clearField(7);

  @$pb.TagNumber(8)
  Settings_ComputerSettings get computerSettings => $_getN(7);
  @$pb.TagNumber(8)
  set computerSettings(Settings_ComputerSettings v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasComputerSettings() => $_has(7);
  @$pb.TagNumber(8)
  void clearComputerSettings() => clearField(8);
  @$pb.TagNumber(8)
  Settings_ComputerSettings ensureComputerSettings() => $_ensure(7);
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

class Response_ChessBoard extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response.ChessBoard', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..aOS(1, 'fen')
    ..aOB(2, 'rotate')
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

  @$pb.TagNumber(2)
  $core.bool get rotate => $_getBF(1);
  @$pb.TagNumber(2)
  set rotate($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRotate() => $_has(1);
  @$pb.TagNumber(2)
  void clearRotate() => clearField(2);
}

class Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..e<Response_Type>(1, 'type', $pb.PbFieldType.OE, defaultOrMaker: Response_Type.NOOP, valueOf: Response_Type.valueOf, enumValues: Response_Type.values)
    ..pc<Response_WifiNetwork>(2, 'networks', $pb.PbFieldType.PM, subBuilder: Response_WifiNetwork.create)
    ..aOM<Settings>(3, 'settings', subBuilder: Settings.create)
    ..aOB(4, 'gameInProgress', protoName: 'gameInProgress')
    ..pPS(5, 'moves')
    ..aOB(6, 'whiteTurn', protoName: 'whiteTurn')
    ..aOS(7, 'state')
    ..aOM<Response_ChessBoard>(8, 'chessBoard', subBuilder: Response_ChessBoard.create)
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
  Response_Type get type => $_getN(0);
  @$pb.TagNumber(1)
  set type(Response_Type v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasType() => $_has(0);
  @$pb.TagNumber(1)
  void clearType() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<Response_WifiNetwork> get networks => $_getList(1);

  @$pb.TagNumber(3)
  Settings get settings => $_getN(2);
  @$pb.TagNumber(3)
  set settings(Settings v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasSettings() => $_has(2);
  @$pb.TagNumber(3)
  void clearSettings() => clearField(3);
  @$pb.TagNumber(3)
  Settings ensureSettings() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.bool get gameInProgress => $_getBF(3);
  @$pb.TagNumber(4)
  set gameInProgress($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasGameInProgress() => $_has(3);
  @$pb.TagNumber(4)
  void clearGameInProgress() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.String> get moves => $_getList(4);

  @$pb.TagNumber(6)
  $core.bool get whiteTurn => $_getBF(5);
  @$pb.TagNumber(6)
  set whiteTurn($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasWhiteTurn() => $_has(5);
  @$pb.TagNumber(6)
  void clearWhiteTurn() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get state => $_getSZ(6);
  @$pb.TagNumber(7)
  set state($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasState() => $_has(6);
  @$pb.TagNumber(7)
  void clearState() => clearField(7);

  @$pb.TagNumber(8)
  Response_ChessBoard get chessBoard => $_getN(7);
  @$pb.TagNumber(8)
  set chessBoard(Response_ChessBoard v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasChessBoard() => $_has(7);
  @$pb.TagNumber(8)
  void clearChessBoard() => clearField(8);
  @$pb.TagNumber(8)
  Response_ChessBoard ensureChessBoard() => $_ensure(7);
}

class Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Request', package: const $pb.PackageName('bluetoothpb'), createEmptyInstance: create)
    ..e<Request_Type>(1, 'type', $pb.PbFieldType.OE, defaultOrMaker: Request_Type.NOOP, valueOf: Request_Type.valueOf, enumValues: Request_Type.values)
    ..aOS(2, 'wifiSsid')
    ..aOS(3, 'wifiPsk')
    ..aOM<Settings>(4, 'settings', subBuilder: Settings.create)
    ..aOS(5, 'move')
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

  @$pb.TagNumber(4)
  Settings get settings => $_getN(3);
  @$pb.TagNumber(4)
  set settings(Settings v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasSettings() => $_has(3);
  @$pb.TagNumber(4)
  void clearSettings() => clearField(4);
  @$pb.TagNumber(4)
  Settings ensureSettings() => $_ensure(3);

  @$pb.TagNumber(5)
  $core.String get move => $_getSZ(4);
  @$pb.TagNumber(5)
  set move($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasMove() => $_has(4);
  @$pb.TagNumber(5)
  void clearMove() => clearField(5);
}

