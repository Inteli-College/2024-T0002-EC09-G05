---
title: Modelo - V1
sidebar_position: 1
---

![img alt](/img/architecture_v1.png)

Nesse primeiro modelo de arquitetura, temos:

- **Sensores:** Representam a camada de coleta de dados, onde sensores diversos capturam as informações em tempo real.

- **MQTT:** Protocolo leve de mensagens usado para o envio de dados dos sensores para o sistema de processamento. É ideal para comunicação em redes de largura de banda limitada e equipamentos distríbuidos com baixa potência.

- **Apache Kafka:** Atua como uma camada de mensageria e persistência, recebendo dados via MQTT. O Kafka permite a pré-filtragem dos dados e organiza-os em tópicos, garantindo a distribuição eficiente para os sistemas do servidor.

- **Servidor:** O servidor é desenvolvido em Go, usando o framework Gin. Ele implementa a lógica de negócios e o processamento dos dados, que envolvem a agregação, análise e transformação dos dados recebidos do Kafka.

- **InfluxDB:** Banco de dados otimizado para armazenar e consultar séries temporais, ou seja, dados que são coletados em intervalos regulares. Adequado para lidar com as grandes quantidades de dados gerados pelos sensores. Também alimenta o dashboard.

- **Frontend (dashboard):**  Implementado com Next.js, uma estrutura de React. O dashboard é onde os dados processados são exibidos visualmente para análise e monitoramento das métricas da cidade inteligente.