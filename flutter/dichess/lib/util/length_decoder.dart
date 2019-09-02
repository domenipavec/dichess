import 'dart:async';
import 'dart:typed_data';
import 'package:flutter/foundation.dart';

Stream<Uint8List> lengthDecoder(Stream<Uint8List> input) async* {
  var buffer = WriteBuffer();
  int wantedLength;
  int gotLength = 0;

  await for (var r in input) {
    buffer.putUint8List(r);
    gotLength += r.length;

    if (wantedLength == null) {
      if (gotLength < 8) {
        continue;
      }

      var data = buffer.done();
      wantedLength = data.getUint64(0);

      buffer = WriteBuffer();
      buffer.putUint8List(data.buffer.asUint8List(8));
      gotLength -= 8;
    }

    if (gotLength < wantedLength) {
      continue;
    }

    var data = buffer.done();
    yield data.buffer.asUint8List(0, wantedLength);

    buffer = WriteBuffer();
    buffer.putUint8List(data.buffer.asUint8List(wantedLength));
    gotLength -= wantedLength;
    wantedLength = null;
  }
}
