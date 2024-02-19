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


if __name__ =="__main__":
    test_pub_on()
#def test_pub_get_data()