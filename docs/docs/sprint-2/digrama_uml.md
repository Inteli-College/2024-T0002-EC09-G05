---
title: Diagramas UML
sidebar_position: 5
---


![Diagrama UML de implantação da solução](/img/uml_diagrama.png)

**Diagrama UML de implantação da solução**

Neste diagrama:

- **Main Application**: Este pacote representa a aplicação principal, que consiste em dois componentes principais: main() e Env Loader.
- **Client**: Este componente representa a lógica do cliente MQTT. Ele interage com o MQTT Client para se conectar ao servidor MQTT e publicar mensagens.
- **Env Loader:** Este componente é responsável por carregar as variáveis de ambiente do arquivo .env.
- **MQTT Client**: Este componente é fornecido pela biblioteca Paho MQTT. Ele é usado para se conectar ao broker MQTT e publicar mensagens.
- **Broker**: Este componente representa o servidor MQTT ao qual o cliente se conecta para publicar mensagens.
- **External** Libraries: Este pacote representa as bibliotecas externas usadas pela aplicação. No caso, inclui a biblioteca Paho MQTT e o .env Loader.


![Diagrama UML da comunicação MQTT](/img/rabbit_diagram.png)

**Diagrama de comunicação MQTT**

Neste diagrama:

- Há dois nós principais: o RabbitMQ e o InfluxDB.
- O RabbitMQ é responsável por receber e distribuir mensagens usando o protocolo AMQP. O protocolo AMQP (Advanced Message Queuing Protocol) é um protocolo de mensagens de rede aberto que fornece uma estrutura padronizada para comunicação entre aplicativos cliente e servidores de mensagens. Ele é acessado pelo serviço de publicação de sensor, que envia dados para ele.
- O InfluxDB é um banco de dados de séries temporais usado para armazenar os dados do sensor. Ele é acessado pelo serviço de processamento de dados, que recebe dados do RabbitMQ e os escreve no InfluxDB.
- Os dois serviços principais são o Sensor Publisher e o Data Processor. 

![Digrama geral](/img/general_diagram.png)

**Diagrama geral da aplicação**

Neste diagrama está reprensentado por meio de uma visão macro os principais componentes do sistema, que são constituido pelos sensores, servidor, os testes e o dashboard. Além disso, estão representados os métodos de cada classe bem como quais são as funções de cada método.
