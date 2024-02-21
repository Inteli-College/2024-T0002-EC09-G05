---
title: Análise Financeira
sidebar_position: 1
---

# Análise Financeira

Mediante aos estudos de orçamento inicial do projeto, pensamos em todos os profissionais que farião sentido nesse desenvolvimento. Todos os valores foram baseados na média salarial de cada profissional utilizando o site Glassdoor. O calculo para contratação CLT foi feito no site iDinheiro, já considerando benefícios de vale transporte, vale alimentação e plano de saúde.

Configuração do calculo:

* Vale Alimentação/Refeição: R$1000

* Vale Trasporte: R$150

* Plano de Saúde: R$ 200


| Pessoas | Área do Conhecimento | Salário Mensal | Custo CLT | Meses | Custo Total Anual |
|---|---|---|---|---|---|
| 3 | Desenvolvedor Full-Stack | $5.500,00 | $9.172,22 | 12 | $330.199,92 |
| 1 | Analista de Dados | $5.500,00 | $9.172,22 | 12 | $110.066,64 |
| 1 | Design UX/UI | $5.000,00 | $8.461,11 | 12 | $101.533,32 |
| 1 | Devops | $6.000,00 | $9.883,33 | 12 | $118.599,96 |
| 1 | CyberSegurança | $7.000,00 | $11.305,56 | 12 | $135.666,72 |
| 1 | Engenheiro de Dados | $10.000,00 | $15.572,22 | 12 | $186.866,64 |
| **Total** | **-** | **-** | **-** | **-** |**$982.933,20** |

<br></br>
<br></br>

Pensando nas máquinas para cada uma das 8 pessoas, foi escolhido um modelo que atende as necessidades de cada.

# Tabela de Custos do Notebook Inspiron 15

| **Item** | **Quantidade** | **Valor Unitário** | **Valor Total** | **Total** |
|---|---|---|---| --- |
| Notebook Inspiron 15 - 12a geração Intel® Core™ i5-1235U Windows 11 16GB DDR4 (2x8GB) SSD de 512GB PCIe NVMe M.2 | 8 | $3.598,00 | $28.784,00 | $28.784,00 |

**Observações:**

* O valor total é calculado multiplicando a quantidade pelo valor unitário.
* A tabela não inclui impostos ou outras taxas.

Avaliando os valores das duas tabelas, podemos conlcuir que no período de 1 ano, o custo do projeto ficaria aproximadamente em: 

### **Investimento Inicial:** R$1.011.717,20

<br></br>
<br></br>

# Custo para Manter o Projeto em Funcionamento

Pensando além do investimento inicial de 1 ano, o custo para manter o projeto funcionando e atualizado, existe dois custos, o operacional e o tecnológico.

## Custo Tecnológico

A plataforma do projeto está hospedada na AWS (Amazon Web Services), utilizando a solução em nuvem. A configuração proposta para sustentar a aplicação é a seguinte:

- Amazon EC2 (Elastic Compute Cloud):
Tipo de instância (t2.micro)
Sistema operacional (Linux)
Tempo de execução (ex: 24 horas por dia, 7 dias por semana)
1 Gb de Memoria

- Elastic IP

De acordo com a calculadora da AWS, a máquina com essa configuração teria um custo mensal aproximado de USD 4,23. Multiplicando este valor por 12 meses e convertendo para reais (considerando a taxa de câmbio de R$ 4,97 por dólar), obtemos um custo anual de **R$ 252,34**

## Custo Operacional

A equipe mínima necessária para manter a aplicação em funcionamento e atualizada consiste em 3 profissionais:

* Analista de Dados: Responsável por analisar os novos dados coletados e adaptá-los à plataforma.
* Desenvolvedor Full-Stack: Responsável por realizar atualizações e correções na plataforma.
* Especialista em Cyber Segurança: Responsável por garantir a segurança e estabilidade das informações.

Consultando a tabela acima, vendo os valores mensais (CLT) de cada funcionário ,podemos chegar a conclusão que o custo operacional em 1 ano ficaria, em um valor aproximado de **RS355.800,00**

## Custo Total de Funcionamento

Somando os custos tecnológicos e operacionais, obtemos o custo anual total para manter o projeto em funcionamento:

R$ 252,34 (Tecnológico) + R$ 355.800,00 (Operacional) = **R$ 356.052,34**


## Referências

Site Dell - Notebook Inspiron 15
Link: https://www.dell.com/pt-br/shop/notebooks/notebook-inspiron-15/spd/inspiron-15-3520-laptop/i3520wadl1012w
Data de acesso: 02/2024

Site iDinheiro - Custo de funcionário para Empresa
link: https://www.idinheiro.com.br/calculadoras/calculadora-custo-de-funcionario-para-empresa/
Data de acesso: 02/2024

Site Glassdoor - Consulta de Salários
Link: https://www.glassdoor.com.br/Sal%C3%A1rios/index.htm
Data de acesso: 02/2024

Site Calculator AWS - Calculadora
Link: https://calculator.aws/#/estimate
Data de acesso: 02/2024
