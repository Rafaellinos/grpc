# gRPC Learning

- Remote Protocol Call
  - Padrão de comunicacao para chamar funcao ou metodo
  - Complexidade de rede é abstraida (sem abrir ‘socket’, etc)
- gRPC (Google, mas na vdd cada release tem um significado como garcia, good, green etc)
  - substituir o Stubby por uma implementacao utilizando HTTP/2
  - CNCF estabelece principios para projeto gRPC
  - Nao funciona com HTTP/1.1
- HTTP/2
  - Somente formato binario (formato menor)
  - Multiplexacao 
    - Unica comunicacao pode-se enviar varios requestes em paraleleto, ao contratio da 1.1
    - Compressao header com HPACK (enviar apenas cabeçalhos que mudaram)
  - HTTP/1.1 tem pipelining, mas é uma req por ordem
- Server Push
  - Servidor envia para o cliente mais coisas do que ele pediu (eg ja manda arquivos CSS e JS)

## Protocol buffer

- Serializacao padrao usado pelo gRPC
- Independente do gRPC, mas é o mais utilizado
- Tipadas e estruturados
- Compacto, cross-lang, otimizado e compacto
- Fornece IDL (interface Definition Language)
  - Mensagens, servicos e arquivo `.proto`.
- Diagrama: 
  - Crie .proto data structure -> Gere codigo usando protoc compiler -> Compile PB no seu projeto
  - PB classes será usado para serdes
  - Parecido com Avro

eg golang

```protobuf
syntax = "proto3";

option go_package = "github.com/Rafaellinos/grpc-go";

package helloworld;

message HelloRequest {
    string name = 1;
}
message HelloReply {
    string message = 1;
}
service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}
```

> On Mac M1, run these commands: `brew install protoc-gen-go` install go generator and add to your ~/.zshrc `export PATH="$PATH:$(go env GOPATH)/bin"`

```bash
protoc --go_out=. --go_opt=module=github.com/Rafaellinos/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/Rafaellinos/grpc-go helloworld/proto/helloworld.proto
```

## Golang comandos

- `go mod download` download dependencies, nao modifica go.mod ou go.sum
- `go mod tidy` limpa e atualiza go.mod e go.sum de acordo com dependencias
- go.mod = arquivo que define nome do modulo, versao e dependencias
- go.sum = arquivo que armazena checksums de dependencias baixadas
- `cd helloworld && go build -o bin/server ./server` buildar o projeto
