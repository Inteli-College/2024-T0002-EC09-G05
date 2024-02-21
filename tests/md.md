
# Teste da Arquitetura

Elaborar testes para a arquitetura proposta é essencial para garantir que o sistema atenda aos requisitos do projeto e funcione corretamente. A seguir, são apresentados os testes a serem realizados para validar a arquitetura proposta.

## Testes Manuais

### Teste de Conexão MQTT

**Objetivo:** Verificar se o cliente MQTT consegue publicar mensagens nos tópicos corretos e se o broker MQTT encaminha corretamente as mensagens para o AWS IoT Core.

**Passos:**
1. Configurar um cliente MQTT para publicar mensagens nos tópicos especificados.
2. Verificar se as mensagens são recebidas pelo broker MQTT.
3. Verificar se as mensagens são encaminhadas corretamente para o AWS IoT Core.

**Ferramentas:** Cliente MQTT (por exemplo, MQTT Explorer)