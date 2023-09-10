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

## Product Backlog

1. Como usuário, quero me cadastrar no sistema
3. Como usuário cadastrado, quero me logar no sistema
4. Como usuário cadastrado, quero recuperar minha senha
5. Como usuário cadastrado logado, quero atualizar minha senha
6. Como usuário cadastrado logado, quero atualizar meu email
7. Como usuário, quero me logar usando minha conta Google ou GitHub
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

2. Como usuário, quero me logar usando minha conta Google ou GitHub
Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [Samuel];
  - Criar e testar rota que implementa a tela de logar usuário cadastrado [Samuel];
  - Criar e testar rota que implementa a tela de logar usuário pelo Google ou Github [Samuel];
  - Implementar no backend a lógica de logar usuário cadastrado [Daniel];
  - Implementar no backend a lógica de logar usuário pelo Google ou Github [Daniel];
  - Criar tela de logar usuário cadastrado [Julio];
  - Criar tela de logar usuário pelo Google ou Github [Julio];

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
