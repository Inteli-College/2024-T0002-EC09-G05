# Requisitos

Requisitos funcionais e não funcionais são as descrições do funcionamento de um sistema. Tal documentação é crucial para o sucesso do projeto de Smart City. Definidos com clareza e precisão, garantem que nosso sistema seja um sistema útil, eficiente, seguro e sustentável, capaz de melhorar a vida da população.

## Contexto

A definição dos requisitos do projeto foi um processo colaborativo. As entrevistas com clientes e stakeholders foram muito elucidativas e permitiu o desenvolvimento inicial da nossa arquitetura e funcionalidade do sistema. Logo após esse momento, em conversas com o grupo e professores, reunimos o conhecimento necessário para o levantamento dos requisitos. 


### Tabela de Requisitos Funcionais

| Código | Requisito Funcional                                                                                                 |
|--------|---------------------------------------------------------------------------------------------------------------------|
| RF 01  | O sistema deve conter um pré processamento de toda entrada de dados, com o objetivo de padronizar antes de serem armazenados ou manuseados por usuários. |
| RF 02  | O Dashboard deve ter personalização pelo proprio usuário para consultas de diferentes objetivos. |
| RF 03  | O Dashboard suportar visualizações de diferentes gráficos, como de linha, de barra, de coluna e de área. |
| RF 04  | O sistema deve conter um sistema de filtragem de dados |
| RF 05  | O sistema deve permitir a exportação dos dados filtrados pelo usuario através de um arquivo CSV. |
| RF 06  | O sistema de conter uma seção pública de acesso ao publico geral, com a disponibilização de informações metereológicas, como intensidade de radição solar, precipitação entre outras . |
| RF 07  | O sistema deve conter uma seção de administração, com acesso restrito a pessoas autorizadas pela Prodam. Nessa seção será possivel gerenciar o sistema, seus usuários e ter acesso a logs de sistema  |
| RF 08  | O sistema deve garantir a retenção de dados por um periodo minimo de 4 anos, sendo possível assim a analise descritiva e preditiva a partir de dados históricos. |

### Tabela de Requisitos Não Funcionais

| Código | Requisito Não Funcional                                                                                                 |
|--------|---------------------------------------------------------------------------------------------------------------------|
| RNF 01  | O Dashboard deve ser intuitivo e acessivel, baseando seu design na nossa persona representante do público geral. |
| RNF 02  | O sistema deve garantir a confiabilidade dos dados armazenado, estando totalmente em conformidade com as lei e regulamentações de proteção de dados. |
| RNF 03  | O sistema deve suportar 2 ou mais (mais de uma) requisições simultâneas, com um tempo maximo de resposta de 12 segundos por requisição. |
| RNF 04  | Após o usuário clicar em exportar dados filtrados, o sistema deve gerar um arquivo csv em até 25s. |
| RNF 05  | O sistema deve cumprir a regulamentação de informações do usuário usando criptgrafia JWT em suas informações. |
| RNF 06  | O sistema deve gerar os gráficos filtrados pelo usuário em no maximo 30 segundo. |
| RNF 07  | O Dashboard deve ser atualizado automaticamente em até 2 minutos, caso o usuario não esteja filtrando por dados históricos. |


### Testes de Requisitos Funcioanais

**RF 01: Pré-processamento de dados**

**Objetivo:** Validar se o sistema está pré-processando os dados de acordo com as regras definidas.
* **Descrição do Teste:** Injetar dados com diferentes formatos e verificar se o sistema os converte para o formato padrão.
* **Resultado Esperado:** Todos os dados devem ser convertidos para o formato padrão.

**RF 02: Personalização do Dashboard**

**Objetivo:** Validar se o usuário pode personalizar o Dashboard de acordo com suas necessidades.
* **Descrição do Teste:** Criar diferentes dashboards com diferentes gráficos e filtros.
* **Resultado Esperado:** O sistema deve permitir a criação e a personalização de dashboards.

**RF 03: Visualização de gráficos**

**Objetivo:** Validar se o sistema está gerando os gráficos corretamente.
* **Descrição do Teste:** Selecionar diferentes tipos de gráficos e verificar se eles são gerados corretamente.
* **Resultado Esperado:** Os gráficos devem ser gerados de acordo com o tipo selecionado.

**RF 04: Filtragem de dados**

**Objetivo:** Validar se o sistema está filtrando os dados de acordo com os critérios definidos pelo usuário.
* **Descrição do Teste:** Filtrar os dados por diferentes critérios e verificar se os resultados são os esperados.
* **Resultado Esperado:** O sistema deve retornar os dados filtrados de acordo com os critérios definidos pelo usuário.

**RF 05: Exportação de dados**

**Objetivo:** Validar se o sistema está exportando os dados filtrados para um arquivo CSV.
* **Descrição do Teste:** Filtrar os dados e exportá-los para um arquivo CSV.
* **Resultado Esperado:** O sistema deve exportar os dados filtrados para um arquivo CSV.

