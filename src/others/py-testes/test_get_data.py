from mqtt.MQTT_Publisher import Sensor
import random

def test_get_data():
    sensor = Sensor(
        name='RadX-300',
        min=200,
        max=1100
    )
    assert (sensor.get_data(random.randrange(start=200, stop=1100)), 'erro no get_data')

test_get_data()