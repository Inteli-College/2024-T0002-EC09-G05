import paho.mqtt.client as mqtt
import time
import random
import numpy as np
import math


class Sensor():
    def __init__(self, name, min, max) -> None:
        self.nameType = name
        self.delta = 0
        self._min = min
        self._max = max


    
    def on(self, broker):
        # Configuração do cliente
        client = mqtt.Client("python_publisher")

        # Conecte ao broker
        client.connect(broker['link'], broker['port'], 30)

        # Loop para publicar mensagens continuamente
        try:
            self.rand_number = random.uniform(self._min, self._max)
            while True:
                message = self.get_data(init=self.rand_number)
                client.publish("test/topic", str(message))
                print(f"{self.nameType}: {str(message)} | {time.strftime("%H:%M:%S")}")
                time.sleep(2)
        except KeyboardInterrupt:
            print("Publicação encerrada")
        client.disconnect()
    
    def get_data(self, init): 
        
        self.delta += 1/8
        if self.delta > 2:
            self.delta = 0
            rand = (random.uniform(self._min, self._max))
            self.rand_number = rand*abs((rand/self._max)+(rand/self._min))
              
        variance = round(abs(np.random.normal(0,5))*math.cos(self.delta),1)
        value = (init+variance)
        if value > self._max:
            value = self._max
        if value < self._min:
            value = self._min
            
        return f'{value:,.2f}'
        
        
    
