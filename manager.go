package ttqm

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttManagerContext struct {
}

type IMqttManager interface {
	ExecuteMessage(client mqtt.Client, msg mqtt.Message)
}

func NewMqttManager() IMqttManager {
	return MqttManagerContext{}
}

func (m MqttManagerContext) ExecuteMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
}
