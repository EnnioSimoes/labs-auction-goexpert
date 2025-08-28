# Concorrência com Golang - Leilão

Objetivo: Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

Toda rotina de criação do leilão e lances já está desenvolvida, entretanto, o projeto clonado necessita de melhoria: adicionar a rotina de fechamento automático a partir de um tempo.

Para essa tarefa, você utilizará o go routines e deverá se concentrar no processo de criação de leilão (auction). A validação do leilão (auction) estar fechado ou aberto na rotina de novos lançes (bid) já está implementado.

Você deverá desenvolver:

Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente;
Uma nova go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e que deverá realizar o update, fechando o leilão (auction);
Um teste para validar se o fechamento está acontecendo de forma automatizada;

Dicas:

Concentre-se na no arquivo internal/infra/database/auction/create_auction.go, você deverá implementar a solução nesse arquivo;
Lembre-se que estamos trabalhando com concorrência, implemente uma solução que solucione isso:
Verifique como o cálculo de intervalo para checar se o leilão (auction) ainda é válido está sendo realizado na rotina de criação de bid;
Para mais informações de como funciona uma goroutine, clique aqui e acesse nosso módulo de Multithreading no curso Go Expert;
 
Entrega:

O código-fonte completo da implementação.

Documentação explicando como rodar o projeto em ambiente dev.

Utilize docker/docker-compose para podermos realizar os testes de sua aplicação.

## Run

docker-compose up -d

* App: localhost:8080
* MongoDB: localhost:27017
* Mongo Express: localhost:8081
  * user: admin
  * senha: pass

Utilize o Mongo Express (http://localhost:8081) para facilitar o acompanhamento da manipulação dos dados.

As chamadas REST estão em api/api.http

1. Crie um usuário para participar do leilão
2. Crie um Leilao
3. Envie Lances
4. Verifique o status do leilão (No momento o encerramento ocorre após 5 minutos)

## Run Test
Entre no container
```docker exec -ti labs-auction-goexpert-app-1 bash ```

Dentro do container rode
```go test ./internal/infra/database/auction -v -run TestCreateAuction```