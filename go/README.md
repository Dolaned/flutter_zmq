# flutter_zmq

This Go package implements the host-side of the Flutter [flutter_zmq](https://github.com/dolaned/flutter_zmq) plugin.

## Usage

Import as:

```go
import flutter_zmq "github.com/dolaned/flutter_zmq/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&flutter_zmq.FlutterZmqPlugin{}),
```
