///
//  Generated code. Do not modify.
//  source: bluetoothpb.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Request_Type extends $pb.ProtobufEnum {
  static const Request_Type NOOP = Request_Type._(0, 'NOOP');
  static const Request_Type START_WIFI_SCAN = Request_Type._(1, 'START_WIFI_SCAN');
  static const Request_Type STOP_WIFI_SCAN = Request_Type._(2, 'STOP_WIFI_SCAN');
  static const Request_Type CONFIGURE_WIFI = Request_Type._(3, 'CONFIGURE_WIFI');
  static const Request_Type FORGET_WIFI = Request_Type._(4, 'FORGET_WIFI');
  static const Request_Type CONNECT_WIFI = Request_Type._(5, 'CONNECT_WIFI');

  static const $core.List<Request_Type> values = <Request_Type> [
    NOOP,
    START_WIFI_SCAN,
    STOP_WIFI_SCAN,
    CONFIGURE_WIFI,
    FORGET_WIFI,
    CONNECT_WIFI,
  ];

  static final $core.Map<$core.int, Request_Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Request_Type valueOf($core.int value) => _byValue[value];

  const Request_Type._($core.int v, $core.String n) : super(v, n);
}

