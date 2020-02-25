import 'dart:async';
import 'dart:math';
import 'dart:typed_data';
import 'package:flutter/cupertino.dart';
import 'package:flutter_test/flutter_test.dart';

import 'length_decoder.dart';

void main() {
  const size = 1000;
  for (var i = 0; i < size; i++) {
    testWidgets('Should be successfully processed for length ${i}', (tester) async {
      var data = List.generate(i, (idx) => idx%256);
      var input = Stream.fromIterable([
        Uint8List.fromList([0, 0, 0, 0, 0, (i/(256*256)).floor(), (i/256).floor()%256, i%256]),
        Uint8List.fromList(data),
      ]);
      final result = await input.transform(
          StreamTransformer.fromBind(lengthDecoder)).toList();
      expect(result, [data]);
    });
  }

  testWidgets('Should be successfully processed for multiple messages', (tester) async {
    var packets = List<Uint8List>(0);
    var expected = [];
    for (var i = 0; i < size; i++) {
      var data = List.generate(i, (idx) => idx%256);
      packets += [
        Uint8List.fromList([0, 0, 0, 0, 0, (i/(256*256)).floor(), (i/256).floor()%256, i%256]),
        Uint8List.fromList(data),
      ];
      expected += [data];
    }
    var input = Stream.fromIterable(packets);
    final result = await input.transform(
        StreamTransformer.fromBind(lengthDecoder)).toList();
    expect(result, expected);
  });

  testWidgets('Should be successfully processed for small random sized messages', (tester) async {
    var packets = List<Uint8List>(0);
    var random = Random();
    var allData = List<int>(0);
    var expected = [];
    for (var i = 0; i < size/2; i++) {
      var data = List.generate(i, (idx) => idx%256);
      allData += [0, 0, 0, 0, 0, (i/(256*256)).floor(), (i/256).floor()%256, i%256];
      allData += data;
      expected += [data];
    }
    var start = 0;
    while (start < allData.length) {
      var length = random.nextInt(min(10, allData.length - start + 1));
      packets += [Uint8List.fromList(allData.sublist(start, start+length))];
      start += length;
    }
    var input = Stream.fromIterable(packets);
    final result = await input.transform(
        StreamTransformer.fromBind(lengthDecoder)).toList();
    expect(result, expected);
  });

  testWidgets('Should be successfully processed for larger random sized messages', (tester) async {
    var packets = List<Uint8List>(0);
    var random = Random();
    var allData = List<int>(0);
    var expected = [];
    for (var i = 0; i < size; i++) {
      var data = List.generate(i, (idx) => idx%256);
      allData += [0, 0, 0, 0, 0, (i/(256*256)).floor(), (i/256).floor()%256, i%256];
      allData += data;
      expected += [data];
    }
    var start = 0;
    while (start < allData.length) {
      var length = random.nextInt(min(200, allData.length - start + 1));
      packets += [Uint8List.fromList(allData.sublist(start, start+length))];
      start += length;
    }
    var input = Stream.fromIterable(packets);
    final result = await input.transform(
        StreamTransformer.fromBind(lengthDecoder)).toList();
    expect(result, expected);
  });
}