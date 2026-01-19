# Copilot Instructions – Zipcode Weather Service (Go)

## Contexto do Projeto
Este repositório implementa um serviço HTTP escrito em Go cujo objetivo é receber um CEP brasileiro, identificar a cidade correspondente e retornar o clima atual (temperatura) nos formatos Celsius, Fahrenheit e Kelvin.

O serviço deve ser executável via Docker e compatível com deploy no Google Cloud Run.

Você (Copilot) deve analisar criticamente o código existente e verificar se todos os requisitos descritos neste documento estão sendo rigorosamente cumpridos.

Caso algum requisito não esteja sendo atendido, você deve:
- Indicar claramente o problema
- Apontar onde ocorre (arquivo/função, quando possível)
- Explicar o impacto no comportamento esperado
- Sugerir correção objetiva

---

## Objetivo Funcional
Desenvolver um sistema em Go que:
1. Receba um CEP brasileiro
2. Valide o formato do CEP (8 dígitos numéricos)
3. Consulte a localização via API de CEP
4. Consulte o clima atual da cidade obtida
5. Retorne a temperatura em Celsius, Fahrenheit e Kelvin
6. Seja executável via Docker e publicável no Google Cloud Run

---

## Requisitos Obrigatórios

### 1. Entrada (CEP)
- O sistema DEVE aceitar apenas CEPs com exatamente 8 dígitos numéricos
- CEPs com letras, símbolos ou tamanho diferente devem ser rejeitados

### 2. Validação de CEP
- A validação de formato deve ocorrer antes de qualquer chamada externa
- Em caso de CEP inválido:
  - HTTP Status: 422
  - Response Body (texto puro): invalid zipcode

### 3. Consulta de CEP
- Utilizar a API ViaCEP (ou equivalente)
- Caso o CEP seja válido, porém não encontrado:
  - HTTP Status: 404
  - Response Body (texto puro): can not find zipcode
- O nome da cidade deve ser corretamente extraído da resposta

### 4. Consulta de Clima
- Utilizar a API WeatherAPI (ou equivalente)
- A consulta deve ser baseada na cidade obtida via CEP
- A temperatura base deve ser obtida em Celsius

### 5. Conversão de Temperaturas
As conversões DEVEM seguir exatamente as fórmulas abaixo:

- Fahrenheit:
  F = C * 1.8 + 32

- Kelvin:
  K = C + 273

### 6. Resposta de Sucesso
Em caso de sucesso:
- HTTP Status: 200
- Response Body (JSON):
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}

---

## Requisitos de Arquitetura

### Código
- Linguagem: Go
- Organização clara (handlers, services, clients, etc.)
- Tratamento explícito de erros
- Não mascarar falhas de APIs externas
- Não retornar stack traces ao cliente

### Testes Automatizados
O projeto DEVE conter testes automatizados que cubram:
- CEP inválido
- CEP válido inexistente
- Fluxo de sucesso

Os testes devem ser executáveis via:
go test ./...

### Docker
- Deve existir um Dockerfile funcional
- A aplicação deve iniciar corretamente em container
- A porta deve ser configurável via variável de ambiente PORT

### Docker Compose
- Deve existir um docker-compose.yml (quando aplicável)
- Deve permitir subir a aplicação localmente para testes

---

## Google Cloud Run
- A aplicação DEVE ser stateless e compatível com Cloud Run
- Deve existir evidência de deploy com URL pública funcional
- Porta dinâmica configurada corretamente