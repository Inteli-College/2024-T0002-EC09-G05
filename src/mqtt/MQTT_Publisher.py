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

        # Configuração do cliente
        self.client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2, "python_publisher")

        self.connection = False
        self.counter = 10

        self.first_number = False
        self.new_number = 0

    def on(self, broker, test={"tested":False,"sec":0}):
        
        # Conecte ao broker
        self.client.connect(broker['link'], broker['port'], 30)
        self.connection = True
        self.counter = test["sec"]
        # Loop para publicar mensagens continuamente
        try:
            rand_number = random.uniform(self._min, self._max)
            while self.connected(test):
                message = self.get_data(init=rand_number)
                self.client.publish("test/topic", str(message))
                time.sleep(1)
        except KeyboardInterrupt:
            print("Publicação encerrada")
        self.client.disconnect()

    def off(self) -> None:
        self.connection = False
        self.client.disconnect()
        print("Publicação encerrada")
    
    def connected(self, test) -> bool:
        if test["tested"] == True:
            self.counter -= 1
            if self.counter < 0:
                self.off()

        return self.connection
    
    def get_data(self, init): 
        if self.first_number == False:
            self.first_number = True
        else:
            init = self.new_number

        self.delta += 1/2
        if self.delta > 2:
            self.delta = 0
            rand = (random.uniform(self._min, self._max))
            self.new_number = rand*abs((rand/self._max)+(rand/self._min))
        
        
        variance = round(abs(np.random.normal(0,5))*math.cos(self.delta),1)
        value = (init+variance)

        if value > self._max:
            value = self._max
        if value < self._min:
            value = self._min
            
        return print(f'Valores do Sensor Simulado: {value:,.2f}')
        
        
    
