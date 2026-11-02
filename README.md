# Xbox 360 Licensing Server

OpenPak's Xbox 360 licensing server (Live service 21, `xbos`).

| | |
| --- | --- |
| Network | Xbox 360 |
| Role | Licensing and the certificate revocation list |
| Depends on | [`x360-common`](https://github.com/openpak/x360-common) (vendored) |
| Image | `ghcr.io/openpak/x360-xbos` |
| Used by | XAM through the security gateway |
| Licence | AGPL-3.0-only |

## What it serves

| Path | XRL | What |
| --- | --- | --- |
| `POST /xbos/GetRevocationList.ashx` | 0x918 | Revocation list: empty and unsigned |
| `GET /healthz` | | Health probe |

Identity headers (`X-X360-Account`, `X-X360-Users`) are trusted only with `X360_GATEWAY_KEY`. Split out of
[`x360-xuacs`](https://github.com/openpak/x360-xuacs), one repository per Live service.

## Run

```sh
GOWORK=off go run -mod=vendor ./cmd/x360-xbos
```


## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
