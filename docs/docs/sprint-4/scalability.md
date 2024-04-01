---
title: Escalabilidade
sidebar_position: 1
---

## Escalabilidade

Todo projeto inicial possui uma versão de prova de conceito, que se concentra na funcionalidade principal capaz de agregar o maior valor. O objetivo é demonstrar que o desenvolvimento inicial é viável e que é possível desenvolvê-lo.
Quando falamos em escalabilidade, estamos nos referindo à capacidade de tornar o projeto mais robusto, refere-se à capacidade de um sistema, rede ou processo lidar com um aumento crescente na demanda de recursos ou de trabalho de forma eficiente. Em outras palavras, é a capacidade de expandir ou contrair um sistema em resposta ao aumento ou diminuição da carga, mantendo ou melhorando o desempenho, a eficiência e a disponibilidade.
Analisando essa lógica, podemos chegar em duas opções de escalabilidade, a Horizontal e a Vertical.

### Escalabilidade Vertical

Também conhecida como "escala para cima", a escalabilidade vertical refere-se à adição de recursos a um único servidor ou máquina. Isso pode incluir a atualização da CPU, adição de mais memória RAM, aumento da capacidade de armazenamento, entre outros.

**Características da Escalabilidade Vertical:**
- Aumento rápido de desempenho: Ao adicionar mais recursos a uma única máquina, é possível obter um aumento imediato no desempenho.
- Limitações físicas: Existe um limite para o quanto você pode escalar verticalmente, pois as máquinas têm capacidades físicas e técnicas máximas.
- Interrupção potencial: Em alguns casos, pode ser necessário reiniciar o servidor ou interromper o serviço para realizar atualizações.

### Escalabilidade Horizontal

Conhecida como "escala para fora", a escalabilidade horizontal envolve adicionar mais máquinas ao sistema, distribuindo a carga de trabalho entre elas. Isso é frequentemente utilizado em ambientes de computação em nuvem e arquiteturas distribuídas.

**Características da Escalabilidade Horizontal:**
- Flexibilidade: É possível adicionar mais máquinas conforme a demanda cresce, permitindo uma escalabilidade mais elástica.
- Complexidade de gerenciamento: Distribuir a carga de trabalho entre várias máquinas pode exigir uma arquitetura mais complexa e estratégias de balanceamento de carga.
- Menor risco de interrupção: Ao adicionar mais máquinas, é possível realizar atualizações ou manutenções em uma delas sem afetar todo o sistema.

### Resumo:

- **Escalabilidade Vertical:** Aumenta os recursos de uma única máquina ou servidor, adequado para sistemas que exigem maior poder de processamento ou capacidade de armazenamento em uma única instância.
  
- **Escalabilidade Horizontal:** Adiciona mais máquinas ao sistema para distribuir a carga de trabalho, adequado para sistemas distribuídos e ambientes que precisam lidar com um grande número de usuários ou requisições.

Pensando nisso, conforme foi apresentado o Diagrama de Sequência na Sprint 3 em Relatório. O projeto além de conseguir demonstrar sua prova de conceito, sua arquitetura foi construída de tal forma que a sua escalabilidade é de maneira Horizontal, como demonstrado na imagem abaixo:

![sequential_diagram_gif_database_system](/img/sequential_diagram_database_system_v1.gif)

O sensores vão gerar os dados que seram enviados para o RabbitMQ, que fará a distribuição dessa informações para diferentes servidores em rede que vão fazer o processamento dos dados e os mesmos vão gerar uma conexão com o Banco de Dados e enviar os mesmo para o InfluxDB.

## Análise de Escalabilidade

Foi dada a proposta de utilizar o Kafka que é uma plataforma de streaming de eventos distribuída, projetada para lidar com alto volume de dados em tempo real. Ele é utilizado para:

    Coletar, armazenar e processar grandes volumes de dados de forma escalável.
    Facilitar a integração entre sistemas e aplicações em tempo real.
    Prover uma arquitetura de mensageria distribuída para stream de dados, permitindo a publicação e subscrição de eventos.
    Suportar aplicações de análise em tempo real, monitoramento, processamento de dados em fluxo contínuo e integração entre microserviços.

Devido à complexidade na configuração e integração do Kafka para a aplicação específica do projeto, assim como algumas funcionalidades que, embora interessantes como o mantimento dos dados até a sua consulta, são inviáveis por enquanto para proposta, optamos por manter o RabbitMQ. Esta decisão foi influenciada pela facilidade de uso do RabbitMQ e pela arquitetura do projeto, que foi projetada para minimizar o congestionamento de dados no broker. A escalabilidade horizontal da solução, distribui dados para os servidores destinados ao processamento, torna injustificável a implementação de uma das principais funcionalidades oferecidas pelo Kafka.