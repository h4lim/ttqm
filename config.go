package ttqm

import (
	"errors"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqttConfigContext struct {
	mqttContext        MqttContext
	mqttManagerContext IMqttManager
}

type IMqttConfig interface {
	Connect() (mqtt.Client, *error)
	Subscribe(client mqtt.Client, topic string, qos byte) *error
	Publish(client mqtt.Client, topic string, payload string)
}

func NewMqttConfig(mqttContext MqttContext, mqttManager IMqttManager) IMqttConfig {
	return mqttConfigContext{
		mqttContext:        mqttContext,
		mqttManagerContext: mqttManager,
	}
}

func (m mqttConfigContext) Connect() (mqtt.Client, *error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + m.mqttContext.Url + ":" + m.mqttContext.Port)
	opts.SetClientID(m.mqttContext.ClientId)

	if m.mqttContext.Username != nil && m.mqttContext.Password != nil {
		opts.SetUsername(*m.mqttContext.Username)
		opts.SetPassword(*m.mqttContext.Password)
	}
	opts.SetCleanSession(true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		newError := errors.New("cannot connect mosquitto mqtt: " + token.Error().Error())
		return nil, &newError
	}

	return client, nil
}

func (m mqttConfigContext) Subscribe(client mqtt.Client, topic string, qos byte) *error {
	if token := client.Subscribe(topic, qos, m.mqttManagerContext.ExecuteMessage); token.Wait() && token.Error() != nil {
		newError := errors.New("error subscribing to topic: " + token.Error().Error())
		return &newError
	}
	return nil
}

func (m mqttConfigContext) Publish(client mqtt.Client, topic string, payload string) {
	token := client.Publish(topic, 2, false, payload)
	token.Wait()
}
