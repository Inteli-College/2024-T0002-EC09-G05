# Testes

O script `run_tests.sh` na pasta `src/` é responsável por executar todos os testes do projeto. Ele executa o sistema completo, além de um container Go responsável por executar os testes unitários e de integração.

Para rodar os testes na sua máquina, execute o script com:

`cd src && ./run_tests.sh`

![commit](img/tests.png)

### Resultados esperados

É esperado que todos os testes passem. Verifique código de saída do script. Qualquer código diferente de 0 indica falha nos testes.

### CI/CD

O script `run_tests.sh` é executado automaticamente pelo Github Actions do projeto. O símbolo de check ao lado de um commit evidência o resultado dos testes para aquele commit específico:

![tests](img/commit.png)
