# Korp_Teste_AntonioWills

Este projeto é uma implementação parcial do desafio técnico para a Korp/Viasoft, focada em uma arquitetura robusta de microsserviços utilizando **Golang** e **Docker**.

# Arquitetura do Projeto

O sistema foi desenhado para ser escalável e resiliente, utilizando os seguintes componentes:

- **API Gateway (Nginx):** Ponto único de entrada que roteia as chamadas para os serviços específicos (/api/*inventory* e */api/invoices*).
- **Inventory Service**: Gerencia o cadastro de produtos e o controle de saldo em estoque.
- **Incoice Service:** Responsável pela criação de notas fiscais e pela orquestração do processo de "impressão", que consome o saldo do estoque.
- **Bancos de Dados:** Instâncias independentes de PostgreSQL para cada serviço, evitando o acoplamento de dados.

# Tecnologias Utilizadas

- **Linguagem:** Golang (Framework Gin).
- **Banco de Dados:** PostgreSQL 16.
- **Containerização:** Docker e Docker Compose.
- **Proxy/Gateway:** Nginx.

# Como Executar

Para subir todo o ecossistema de backend, utilize o comando:
```bash
docker-compose up --build
```

Os serviços estarão disponíveis em:
- Gateway: *http://localhost:80*
- Inventory API: *http://localhost:8080/api/inventory*
- Invoices API: *http://localhost:8081/api/invoices*

# Tratamento de Erros e Resiliência

O sistema foi projetado para lidar com falhas distribuídas:

- Isolamento: A falha de um banco de dados não derruba o serviço adjacente.
- Healthchecks: O Docker Compose monitora a saúde dos bancos antes de iniciar os serviços.
- Códigos de Status HTTP: O InvoiceController mapeia erros de negócio para status específicos, como 409 Conflict para estoque insuficiente e 503 Service Unavailable caso o serviço de inventário esteja fora do ar.