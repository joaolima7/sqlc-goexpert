# Guia de Instalação e Uso do SQLC e Migrations
## Instalação do Golang Migrate
Para gerenciar migrações de banco de dados em projetos Go, utilizamos a ferramenta `golang-migrate`. Para instalá-la, execute o seguinte comando:

```bash
brew install golang-migrate
```
Este comando utiliza o Homebrew, um gerenciador de pacotes para macOS, para instalar a ferramenta `golang-migrate`, que é essencial para criar e aplicar migrações de banco de dados.
## Criação de Migrações
Após instalar o `golang-migrate`, você pode criar uma nova migração utilizando o seguinte comando:

```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```
Este comando cria um novo arquivo de migração com a extensão `.sql` no diretório `sql/migrations`. A opção `-seq` indica que as migrações serão sequenciais, e `init` é o nome da migração. O arquivo gerado conterá dois blocos: um para aplicar a migração (`up`) e outro para revertê-la (`down`).

## Aplicação de Migrações
Para aplicar as migrações criadas, utilize o seguinte comando:

```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/mydb" -verbose up
```
Este comando aplica todas as migrações pendentes no banco de dados especificado. A opção `-verbose` fornece informações detalhadas sobre o processo de migração. Certifique-se de substituir a string de conexão do banco de dados conforme necessário para o seu ambiente.
## Reversão de Migrações
Para reverter a última migração aplicada, você pode usar o seguinte comando:

```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/mydb" -verbose down
```
Este comando reverte a última migração aplicada no banco de dados. Assim como na aplicação, a opção `-verbose` fornece detalhes sobre o processo de reversão.

## Explicação do Makefile
O Makefile presente no projeto contém três comandos principais para facilitar o gerenciamento de migrações:
- `createmigration`: Este comando cria uma nova migração com a extensão `.sql` no diretório `sql/migrations`. Ele utiliza o comando `migrate create` com as opções apropriadas para definir o nome e a sequência da migração.
- `migrate`: Este comando aplica todas as migrações pendentes no banco de dados especificado. Ele utiliza o comando `migrate up` com o caminho das migrações e a string de conexão do banco de dados.
- `migrate-down`: Este comando reverte a última migração aplicada no banco de dados. Ele utiliza o comando `migrate down` com o caminho das migrações e a string de conexão do banco de dados.
Esses comandos são definidos como `.PHONY`, o que significa que eles não correspondem a arquivos reais, mas sim a ações que devem ser executadas quando chamados. Isso garante que o Makefile sempre execute os comandos, independentemente da existência de arquivos com os mesmos nomes.