# fc-go-gcr-weatherCEP

Sistema em Go que recebe um CEP válido de 8 dígitos, identifica a cidade e retorna o clima atual em Celsius, Fahrenheit e Kelvin.

## URL do serviço no Cloud Run

```
https://fc-go-gcr-weathercep-yjkoxdbpzq-uc.a.run.app
```

## Uso da API

### Sucesso (200)

```
GET /cep/01001000/weather
```

```json
{ "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }
```

### CEP inválido (422)

```
GET /cep/123/weather
```

```
invalid zipcode
```

### CEP não encontrado (404)

```
GET /cep/99999999/weather
```

```
can not find zipcode
```

## Como rodar localmente com Docker

```bash
docker build -t weathercep .
docker run -p 8080:8080 -e API_KEY=SUA_CHAVE_WEATHERAPI weathercep
```

A aplicação estará disponível em `http://localhost:8080`.

## Como rodar os testes

```bash
API_KEY=SUA_CHAVE_WEATHERAPI go test ./...
```