---
title: Setup do Postgres
sidebar_position: 3
---


Essa implementação utiliza o GORM, uma biblioteca de mapeamento objeto-relacional para Go, para facilitar a interação com um banco de dados PostgreSQL. Cinco estruturas de dados são definidas: User, Sensor, Props, Elements, e Layout, cada uma representando entidades diferentes em um sistema.

A estrutura User representa os usuários do sistema, com campos como nome, e-mail e senha, e associações com outras entidades como Directorate e Layout. Directorate representa as diretorias no sistema, enquanto Layout representa os layouts de interface. A estrutura Elements representa os elementos que compõem um layout, e Props armazena propriedades genéricas associadas a esses elementos.



![shcema1](/img/schema1.png)

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária autoincrementada única do usuário
3) Name representa o nome do usuário
4) Email é o endereço de e-mail do usuário
5) Password é a senha do usuário
6) Role é o papel do usuário, com um valor padrão de 1
7) DirectorateRefer é uma referência à diretoria do usuário
8) Directorate é a diretoria à qual o usuário pertence
9) LayoutRefer é uma referência ao layout do usuário
10) Layout é o layout associado ao usuário


![schema2](/img/schema2.png)

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária autoincrementada única 
3) Directorate representa o nome da diretoria


![schema3](/img/schema3.png)

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária autoincrementada única do elemento
3) Name é o nome do elemento
4) PropsRefer é uma referência aos atributos (props) do elemento
5) Props é a estrutura de atributos associada ao elemento



![schema4](/img/schema4.png)
1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária autoincrementada única do atributo
3) Value é o valor do atributo

![schema5](/img/schema5.png)
1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária do layout
3) ElementsRefer é uma referência aos elementos do layout
4) Elements é a estrutura de elementos associada ao layout
5) Index é a posição do layout


![schema6](/img/schema6.png)

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) Identifier é o identificador do sensor
3) Name é o nome do sensor
3) CoordX é a coordenada X do sensor
4) CoordY é a coordenada Y do sensor





