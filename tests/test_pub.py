from src import main
from src.MQTT_Publisher import Sensor

pub_broker = {
    "link": "broker.hivemq.com",
    "port": 1883
}

def test_pub_on():
    sensor = Sensor(
        name="tempTest",
        min=-20,
        max=60,
        broker=pub_broker
    )
    assert sensor.on()

#def test_pub_get_data()