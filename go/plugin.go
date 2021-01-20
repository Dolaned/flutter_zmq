package flutter_zmq

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	zmq "github.com/pebbe/zmq4"
)

const channelName = "flutter_zmq"

// FlutterZmqPlugin implements flutter.Plugin and handles method.
type FlutterZmqPlugin struct{}

var _ flutter.Plugin = &FlutterZmqPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *FlutterZmqPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	return nil
}

func (p *FlutterZmqPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}
