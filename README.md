# ttqm

simple library for handling mqtt message

### Getting the library

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/h4lim/ttqm"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `qr` package:

```sh
$ go get -u github.com/h4lim/ttqm
```

### Use ttqm

```go
package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/h4lim/ttqm"
)

type CustomMqttManager struct {
	ttqm.MqttManagerContext
}

func (m CustomMqttManager) ExecuteMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Custom Handling: Topic: %s, Message: %s\n", msg.Topic(), msg.Payload())
}

func main() {
	// Create your custom MQTT manager
	customManager := CustomMqttManager{}

	// Initialize MqttContext with your custom manager
	mqttContext := ttqm.MqttContext{
		Url:         "localhost",
		Port:        "1883",
		ClientId:    "exampleClient",
		MqttManager: customManager, // Inject your custom manager here
	}

	// Use the context to send a message and subscribe
	if err := mqttContext.SendToMqtt("example/hierarchy", "Hello World!"); err != nil {
		fmt.Println("Error sending to MQTT:", *err)
	}
}
```

And use the Go command to run the demo:

```
# run main.go
$ go run main.go
```

And use method ExecuteMessage to handling your business logic