///
//  Generated code. Do not modify.
//  source: bluetoothpb.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, pragma, String;

import 'package:protobuf/protobuf.dart' as $pb;

class Response_ChessBoard extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response.ChessBoard', package: const $pb.PackageName('bluetoothpb'))
    ..a<$core.List<$core.int>>(1, 'image', $pb.PbFieldType.OY)
    ..aOS(2, 'pgn')
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

  $core.List<$core.int> get image => $_getN(0);
  set image($core.List<$core.int> v) { $_setBytes(0, v); }
  $core.bool hasImage() => $_has(0);
  void clearImage() => clearField(1);

  $core.String get pgn => $_getS(1, '');
  set pgn($core.String v) { $_setString(1, v); }
  $core.bool hasPgn() => $_has(1);
  void clearPgn() => clearField(2);
}

class Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Response', package: const $pb.PackageName('bluetoothpb'))
    ..a<Response_ChessBoard>(1, 'chessBoard', $pb.PbFieldType.OM, Response_ChessBoard.getDefault, Response_ChessBoard.create)
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
}

