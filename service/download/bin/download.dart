/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

import 'dart:io';

const contentType = 'application/octet-stream';

Future main(List<String> args) async {
  if (args.length != 4) {
    print('please input args');
    return;
  }

  var port = int.parse(args[0]);
  var path = args[1];
  var publicPath = args[2];
  var privatePath = args[3];

  var serverContext = SecurityContext();
  serverContext.useCertificateChain(publicPath);
  serverContext.usePrivateKey(privatePath);

  var httpServer = await HttpServer.bindSecure(
    InternetAddress.anyIPv4,
    port,
    serverContext,
  );

  await for (HttpRequest httpRequest in httpServer) {
    print(path + httpRequest.requestedUri.path);
    var file = File(path + httpRequest.requestedUri.path);
    file.exists().then((found) {
      if (found) {
        httpRequest.response.headers.add('content-type', contentType);
        file.openRead().pipe(httpRequest.response).catchError((e) {
          print(e);
        });
      } else {
        httpRequest.response.statusCode = HttpStatus.notFound;
        httpRequest.response.close();
      }
    });
  }
}
