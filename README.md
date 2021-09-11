# Buttress Server

_Buttress is the __Access Control as a Service__ solution based on [Casbin][anchor-casbin]._

---

__Buttress Server__ provides a gRPC interface as an access control management service.

## Supported Databases

| Name         | Version                |
| :----------- | :--------------------- |
| `PostgreSQL` | `>=9.0.0`              |
| `MySQL`      | `>=5.7.0`, `>=8.0.0`   |
| `MongoDB`    | `>=4.0.0`              |

## Available Clients and Implementations

- [gRPC Protocol Buffers][anchor-br-protos]
- [Official Golang Client][anchor-brc-go]

[anchor-casbin]: https://casbin.org
[anchor-br-protos]: https://github.com/merajsahebdar/buttress-implementation-go
[anchor-brc-go]: https://github.com/merajsahebdar/buttress-client-go
