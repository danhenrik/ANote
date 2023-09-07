# ANote
Um webapp de anotações com a possibilidade de menção de usuários e de outras anotações, comentário e compartilhamento em comunidade. Essa aplicação possuirá um mecanismo de pesquisa eficiente para facilitar a localização de anotações, os filtros pooderam ser feitas por data de publicação, autor, conteúdo, título, menção e comunidade, separados ou ao mesmo tempo.

## Integrantes
- Daniel Henrique Toledo Santos: Backend
- Julio Cesar de Paula Ferreira: Full-Stack
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

## 1st Sprint Backlog
1. Como usuário, quero me cadastrar no sistema
Tarefas e Responsáveis:
  - Instalar tecnologias (React, PostgreSQL, Docker) [];
  - Criar tabelas relacionadas ao usuário no banco e configurá-las no sistema [];
  - Criar e testar rota que implementa a tela de cadastro do usuário [];
  - Implementar no backend a lógica de cadastrar usuário [];
  - Criar tela de cadastrar usuário [];

2. Como usuário, quero me logar usando minha conta Google ou GitHub
Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [];
  - Criar e testar rota que implementa a tela de logar usuário cadastrado [];
  - Criar e testar rota que implementa a tela de logar usuário pelo Google ou Github [];
  - Implementar no backend a lógica de logar usuário cadastrado [];
  - Implementar no backend a lógica de logar usuário pelo Google ou Github [];
  - Criar tela de logar usuário cadastrado [];
  - Criar tela de logar usuário pelo Google ou Github [];

3. Como usuário cadastrado, quero recuperar minha senha
Tarefas e Responsáveis:
  - Criar e testar rota que recupera a senha do usuário [];
  - Implementar no backend a lógica de recuperar senha do usuário [];
  - Criar tela de recuperar senha do usuário [];

4. Como usuário logado, quero criar, editar ou excluir anotações
Tarefas e Responsáveis:
  - Criar tabelas relacionadas às anotações no banco e configurá-las no sistema [];
  - Criar e testar rotas que cria, edita e exclui anotações [];
  - Implementar no backend a lógica de cadastrar, editar e excluir usuário [];
  - Criar tela inicial de criar, editar ou excluir anotações [];

5. Como usuário, quero filtrar as anotações baseado em data de publicação, autor, conteúdo, título, menção e comunidade
Tarefas e Responsáveis:
  - Criar e testar rota que filtra anotações [];
  - Implementar no backend a lógica de filtrar anotações [];
  - Criar tela de filtrar anotações [];

6. Como usuário logado, quero criar, edtiar ou excluir comunidades
Tarefas e Responsáveis:
  - Criar tabelas relacionadas às comunidades no banco e configurá-las no sistema [];
  - Criar e testar rotas que implementam a criação, edição e exclusão de comunidades [];
  - Implementar no backend a lógica de criar, editar e excluir comunidades [];
  - Criar tela de cadastrar, editar e excluir comunidades [];

7. Como usuário logado, quero compartilhar uma anotação pública
Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [];
  - Criar e testar rota que implementa o compartilhamento de anotações [];
  - Implementar no backend a lógica de compartilhar anotações [];
  - Criar tela de compartilhar anotações [];

8. Como usuário logado, quero curtir/comentar anotações
Tarefas e Responsáveis:
  - Criar tabelas da história no banco e configurá-las no sistema [];
  - Criar e testar rotas responsáveis por curtir e comentar anotações [];
  - Implementar no backend a lógica de curtir e comentar anotações [];
  - Criar tela de curtir e comentar anotações [];
