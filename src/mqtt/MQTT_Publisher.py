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

        self.status = "new wave"
        self.start_wave = True
        self.current_position = 0
        self.go_position = 0
        self.wave_size = 6

    def on(self, broker, test={"tested":False,"sec":0}):
        
        # Conecte ao broker
        self.client.connect(broker['link'], broker['port'], 30)
        self.connection = True
        self.counter = test["sec"]
        # Loop para publicar mensagens continuamente
        try:
            
            while self.connected(test):
                message = self.generate_wave_data()
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
    
    def generate_wave_data(self): 
        value = 0

        if self.status == "new wave":
            self.wave_size = 6
            if self.start_wave == True:
                self.current_position = (random.randrange(self._min, self._max))
                self.start_wave == False
                value = self.current_position
            self.status = "wave"
        elif self.status == "wave":
            self.wave_size -= 1
            value = self.current_position + round(abs(np.random.normal(0,5))*math.cos(self.delta),1)
            if ( self.wave_size <= 0):
                self.status = "transtion"
                self.go_position = (random.randrange(self._min, self._max))
            
        elif self.status == "transtion":
            print(f"Indo de {self.current_position} para {self.go_position}")
            if ( self.current_position == self.go_position):
                 self.current_position = self.go_position
                 self.status = "new wave"
        else: 
            "Erro"


        if value > self._max:
            value = self._max
        if value < self._min:
            value = self._min
            
        return print(f'Valores do Sensor Simulado: {value:,.2f}')
        
        
    






# init = random.uniform(self._min, self._max)


# self.delta += 1/2
# if self.delta > 2:
#     self.delta = 0
#     rand = (random.uniform(self._min, self._max))
#     self.new_number = rand*abs((rand/self._max)+(rand/self._min))


# variance = round(abs(np.random.normal(0,5))*math.cos(self.delta),1)
# value = (init+variance)