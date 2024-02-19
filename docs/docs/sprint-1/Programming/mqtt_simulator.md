---
title: Simulador MQTT
sidebar_position: 1
---

# Simulador MQTT

Esta seção da documentação descreve o módulo responsável pela realização dos testes iniciais da infraestrutura da solução. Para este propósito, desenvolvemos o módulo "MQTT_publisher", que possibilita a simulação de um sensor e o envio dos dados correspondentes a ele.

## Setup do ambiente 

Este tutorial pressupõe que você já tenha instalado as seguintes tecnologias:
- [Python](https://www.python.org)
- [Mosquitto](https://mosquitto.org)

Para o uso dessa funcionalidade siga as etapas abaixo:

Entre no diretório '2024-T0002-EC09-G05\src\mqtt':
```
cd 2024-T0002-EC09-G05\src\mqtt
```

É recomendavel o uso de virtual environment. Para isso, rode:
```
python -m venv venv
```

Para ativar o venv em MAC e UBUNTU, use:
```
source venv/bin/activate
```

Para ativar o venv no WINDOWS, use:
```
venv\Scripts\activate
```

Agora, instale as dependências do projeto:
```
pip install -r requirements.txt 
```

 
**Objetivo**: Verificar a publicação de mensagens no Broker publico HiveMQ a partir de um sensor simulado

**Estrutura do teste**: 
Primeiro, no arquivo 'test_pub.py' dentro da pasta 'tests', configuramos um cliente MQTT (Paho MQTT, usado no nosso caso), instanciado a partir da classe de um Sensor 'Genérico'. A partir dele publicamos os dados em um Broker, para isso usamos o Broker publico da HiveMQ

Para rodar, dentro da pasta 'tests' execute o seguinte código:
```
pytest test_pub.py
```

**Resultado esperado**
![img alt](/img/code_main-sensor-test_sprint_01.png)
![img alt](/img/code_sensor-def-get_data_sprint_01.png)
![img alt](/img/code_sensor-def-on_sprint_01.png)
![img alt](/img/code_sensor-metodo-construtor_sprint_01.png)