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


    
    def on(self, broker, sec):
        # Configuração do cliente
        client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2, "python_publisher")

        # Conecte ao broker
        client.connect(broker['link'], broker['port'], 30)

        # COntagem de segundos para sair do loop do pytest
        count = 0
        # Loop para publicar mensagens continuamente
        try:
            rand_number = random.uniform(self._min, self._max)
            while count <= sec:
                message = self.get_data(init=rand_number)
                client.publish("test/topic", str(message))
                print(f"{self.nameType}: {str(message)} ")
                count += 1
                time.sleep(2)
        except KeyboardInterrupt:
            print("Publicação encerrada")
        client.disconnect()

    
    def get_data(self, init): 
        
        self.delta += 1/8
        if self.delta > 2:
            self.delta = 0
            rand = (random.uniform(self._min, self._max))
            init = rand*abs((rand/self._max)+(rand/self._min))
              
        variance = round(abs(np.random.normal(0,5))*math.cos(self.delta),1)
        value = (init+variance)
        if value > self._max:
            value = self._max
        if value < self._min:
            value = self._min
            
        return print(f'Valores do Sensor Simulado: {value:,.2f}')
        
        
    
