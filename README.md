# ANote
Um webapp de anotações com a possibilidade de menção de usuários e de outras anotações, comentário e compartilhamento em comunidade. Essa aplicação possuirá um mecanismo de pesquisa eficiente para facilitar a localização de anotações, os filtros pooderam ser feitas por data de publicação, autor, conteúdo, título, menção e comunidade, separados ou ao mesmo tempo.

## Integrantes
- Daniel Henrique Toledo Santos: Backend
- Julio Cesar de Paula Ferreira: Frontend
- Samuel Henrique Miranda Alves: Full-Stack

## Tecnologias
- React
- Go
- PostgreSQL
- Elasticsearch
- Docker

## How to run

Para rodar o sistema será necessário que você tenha instalado docker em sua máquina. 

Com o docker instalado basta na pasta root (```../ANote/```) do projeto rodar o comando ```docker compose up -d``` para inicializar o projeto (Frontend + Backend + Elasticsearch + PostgreSQL + Ferramentas de interação com as bases de dados). O frontend pode ser acessado em localhost:3000.

OBS: Isso pode demorar um bom tempo na primeira vez.

## Architecture

![image](https://github.com/danhenrik/ANote/assets/42657692/84b2bccc-fd42-4256-be59-ae25ad649e52)

A arquitetura consiste em um client e um server monolítico. A comunicação entre o client e o server é feita via HTTP.

Para a funcionalidade de pesquisa foi utilizado elasticsearch (ES), para replicar os dados do banco de dados PostgreSQL (PG) para o elasticsearch foi desenvolvido um componente chamado es_replicate. 

Este componente se conecta ao PostgreSQL e se inscreve no tópico de notificação "es_replicate" criado no próprio PosgreSQL, esse tópico recebe notificações toda vez que tabelas de interesse sofrem alterações no PostgreSQL, isto é feito por meio do uso de triggers. 

Quando uma notificação é recebida o componente "es_replicate" busca no banco de dados as mudanças desde sua última atualização, isto é possível devido ao mecanismo de replication slot disponível neste SGBD que possibilita o acesso ao WAL (Write Ahead Log) em formato json, o que torna fácil seu consumo.

Write Ahead Log em uma visão simplificada é um componente do PostgreSQL que guarda todas as mudanças feitas no banco de dados.

Uma forma mais fácil de visualizar essa interação pode ser vista a seguir:

![image](https://github.com/danhenrik/ANote/assets/42657692/ef257bd6-5f84-41bb-869a-ddb1bfbfae42)
1. Server escreve no banco
2. Antes mesmo da escrita de Server a escrita é persistida no WAL
3. Ao finalizar a escrita de Server o trigger é ativado e o componente "es_replicate" é notificado
4. "es_replicate" busca as mudanças no WAL 
5. "es_replicate" interpreta as mudanças de acordo com a tabela onde foi feita a mudança e replica essa mudança nos documentos

## Product Backlog

1. Como usuário, quero me cadastrar no sistema
3. Como usuário cadastrado, quero me logar no sistema
4. Como usuário cadastrado, quero recuperar minha senha
5. Como usuário cadastrado logado, quero atualizar minha senha
6. Como usuário cadastrado logado, quero atualizar meu email
7. Como usuário, quero me logar usando minha conta Google ou Facebook
8. Como usuário logado, quero criar comunidades
9. Como usuário logado, quero editar comunidades (Nome, Privado/Público)
10. Como usuário logado, quero excluir comunidades sem membros ou posts
11. Como usuário logado, quero criar novas anotações
12. Como usuário, quero ter acesso a anotações em comunidades públicas
13. Como usuário, quero filtrar as anotações baseado em data de publicação, autor, conteúdo, título, menção e comunidade
14. Como usuário logado, quero acessar minhas anotações privadas
15. Como usuário logado, quero editar anotações das quais sou dono
16. Como usuário logado, quero excluir anotações das quais sou dono
17. Como usuário logado, quero seguir outros usuários
18. Como usuário logado, quero compartilhar uma anotação pública
19. Como usuário logado, quero comentar anotações
20. Como usuário logado, quero curtir anotações

## Protótipo

<a href="https://www.figma.com/file/Xcmx6LAl2FMChm8Mb5psoX/ANote?type=design&node-id=0%3A1&mode=design&t=iqdVrLYjdwgdlkHh-1">Protótipo</a>

## 1st Sprint Backlog
1. Como usuário, quero me cadastrar no sistema

Tarefas e Responsáveis:
  - Instalar tecnologias (React, PostgreSQL, Docker) [Daniel];
  - Criar tabelas relacionadas ao usuário no banco e configurá-las no sistema [Samuel];
  - Criar e testar rota que implementa a tela de cadastro do usuário [Samuel];
  - Implementar no backend a lógica de cadastrar usuário [Daniel];
  - Criar tela de cadastrar usuário [Julio];

2. Como usuário, quero me logar usando minha conta Google ou Facebook

Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [Samuel];
  - Criar e testar rota que implementa a tela de logar usuário cadastrado [Samuel];
  - Criar e testar rota que implementa a tela de logar usuário pelo Google ou Facebook [Samuel];
  - Implementar no backend a lógica de logar usuário cadastrado [Daniel];
  - Implementar no backend a lógica de logar usuário pelo Google ou Facebook [Daniel];
  - Criar tela de logar usuário cadastrado [Julio];
  - Criar tela de logar usuário pelo Google ou Facebook [Julio];

3. Como usuário cadastrado, quero recuperar minha senha

Tarefas e Responsáveis:
  - Criar e testar rota que recupera a senha do usuário [Samuel];
  - Implementar no backend a lógica de recuperar senha do usuário [Daniel];
  - Criar tela de recuperar senha do usuário [Julio];

4. Como usuário logado, quero criar, editar ou excluir anotações

Tarefas e Responsáveis:
  - Criar tabelas relacionadas às anotações no banco e configurá-las no sistema [Samuel];
  - Criar e testar rotas que cria, edita e exclui anotações [Samuel];
  - Implementar no backend a lógica de cadastrar, editar e excluir usuário [Daniel];
  - Criar tela inicial de criar, editar ou excluir anotações [Julio];

5. Como usuário, quero filtrar as anotações baseado em data de publicação, autor, conteúdo, título, menção e comunidade

Tarefas e Responsáveis:
  - Criar e testar rota que filtra anotações [Samuel];
  - Implementar no backend a lógica de filtrar anotações [Daniel];
  - Criar tela de filtrar anotações [Julio];

6. Como usuário logado, quero criar, edtiar ou excluir comunidades

Tarefas e Responsáveis:
  - Criar tabelas relacionadas às comunidades no banco e configurá-las no sistema [Samuel];
  - Criar e testar rotas que implementam a criação, edição e exclusão de comunidades [Samuel];
  - Implementar no backend a lógica de criar, editar e excluir comunidades [Daniel];
  - Criar tela de cadastrar, editar e excluir comunidades [Julio];

7. Como usuário logado, quero compartilhar uma anotação pública

Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [Samuel];
  - Criar e testar rota que implementa o compartilhamento de anotações [Samuel];
  - Implementar no backend a lógica de compartilhar anotações [Daniel];
  - Criar tela de compartilhar anotações [Julio];

8. Como usuário logado, quero curtir/comentar anotações

Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [Samuel];
  - Criar e testar rotas responsáveis por curtir e comentar anotações [Samuel];
  - Implementar no backend a lógica de curtir e comentar anotações [Daniel];
  - Criar tela de curtir e comentar anotações [Julio];
