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
cd src\mqtt
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

Com as dependências instaladas basta rodar o projeto com o comando abaixo:
```
python main.py
```

 ## Funcionamento do módulo 
Caso todas as etapas acima estejam satisfeitas, o resultado esperado encontra-se neste vídeo de demonstração: [Click Para Ver](https://drive.google.com/drive/folders/1mXoaW-fK-zhebGrNRBtqksd4BWKMuF36?usp=sharing)

