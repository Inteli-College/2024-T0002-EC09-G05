from MQTT_Publisher import Sensor

pub_broker = {
    "link":"broker.hivemq.com",
    "port":1883
}

temperature_sensor = Sensor(
    name="temperatura",
    min = -10,
    max = 10
)



if __name__ == "__main__":
    temperature_sensor.on(pub_broker)
    

