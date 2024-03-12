from mqtt.MQTT_Publisher import Sensor

pub_broker = {
    "link": "broker.hivemq.com",
    "port": 1883
}

def test_pub_on():
    sensor = Sensor(
        name="tempTest",
        min=-20,
        max=60,
    )
    assert (sensor.on(pub_broker, 5), 'Erro no pub')



test_pub_on()
