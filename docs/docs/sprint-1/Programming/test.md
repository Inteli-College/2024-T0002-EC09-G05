---
title: Testes e Resultados
sidebar_position: 2
---

# Testes e Resultados

Nessa seção se encontram os testes produzidos para validar nossa implementação inicial do projeto, garantindo o bom funcionamento do sistema

## Setup inicial do ambiente de testes

Nessa primeira versão do sistema foi utilizado python, e para reproduzir os testes nesse ambiente é necessário algumas configurações

A primeira delas, é recomendavel o uso de virtual environment. Para isso, rode:
```
python3 virtualenv venv
```

Para ativar o venv, use:
```
source venv/bin/activate
```

Agora dentro da pasta **src**, instale o pacote do projeto:
```
pip install -e .
```

E por ultimo instale as dependencias necessárias para o teste:
```
pip install -e ".[test]"
```

### Teste de publicação de mensageria via MQTT

**Objetivo**: Verificar a publicação de mensagens no Broker publico HiveMQ a partir de um sensor simulado

**Estrutura do teste**: 
Primeiro, no arquivo 'test_pub.py' dentro da pasta 'tests', configuramos um cliente MQTT (Paho MQTT, usado no nosso caso), instanciado a partir da classe de um Sensor 'Genérico'. A partir dele publicamos os dados em um Broker, para isso usamos o Broker publico da HiveMQ

Para rodar, dentro da pasta 'tests' execute o seguinte código:
```
pytest test_pub.py
```

**Resultado esperado**
![Teste de publicação MQTT](/img/pytestPub.png)


### Teste de geração e envio de dados simulado

**Objetivo**: Verificar se o código gerador de dados simulaos funciona e se ele envia para a instancia que o chamar, como qualquer instancia do nosso Sensor

**Estrutura do teste**: 
Primeiro, no arquivo 'test_get_data.py' dentro da pasta 'tests', instanciamos um novo objeto da classe 'Sensor', de nome 'Radx-300', com valor minimo de 200 e valor maximo de 1100 (nome e valores inspirados nesse [sensor de radiação solar](https://sigmasensors.com.br/produtos/sensor-de-radiacao-solar-sem-fio-hobonet-rxw-lib-900))
A cada

Para rodar, dentro da pasta 'tests' execute o seguinte código:
```
pytest test_get_data.py
```

**Resultado esperado**

![Teste de geração de dados](/img/pytestGetData.png)

### Rodando todos de uma vez

Também é possivel rodar todos os testes de uma unica vez, para isso, dentro da pasta 'src', execeute o comando:

```
pytest
```

Resultado na linha de comando:

![Teste geral](/img/pytestFull.png)