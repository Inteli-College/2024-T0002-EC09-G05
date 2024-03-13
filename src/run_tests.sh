#!/bin/bash

status=1

echo "Iniciando execução do script de teste..."
echo "subindo contêineres de teste..."
echo
docker compose -f docker-compose-test.yaml  up --build -d --quiet-pull  >> /dev/null
echo
i=0
while [ $i -le 10 ]; do
    i=$((i+1))
    echo
    echo "Checkando existência do container de testes... ($i de 10)"
    
    if docker container inspect -f '{{.State.Running}}' tests; then
        echo "Rodando testes... aguarde..."
        if docker exec -w /app tests go test; then
            echo
            echo "--> Teste executado com sucesso."
            status=0
        else
            echo "Teste falhou"
        fi
        break
    else
        echo "Falha ao encontrar container de teste, tentando novamente em 10 segundos..."
        sleep 10
    fi
    
    if [ $i -eq 10 ]; then
        echo "Tentativas de execução do script de teste falharam. Verifique os contêineres e o ambiente."
        break
    fi
done


echo "Limpando contêineres..."
echo
docker compose -f docker-compose-test.yaml down

exit $status
