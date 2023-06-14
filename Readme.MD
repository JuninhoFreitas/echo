# Lixão Go Projects

![Go Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1920px-Go_Logo_Blue.svg.png)

> Praticando tudo que é possível fazer com Go!

Bem-vindo ao repositório Lixão Go Projects! Este é um conjunto de projetos incríveis criados para ajudá-lo a aprender e aprimorar suas habilidades em Go. Cada projeto tem uma funcionalidade diferente, mas todos foram desenvolvidos para que você possa praticar e explorar todo o potencial da linguagem Go.



## Roadmap

Este repositório segue o roadmap do Go fornecido por [roadmap.sh](https://roadmap.sh/golang). O roadmap oferece um guia abrangente para você seguir à medida que explora os recursos e conceitos do Go. Sinta-se à vontade para navegar pelo roadmap e escolher os projetos que mais lhe interessam.

## Ponto do Roadmap atual
![Aprendendo Echo](https://i.imgur.com/kFwwOHm.png)
## Projeto

Aqui estão alguns dos conhecimentos explorados neste repositório:

- **Echo:** Usei esse web framework para fazer um crud básico, sem banco de dados nem nada, futuramente pretendo fazer os testes unitários e adicionar conexões a bancos, assim como filas, testes de integração e feature, mas por enquanto, o objetivo era somente testar o framework.

- **Arrays** Foi útil esse projeto para que eu entendesse um pouco mais sobre arrays, como atualizar, como procurar com eficiência, primeiramente eu tinha feito o FindById alocando uma váriavel para cada laço do range, mas li que isso gastava memória e o mais interessante seria usar o proprio index da lista, e assim está hoje.

- **JSON** Também explorei um pouco de Json no GO para tratar os bodys 

- **Strconv** Foi interessante explorar um pouco mais sobre conversão de tipos no go, algo que eu tinha facilmente no Javascript, mas aqui é mais complicado.

## Como usar

- Primeiro suba o docker-compose
```
docker-compose up -d
```
- Conecte ao container
```
docker exec -it echoapp bash
```
- Inicie o Servidor
```
go run server.go
```
- Faça suas requisições(dentro de /har-collections tem um arquivo para ser importado no insomnia ou postman)
```curl
//Hello World
curl --request GET \
  --url http://localhost:3000/
//Get Users
curl --request GET \
  --url http://localhost:3000/users
//Create User
curl --request POST \
  --url http://localhost:3000/users \
  --header 'Content-Type: application/json' \
  --data '{
	"id": 1,
	"name": "Juninho",
	"age": 20
}'
//Update User
curl --request PUT \
  --url http://localhost:3000/users/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"id": 1,
	"name": "Oscar",
	"age": 25
}'
//Delete User
curl --request DELETE \
  --url http://localhost:3000/users/1
```
#

Vamos lá! Desafie-se e comece a praticar Go hoje mesmo! 🚀

Texto gerado pelo chat gpt, ficou legal né?