**RF 06: Seção pública**

**Objetivo:** Validar se a seção pública está disponível para o público geral e se as informações estão corretas.
* **Descrição do Teste:** Acessar a seção pública e verificar se as informações estão disponíveis e corretas.
* **Resultado Esperado:** As informações da seção pública devem estar disponíveis e corretas.

**RF 07: Seção de administração**

**Objetivo:** Validar se a seção de administração está disponível para usuários autorizados e se as funcionalidades estão funcionando corretamente.
* **Descrição do Teste:** Acessar a seção de administração e verificar se as funcionalidades estão funcionando corretamente.
* **Resultado Esperado:** As funcionalidades da seção de administração devem estar funcionando corretamente.

**RF 08: Retenção de dados**

**Objetivo:** Validar se o sistema está retendo os dados por um período mínimo de 4 anos.
* **Descrição do Teste:** Verificar se os dados coletados há 4 anos ainda estão disponíveis no sistema.
* **Resultado Esperado:** Os dados coletados há 4 anos devem estar disponíveis no sistema.


### Teste de Requisitos Não Funcionais

**RNF 01: Usabilidade do Dashboard**

**Objetivo:** Validar se o Dashboard é intuitivo e acessível para o público geral.
* **Descrição do Teste:** Aplicar testes de usabilidade com público semelhante a persona, observar os resultados e coletar feedbacks.
* **Resultado Esperado:** A persona deve ser capaz de navegar no Dashboard facilmente e encontrar as informações que precisa.

**RNF 02: Confiabilidade dos dados**

**Objetivo:** Validar se o sistema garante a confiabilidade dos dados armazenados.
* **Descrição do Teste:** Verificar se o sistema está em conformidade com as leis e regulamentações de proteção de dados. Conferir se os dados armazenados no sistema são precisos, completos e consistentes.
* **Resultado Esperado:** Os dados armazenados no sistema devem ser precisos, completos e consistentes. O sistema deve ter mecanismos para garantir a segurança dos dados armazenados.

**RNF 03: Desempenho do sistema**

**Objetivo:** Validar se o sistema suporta 2 ou mais requisições simultâneas com um tempo máximo de resposta de 12 segundos por requisição.
* **Descrição do Teste:** Simular 2 ou mais requisições simultâneas ao sistema e medir o tempo de resposta de cada requisição. Ao final, verificar se o tempo de resposta de cada requisição é inferior a 12 segundos.
* **Resultado Esperado:** O sistema deve suportar 2 ou mais requisições simultâneas com um tempo máximo de resposta de 12 segundos por requisição.

**RNF 04: Tempo de geração de arquivos CSV**

**Objetivo:** Validar se o sistema gera um arquivo CSV em até 25 segundos após o usuário clicar em exportar dados filtrados.
* **Descrição do Teste:** Filtrar dados no sistema e clicar em exportar para CSV. Verificar se o tempo de geração do arquivo CSV é inferior a 25 segundos.
* **Resultado Esperado:** O sistema deve gerar um arquivo CSV em até 25 segundos após o usuário clicar em exportar dados filtrados.

**RNF 05: Segurança da informação do usuário**

**Objetivo:** Validar se o sistema cumpre a regulamentação de informações do usuário usando criptografia JWT em suas informações.
* **Descrição do Teste:** Testar se a criptografia JWT está sendo implementada de acordo com a regulamentação de informações do usuário.
* **Resultado Esperado:** O sistema deve estar usando criptografia JWT para proteger as informações do usuário. A criptografia JWT deve estar sendo implementada de acordo com a regulamentação de informações do usuário.

**RNF 06: Tempo de geração de gráficos**

**Objetivo:** Validar se o sistema gera os gráficos filtrados pelo usuário em no maximo 30 segudos.
* **Descrição do Teste:** Filtrar dados no sistema e selecionar a opção de gerar gráficos. Verificar se o tempo de geração dos gráficos é inferior a 30 segundos.
* **Resultado Esperado:** O sistema deve gerar os gráficos filtrados pelo usuário em no máximo 30 segundos.

**RNF 07: Atualização automática do Dashboard**

**Objetivo:** Validar se o Dashboard é atualizado automaticamente em até 2 minutos, caso o usuário não esteja filtrando por dados históricos.
* **Descrição do Teste:** Acessar o Dashboard e verificar se ele é atualizado automaticamente em até 2 minutos. Observar se a atualização do Dashboard é suave e não interfere na navegação do usuário.
* **Resultado Esperado:** O Dashboard deve ser atualizado automaticamente em até 2 minutos, caso o usuário não esteja filtrando por dados históricos. A atualização do Dashboard deve ser suave e não interferir na navegação do usuário.
