<!-- markdownlint-disable MD007 MD031 MD034 -->

# Go Expert Labs - Cloud Run challenge

Aplicação web em Go que receba um CEP, identifica a cidade e retorna o clima atual em Celsius, Fahrenheit e Kelvin.

## Executando em PROD (Cloud Run)

A aplicação está disponível para acesso no serviço Google Cloud Run e pode ser acessado seguindo os seguintes parâmetros:

-   **Endpoint:** https://cloudrun-goexpert-hkg7fv4fwa-rj.a.run.app/
-   **Método:** GET
-   **Query Params:**
    -   **zipcode:** CEP a ser consultado

Exemplo de requisição com `curl`:

```bash
curl -X GET https://cloudrun-goexpert-hkg7fv4fwa-rj.a.run.app/?zipcode=01153000

{
    "temp_C": 15.3,
    "temp_F": 59.5,
    "temp_K": 288.3
}
```

## Executando localmente (dev)

### Requisitos

1. Clone o repositório;
2. Execute o comando `cp .env.example .env` para criar o arquivo de variáveis de ambiente;
3. Edite o novo arquivo `.env` e insira sua chave de acesso à API do [WeatherAPI](https://www.weatherapi.com/) à variável `WEATHER_API_KEY`;

### Via Docker

1. Execute o comando `docker compose up` para realizar o build do container e iniciar a aplicação na porta declarada no arquivo `.env`;

## Documentação do endpoint

### Request

| Endpoint | Descrição                                 | Método | Parâmetro |
| -------- | ----------------------------------------- | ------ | --------- |
| /        | Calcula a temperatura atual em uma cidade | GET    | zipcode   |

### Response

-   Sucesso:

    -   **Código:** 200
    -   **Body:**
        ```json
        {
        	"temp_C": 14.2,
        	"temp_F": 57.6,
        	"temp_K": 287.2
        }
        ```

-   CEP não encontrado:

    -   Ex: 00000-000
    -   **Código:** 404
    -   **Body:**
        ```json
        {
        	"message": "zipcode not found"
        }
        ```

-   CEP inválido:
    -   Ex: 00000-000a
    -   **Código:** 422
    -   **Body:**
        ```json
        {
        	"message": "invalid zipcode"
        }
        ```
