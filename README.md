# goexpert-weather-api
Exercise Weather API for postgraduate Go Expert

## Description

### Objetivo

Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

### Requisitos

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve responder adequadamente nos seguintes cenários:
    - Em caso de sucesso:
        Código HTTP: 200
        Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
    - Em caso de falha, caso o CEP não seja válido (com formato correto):
        Código HTTP: 422
        Mensagem: invalid zipcode
    - ​​​Em caso de falha, caso o CEP não seja encontrado:
        Código HTTP: 404
        Mensagem: can not find zipcode
    - Deverá ser realizado o deploy no Google Cloud Run.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.
