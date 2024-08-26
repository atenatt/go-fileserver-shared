# Servidor de Arquivos em Go com Autenticação Básica

Este projeto implementa um servidor de arquivos simples em Go que serve arquivos a partir de um diretório especificado e oferece autenticação básica para proteger o acesso aos arquivos.

## Funcionalidades

- Servir arquivos de um diretório especificado.
- Proteger o acesso aos arquivos usando autenticação básica (usuário/senha).

## Requisitos

- Go 1.16 ou superior
- Biblioteca externa: [go-http-auth](https://github.com/abbot/go-http-auth)

## Instalação

1. Clone este repositório:

   ```bash
   git clone https://github.com/seuusuario/nome-do-repositorio.git
   cd nome-do-repositorio
   ```

2. Instale a dependência:

   ```bash
   go get github.com/abbot/go-http-auth
   ```

3. Compile o projeto:

   ```bash
   go build -o fileserver main.go
   ```

## Uso

Para rodar o servidor, use o comando:

```bash
go run main.go <diretorio> <porta>
```

- `<diretorio>`: O diretório de onde os arquivos serão servidos.
- `<porta>`: A porta onde o servidor será executado.

Exemplo:

```bash
go run main.go ./shared 8080
```

Neste exemplo, o servidor será iniciado na porta `8080` e servirá os arquivos presentes no diretório `./shared`.

## Autenticação

- O servidor utiliza autenticação básica para proteger o acesso aos arquivos.
- Usuário permitido: `root`
- Senha: `devops1`

## Personalização

Você pode alterar as credenciais de autenticação modificando a função `Secret` no código:

```go
func Secret(user, realm string) string {
    if user == "root" {
        // Password: devops1
        return "$1$8UCs.Hxd$MY6u51wFLuAuget/DZ72j."
    }
    return ""
}
```

Para gerar uma nova senha criptografada, você pode usar a ferramenta `htpasswd`:

```bash
htpasswd -nb root novasenha
```

Substitua a string retornada pela função `Secret` com a nova senha gerada.

## Segurança

Este projeto é para fins educacionais e de uso local. Não é recomendado para uso em produção sem revisões adicionais de segurança, como a implementação de HTTPS.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
```