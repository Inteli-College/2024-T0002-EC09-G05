#!/bin/bash

status=1

echo "Iniciando execução do script de teste..."
echo "subindo contêineres de teste..."
echo
docker compose -f docker-compose-test.yaml up --build -d > /dev/null
echo
for (( i=1; i<=10; i++ ))
do
    echo "Checkando existência do container de testes... ($i de 10)"
    
    # Executa o script de teste unitário dentro do container
    if docker container inspect -f '{{.State.Running}}' tests; then
        sleep 2
        if docker exec  tests go test; then
            echo "Teste executado com sucesso."
            status=0
        else
            echo "Teste falhou"
        fi
        break
    else
        echo "Falha ao encontrar container de teste, tentando novamente em 10 segundos..."
        sleep 10
    fi
    
    # Se atingir a última tentativa e falhar, mostra uma mensagem
    if [ $i -eq 10 ]; then
        echo "Tentativas de execução do script de teste falharam. Verifique os contêineres e o ambiente."
    fi
done



docker compose -f docker-compose-test.yaml down

exit $status
