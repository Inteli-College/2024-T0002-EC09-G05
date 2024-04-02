---
title: Setup do Postgres
sidebar_position: 3
---

Com o objetivo de criar um sistema mais abrangente e adaptável, nesta etapa desenvolvemos a arquitetura de um banco de dados escalável capaz de suportar a criação de múltiplos layouts.

![shcema1](/img/Database_diagram.png)

Essa estrutura se baseia em 5 entidades: User, Directorate, Layout, Elements, e Props, cada uma tem o seu objetivo explicado mais a frente. Mas antes vale resaltar que essa implementação utiliza o GORM, uma biblioteca de mapeamento objeto-relacional para Go, para facilitar a interação com um banco de dados PostgreSQL.


![shcema1](/img/schema1.png)

A entidade User representa um usuário logado em nossa aplicação, tendo os campos email, password, name, role e DirectorateRefer, que é a relação estabelecida com a entidade de Directorate. Assim, um usuário pode estar associado a várias diretórias.

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária autoincrementada única do usuário
3) Name representa o nome do usuário
4) Email é o endereço de e-mail do usuário
5) Password é a senha do usuário
6) Role é o papel do usuário, com um valor padrão de 1
7) DirectorateRefer é uma referência à diretoria do usuário
8) Directorate é a diretoria à qual o usuário pertence


![schema2](/img/schema2.png)

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) ID é a chave primária autoincrementada única 
3) Directorate representa o nome da diretoria

![schema5](/img/schema5.png)

A entidade Layout pode ser associada a n usuários, tendo em seus campos todas as informações necessarias para constuir e organizar n componentes.

1) Inclui os campos padrão do GORM para um modelo de banco de dados.
2) ID é a chave primária do layout.
3) ElementsRefer é uma referência aos elementos do layout.
4) UserRefer é uma referência ao usuário proprietário do layout.
5) Elements é a estrutura de elementos associada ao layout.
6) Index é a posição dos elementos no layout.


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



![schema6](/img/schema6.png)

1) Incorpora os campos padrão do GORM para um modelo de banco de dados
2) Identifier é o identificador do sensor
3) Name é o nome do sensor
3) CoordX é a coordenada X do sensor
4) CoordY é a coordenada Y do sensor





