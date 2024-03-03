---
title: Análise de Impacto Ético
sidebar_position: 1
---

# Análise de Impacto Ético

## Privacidade e proteção de dados (Uso e armazenamento)

A LGPD, ou Lei Geral de Proteção de Dados, é uma legislação brasileira que visa regulamentar o tratamento de dados pessoais no país. Tem como objetivo proteger os direitos fundamentais de liberdade e privacidade, bem como o livre desenvolvimento da personalidade da pessoa natural. Ela estabelece regras claras sobre a coleta, armazenamento, processamento e compartilhamento de dados pessoais, visando garantir a privacidade e a segurança dos dados dos cidadãos brasileiros.

Pensando nessa lógica, a solução desenvolvida visa coletar, armazenar e processar os seguintes dados para o pleno funcionamento e que entram no regulamento da LGPD:

- **Login** - Nome, Idade, Email e Senha
- **Dados para análise (Input do Usuário Logado)** - Formulário com o Dataset, dados dos sensores e/ou descrição do lugar e da situação.

Os dados de login serão armazenados em um banco de dados na AWS, que possui certificações de segurança reconhecidas, e será desenvolvido configurações necessárias para aumentar ainda mais a credibilidade. É possível acessar esse documento no site [Amazon RDS - Guia do Usuário](https://docs.aws.amazon.com/pt_br/AmazonRDS/latest/UserGuide/UsingWithRDS.html). Os dados coletados durante o processo de cadastramento de login têm a finalidade principal de garantir a segurança do perfil do usuário. Em casos de esquecimento de senha, o e-mail fornecido durante o cadastro é utilizado como parte do processo de validação em uma plataforma terceira para recuperar a conta. Dessa forma, o uso desses dados visa exclusivamente proteger a integridade e o acesso seguro à conta do usuário, reforçando a confiança e a segurança em nossos serviços.

Os dados para análise, disponíveis por meio de um formulário opcional, não são requisito obrigatório para utilização da plataforma. No entanto, caso o usuário deseje compartilhar informações específicas de seu domínio, ele tem a opção de preencher o documento e enviá-lo aos administradores. Dada a ampla abrangência da solução, os dados submetidos passam por um processo de validação rigoroso para garantir sua procedência, autenticidade e integridade antes de serem aprovados para exposição no Dashboard. É importante ressaltar que a funcionalidade de validação e exposição dos dados é exclusivamente responsabilidade dos administradores da solução, que têm o papel de garantir a segurança e a precisão das informações apresentadas.

## Equidade e Justiça

A análise de Equidade e Justiça pode ser representado como métricas e indicadores que avaliam o acesso igualitário a recursos, serviços e oportunidades, bem como a distribuição justa de benefícios e ônus dentro de uma sociedade, organização ou sistema. Essa abordagem visa promover a inclusão, a diversidade e a dignidade de todos os indivíduos, buscando eliminar disparidades injustas e promover um ambiente de respeito, cooperação e justiça social.

Nesse contexto, consideremos um cenário hipotético: regiões de maior poder econômico, localizadas próximas às principais avenidas com grande atividade empresarial ou fluxo de pessoas, podem receber mais investimentos em infraestrutura e serviços do que as regiões periféricas. Essa disparidade de investimento pode levar a uma desigualdade na distribuição de recursos, incluindo a instalação de sensores para monitoramento.

Por exemplo, enquanto nas áreas periféricas pode haver uma demanda maior por infraestrutura devido a problemas como alagamentos ou deslizamentos de terra, nas cidades maiores esse cenário pode variar. Portanto, é essencial ajustar os recursos de acordo com as necessidades específicas de cada região, garantindo uma abordagem equitativa e justa no planejamento e na distribuição de investimentos em infraestrutura e tecnologia.

Ao considerar a análise de Equidade e Justiça, é importante reconhecer e mitigar essas disparidades, assegurando que todos os cidadãos tenham acesso igualitário aos recursos e serviços essenciais para uma qualidade de vida adequada, independentemente de sua localização geográfica ou status socioeconômico. Essa abordagem não apenas promove a justiça social, mas também fortalece o tecido social e econômico da comunidade como um todo.

## Transparência e consentimento informado

A transparência desempenha um papel crucial ao garantir que os cidadãos compreendam como seus dados estão sendo coletados, armazenados e utilizados. Através de uma apresentação clara e acessível das informações, o dashboard permite que os residentes da cidade compreendam os processos subjacentes e tomem decisões informadas sobre sua participação e envolvimento.

Além disso, o consentimento informado surge como um princípio ético essencial. Os cidadãos devem ter a oportunidade de consentir - ou não - com a coleta e o uso de seus dados pessoais. Isso implica em garantir que os indivíduos estejam plenamente cientes de como suas informações serão utilizadas e que tenham o direito de optar por não participar, se assim desejarem.

Para promover a transparência e o consentimento informado, os termos de serviço da plataforma estarão disponíveis para todos os usuários interessados, acessíveis no rodapé da tela. Ao se cadastrar na plataforma, os usuários serão solicitados a consentir com o uso de seus dados, conforme necessário para o funcionamento da plataforma.

É importante ressaltar que o não consentimento do usuário para o compartilhamento de dados não o impedirá de acessar a solução. No entanto, certas restrições serão aplicadas. Por exemplo, usuários que optarem por não compartilhar seus dados terão acesso apenas ao dashboard público, sem personalizações, e não poderão enviar dados para análise pela plataforma, devido à necessidade de identificação.

Essas medidas garantem que os usuários tenham controle sobre suas informações pessoais e possam fazer escolhas informadas sobre o compartilhamento de dados. Ao mesmo tempo, mantemos o compromisso de proteger a privacidade e a segurança dos dados dos nossos usuários, garantindo uma experiência confiável e transparente para todos.

##  Responsabilidade social

A responsabilidade social analisa o possível impacto (seja positivo ou negativo) que um projeto tem e/ou pode ter em diferentes ambientes sociais. Essa análise é importante, pois, dependendo do impacto que o desenvolvimento possa oferecer, ele precisa se enquadrar em regras, como por exemplo, a LGPD, leis (sejam estaduais, federais ou de qualquer natureza parecida), para evitar problemas sociais, de privacidade ou qualquer violação que afete a integridade física, intelectual, moral ou ética de qualquer pessoa, cnpj, govenamental ou de patrimônio público ou privado.

Procurando análisar o impacto que o projeto pode ofecer, tanto positivamente quanto negativamente, esse forma os possíveis pontos encontrados:

### Impacto Sociais Positivos:

- **Empoderamento da Comunidade:** Ao disponibilizar informações detalhadas sobre condições atmosféricas, qualidade do ar, entre outros, o Dashboard capacita os cidadãos a tomar decisões mais informadas sobre suas atividades diárias e rotinas, como escolher rotas mais seguras ou evitar áreas com alta poluição.

- **Melhoria da Saúde Pública (Exemplo de aplicação da Persona da Secretaria da Saúde):** Fornecer dados sobre poluição do ar, níveis de radiação UV e outras condições ambientais permite que profissionais ou secretarias criem campanhas de prevenção para que os residentes adotem medidas preventivas para proteger sua saúde, como usar protetor solar adequado ou evitar atividades ao ar livre durante períodos de alta poluição.

- **Apoio a Decisões Urbanas:** Os dados coletados podem ajudar as autoridades locais e políticas a tomar decisões mais eficazes em relação ao planejamento urbano, gestão de emergências e mitigação de desastres, contribuindo para uma cidade mais resiliente e segura.


### Impactos Sociais Negativos:

- **Desigualdade Digital:** Pessoas com acesso limitado à internet, familiaridade com tecnologia ou ainda incompatibilidade com a tecnologia desenvolvida(equipamentos antigos e desatualizados) podem ser excluídas do benefício do Dashboard, aprofundando a divisão digital e a desigualdade de acesso à informação.

- **Privacidade e Segurança dos Dados:** A coleta e o armazenamento de dados pessoais podem representar riscos para a privacidade dos cidadãos se não forem implementadas medidas de segurança adequadas, como anonimização de dados e protocolos de segurança robustos ou ainda a inseguraça de pessoas principalmente com mais idade de compartilhar os dados com a internet (No caso do login, por exemplo, que exige o nome completo).

- **Viés nos Dados:** Se os sensores forem implantados de forma desigual pela cidade, pode ocorrer um viés nos dados coletados, favorecendo certas áreas em detrimento de outras. Isso pode levar a decisões injustas ou políticas públicas mal informadas. Isso no tópico explicado acima de Equidade e Justiça.

Entrando um pouco em sobre o comprisso públicos, o Brasil é um dos membros fundadores da Oranização das Nações Unidas (ONU) segundo o site do [gov.br - Ministério das Relações Exteriores](https://www.gov.br/mre/pt-br/delbrasonu/a-missao-do-brasil/a-missao-do-brasil). As ODS (Objetivos de Desenvolvimento Sustentável) são uma coleção de 17 objetivos globais estabelecidos pela ONU em 2015 como parte da Agenda 2030 para o Desenvolvimento Sustentável. Esses objetivos foram desenvolvidos como um apelo universal à ação para acabar com a pobreza, proteger o planeta e garantir que todas as pessoas desfrutem de paz e prosperidade até 2030.

Falando um pouco sobre as ODs's, podemos alinha o projeto a elas (em algumas) nesse pontos:

- ODS 3 - Saúde e Bem-Estar:
    O fornecimento de dados sobre poluição do ar, níveis de radiação UV e outras condições ambientais pode ajudar a promover a saúde pública, permitindo que as pessoas adotem medidas preventivas para proteger sua saúde.

- ODS 11 - Cidades e Comunidades Sustentáveis:
    O uso do Dashboard para apoiar decisões urbanas pode contribuir para o desenvolvimento de cidades mais resilientes e seguras, promovendo o planejamento urbano sustentável e a gestão de emergências.

- ODS 16 - Paz, Justiça e Instituições Eficazes:
    O acesso transparente a informações sobre condições ambientais e urbanas pode promover a transparência, a prestação de contas e a participação cívica, fortalecendo as instituições e contribuindo para a construção de sociedades mais justas e pacíficas.

Com certeza o projeto pode se enquadrar em outras, mas para fins de documentão só seram citadas a 3 a cima.


## Viés e discriminação

Como comentado acima, O viés nos dados, se não for devidamente identificado e mitigado, pode resultar em discriminação e injustiças sociais. Quando os sensores são implantados de forma desigual pela cidade, tendem a favorecer certas áreas em detrimento de outras. Isso pode levar a decisões injustas ou políticas públicas mal informadas, perpetuando desigualdades existentes e aprofundando divisões sociais.

Por exemplo, se os sensores forem predominantemente instalados em bairros mais ricos ou áreas urbanas centrais, os dados coletados podem não refletir adequadamente as condições em áreas mais periféricas ou de baixa renda. Isso pode levar a uma alocação desigual de recursos e serviços públicos, exacerbando disparidades socioeconômicas e de qualidade de vida.

Além disso, é importante destacar que os algoritmos e modelos utilizados na análise dos dados podem conter viéses, refletindo preconceitos existentes na sociedade. Essa tendência pode culminar em decisões automatizadas que discriminam determinados grupos com base em características como raça, gênero, classe socioeconômica e outros fatores. Casos emblemáticos de uso da inteligência artificial que resultaram em discriminação foram documentados em diversos contextos. Para mais informações, recomenda-se a leitura do artigo disponível em:  jornal.usp acesso o site [Inteligência Artificial Utiliza Base de Dados que Refletem Preconceitos e Desigualdades](https://jornal.usp.br/atualidades/inteligencia-artificial-utiliza-base-de-dados-que-refletem-preconceitos-e-desigualdades/)

Para combater o viés e a discriminação nos dados, é fundamental adotar abordagens transparentes e inclusivas em todas as etapas do processo de coleta, análise e tomada de decisões. Isso implica em garantir a representatividade dos dados, considerar perspectivas diversas e envolver as comunidades afetadas no desenvolvimento e implementação de políticas e soluções. Uma distribuição equitativa dos sensores é essencial para garantir que a equidade seja verdadeira e que os dados coletados sejam utilizados de forma consciente e responsável pelos gestores públicos. Essa abordagem não apenas reduzirá os viéses nos dados, mas também promoverá a inclusão e a participação ativa das comunidades na construção de um ambiente urbano mais justo, igualitário, e tambem seguro para soluções que vão utilizar a mesma base de dados.


# Referências
* Universidade de São Paulo (USP). "Inteligência Artificial Utiliza Base de Dados que Refletem Preconceitos e Desigualdades". Disponível em: https://jornal.usp.br/atualidades/inteligencia-artificial-utiliza-base-de-dados-que-refletem-preconceitos-e-desigualdades/. Acessado em 03/2024.

* Ministério das Relações Exteriores (MRE) - Brasil. "A Missão do Brasil na ONU". Disponível em: https://www.gov.br/mre/pt-br/delbrasonu/a-missao-do-brasil/a-missao-do-brasil. Acessado em 03/2024.

* Amazon Web Services (AWS). "Amazon RDS - Guia do Usuário". Disponível em: https://docs.aws.amazon.com/pt_br/AmazonRDS/latest/UserGuide/UsingWithRDS.html. Acessado em 03/2024.