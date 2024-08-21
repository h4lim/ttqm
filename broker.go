package ttqm

type MqttContext struct {
	Broker      string
	Port        string
	ClientId    string
	Username    *string
	Password    *string
	MqttManager IMqttManager
}

type IMqtt interface {
	SendToMqtt(hierarchy string, message string) *error
}

func NewMqtt(model MqttContext) IMqtt {
	return model
}

func (m MqttContext) SendToMqtt(hierarchy string, message string) *error {

	mqttConfig := NewMqttConfig(m, m.MqttManager)
	client, err := mqttConfig.Connect()
	if err != nil {
		return err
	}

	defer client.Disconnect(250)

	if err := mqttConfig.Subscribe(client, "#", 2); err != nil {
		return err
	}

	mqttConfig.Publish(client, hierarchy, message)

	return nil
}
