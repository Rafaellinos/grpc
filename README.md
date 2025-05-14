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
- Contratos desatualizados n geram erros, mas campo sao ignorados

eg golang

```protobuf
syntax = "proto3";

option go_package = "github.com/Rafaellinos/grpc/helloworld";

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
cd helloworld && protoc --go_out=. --go_opt=module=github.com/Rafaellinos/grpc/helloworld --go-grpc_out=. --go-grpc_opt=module=github.com/Rafaellinos/grpc/helloworld ./proto/helloworld.proto
```

### Tipos Escalares - Numericos

- float
- int32 (para valores negativos, use sint32)
- int64 (para valores negativos, use sint64)
- uint32 (sem sinal)
- uint64 (sem final)
- sint32
- sint64
- fixed32 (valores fixos, bom para valores grandes 2^28)
- fixed64 (valores fixos, bom para valores grande 2^58)
- sfixes32 (sempre 4 bytes)
- sfixed64 (sempre 8 bytes)
- bool
- string (utf8 ou ASCI, n pode ser maior que 2ˆ32)
- bytes (n pode ser maior que 2^32)

- Valores default: numerico = 0, bytes = vazio, string = vazio, bool = false
- Quando o campo é valor default, n é enviado na msg, mesmo que tenha sido setado.

### Cardinallidade

- campos unicos (singular)
  - implicit
  - optional
- Colecoes: repeated
- Mapas: maps

```protobuf
syntax = "proto3";
message UserDetails {
  string name = 1;
  uint32 age = 2;
  optional bool is_active = 3;
}
```

Quando campos esta com optional, no golang será um ponteiro (`*isActive`) entao é possivel verificar se é `nil`
Campos do tipo Message (outra referencia) é um ponteiro por default

### Protobuf decoder / wire format

É possível "decodar" o hexadecimal para obter os detalhes da mensagem.

Para gerar a mensagem em formato binario, usa o comando abaixo:

```shell
protoc --encode=helloworld.GreetRequest helloworld.proto < request.txt > wire_format.bin
```

Arquivo protobuf text format

```text
type: 0
users: [{
    name: "Rafael"
    age: 35
},
{
    name: "Yasmin"
    age: 27
    is_active: true
}]
```

- Wire format (tipos de arquivo binarios)
  - VARINT (int32, int64, sint64, bool, enum)
  - I64 (fixed64, sfixed64, double)
  - LEN (string, bytes, embeedded msgs, packed repeated fields)
  - I32 (fixed32, sfixed32, double)
- Estrutura
  - field number | type (tipo do campo) | payload
  - Nao é possivel saber o nome do campo pelo wire format, apenas o indice (field number)
  - Nao é possivel saber com exatidao o tipo do campo, apenas com o contrato (.proto)
- Field numbers (ids/index de campos, como `string name = 1`)
  - Nao podem se repetir dentro da mesma msg
  - 19_000 a 19_999 eh reservado
  - é possivel reservar numeros
  - Quando adicionar novo campo, boa pratica será usar `optional`
  - !warning cuidado ao mudar o nome e tipo do campo e manter o mesmo ID, pois o servidor vai considerar como mesmo campo
  - use `reserved <ID>` para reservar campos, util para evitar erros quando removemos campos

## Modelos de comunicacao gRPC

- unica mensagem, mais simples
- server streaming (comunicacao constante, mantem comunicacao aberta)
- client streaming
- bidirectional streaming


## Java spring grpc (stockmarket)

- must have the plubin `protobuf-maven-plugin`
- run `mvn generate-sources` to generate classes based on .proto file
- :warning: Remember to add target/generated-sources/protobuf/grpc-java and /java in /target as source root

eg em java:

```protobuf
syntax = "proto3";

option java_package = "br.com.rafaellino.stockmarket";
option java_multiple_files = true;

package stock_market;


service StockPrice {
  rpc GetStockPrice (StockRequest) returns (StockResponse);
}

message StockRequest {
  string symbol = 1;
}

message StockResponse {
  string symbol = 1;
  double price = 2;
  int64 timestamp = 3;
}
```


## Python grpc (stockmarket_consumer)

Entre no directory stockmarket_consumer

```shell
python -m venv .venv # crie o virtual environment
```

```shell
source .venv/bin/activate # ative o virtual environment
```

```shell
pip install -r requirements.txt # instale dependencias
```

```shell
python main.py # com a aplicacao stockmarket rodando, execute o main.py
```

Para gerar o arquivo com base no proto file:

```shell
python -m grpc_tools.protoc -I . --python_out=. --grpc_python_out=. stock_market.proto
```

## Golang comandos

- `go mod download` download dependencies, nao modifica go.mod ou go.sum
- `go mod tidy` limpa e atualiza go.mod e go.sum de acordo com dependencias
- go.mod = arquivo que define nome do modulo, versao e dependencias
- go.sum = arquivo que armazena checksums de dependencias baixadas
- `cd helloworld && go build -o bin/server ./server` buildar o projeto


## Timeout vs Deadline

Tempo maximo que o cliente espera pela resposta do servidor. Quando o client chega ao deadline, recebe o erro: `DEADLINE_EXCEEDED`

## Unario (unary)

- Uma unica mensagem, request e response, parecido com HTTP/REST

## Streaming


e.g.:

```protobuf
syntax = "proto3";

message HelloRequest {
  string name = 1;
}
message HelloReply {
  string message = 1;
}
service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
  rpc SayHelloStream (HelloRequest) returns (stream HelloReply);
}
```

- chave `stream` especifica o tipo


