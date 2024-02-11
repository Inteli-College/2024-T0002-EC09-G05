---
title: Matriz de Riscos
sidebar_position: 2
---
# Matriz de Riscos
Para a identificação dos riscos do projeto, foi utilizada a matriz de riscos, neste aparato são elencados os principais ofensores na execução do projeto, bem como suas respectivas probabilidades e impactos. A partir do pontos elencados abaixo, foram traçadas estratégias de mitigação na ocorrência de tais fatores que podem ser vistos após a matriz.

![img alt](/img/risk_matrix_v1.jpg)

# Estratégias de Mitigação
## Riscos 
### Alta Latência:
A alta latência é um dos principais riscos do projeto, pois pode impactar diretamente na experiência do usuário. Para mitigar este risco, pode-se utilizar o kafka para a comunicação entre os microsserviços, pois o mesmo é uma plataforma de streaming distribuído que pode ser utilizado para a ingestão de dados em tempo real, processamento de dados e análise de dados.

### Perda de Dados por Falha no Equipamento:
Para mitigar a perda de dados por falha no equipamento, pode-se implementar um sistema de avalição da integridade física do sensor, de forma que, ao identificar uma possível desconexão com a rede ou fornecimento de energia, por exemplo, haja um alerta para a manutenção do equipamento.

### Dashboard Ineficiente:
Em decorrência da quantidade de stakeholders do projeto, cabe elencar os requisitos necessários pelos principais demandantes da solução e, a partir disso, desenvolver um dashboard que atenda as necessidades de cada um, através da criação de filtros, visualizações e relatórios personalizados.

### Furto do Equipamento:
Dada a natureza do projeto, a implementação de uma solução para cidades inteligentes, é possível que o equipamento seja alvo de furto. Este fator impacta significativamente nos custos de operação, portanto, pode-se implementar um sistema de rastreamento do equipamento, GPS, de forma que, ao identificar uma possível movimentação do sensor, a equipe responsável seja notificada.

### Falta de Protocolos de Prevenção à Falhas Sistêmicas:
Em se tratando da volumetria de informações que acarreta o projeto de monitoramento de dados do meio ambiente, é imprescindível a implementação de protocolos de prevenção à falhas sistêmicas. 

Durante a concepção do produto, pode-se desenvolver testes de carga para os microsserviços de modo que seja possível identificar gargalos e pontos de falha, bem como a implementação de um sistema de monitoramento de logs, que possa identificar possíveis falhas e alertar a equipe responsável pela implementação da solução. 

Além disto, utilizando sistemas robustos como o Kafka, é possível garantir a integridade dos dados, uma vez que o mesmo é um sistema distribuído e tolerante a falhas, mesmo em casos de falhas de hardware.