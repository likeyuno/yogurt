/*
 * Copyright (C) 2019-2020 KallyDev
 * This program under GNU General Public License version 3.0, you
 * can redistribute it or modify it under the terms of the, see
 * the link below for more details
 *
 * https://github.com/kallydev/yogurt/blob/master/LICENSE
 */

import 'dart:io';

var indexPath = 'service/download/public';
var htmlPath = 'service/download/template/index.html';
var publicPath = 'config/cert/public.pem';
var privatePath = 'config/cert/private.pem';

const fileType = 'application/octet-stream';
const htmlType = 'text/html';

Future main(List<String> args) async {
  if (args.length < 1) {
    print('please input port');
    return;
  }
  var port = int.parse(args[0]);
  var serverContext = SecurityContext();
  serverContext.useCertificateChain(publicPath);
  serverContext.usePrivateKey(privatePath);
  var httpServer = await HttpServer.bindSecure(
    InternetAddress.anyIPv4,
    port,
    serverContext,
  );

  await for (HttpRequest httpRequest in httpServer) {
    var requestPath = httpRequest.requestedUri.path;
    var filePath = indexPath + requestPath;
    var directory = Directory(filePath);
    if (await directory.exists()) {
      var directoryMap = Map<String, String>();
      var fileMap = Map<String, String>();
      for (var element in await directory.list().toList()) {
        var paths = element.path.split('/');
        print(paths);
        var name = paths[paths.length - 1];
        var path = element.path.split(indexPath)[1];
        if (await Directory(element.path).exists()) {
          directoryMap[name] = path;
        } else {
          fileMap[name] = path;
        }
      }
      var list = '';
      directoryMap.forEach((k, v) {
        list += '<li><a href="$v">$k</a></li>';
      });
      fileMap.forEach((k, v) {
        list += '<li><a href="$v">$k</a></li>';
      });
      var requestPaths = requestPath.split('/');
      var upName = requestPaths[requestPaths.length - 2];
      var upPath = requestPath.split(upName)[0];
      if (requestPath != "/" && upName == '') {
        upName = '/';
      }
      if (upName != '') {
        if (!upName.startsWith('/')) {
          upName = '/$upName';
        }
        upName = '..$upName';
      }
      var up = '<li><a href="$upPath">$upName</a></li>';
      var title = requestPaths[requestPaths.length - 1];
      var htmlDate = await File(htmlPath).readAsString();
      htmlDate = htmlDate.replaceAll('{ Title }', (title == '') ? '/' : title);
      htmlDate = htmlDate.replaceAll("{ List }", list);
      htmlDate = htmlDate.replaceAll("{ Up }", up);
      httpRequest.response.headers.add('content-type', htmlType);
      httpRequest.response.write(htmlDate);
      httpRequest.response.close();
      continue;
    }

    var file = File(filePath);
    if (await file.exists()) {
      httpRequest.response.headers.add('content-type', fileType);
      file.openRead().pipe(httpRequest.response).catchError((e) {
        print(e);
      });
      continue;
    }

    httpRequest.response.statusCode = HttpStatus.notFound;
    httpRequest.response.close();
  }
}
