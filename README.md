# Hanzo S3 KES (Key Encryption Service)

**KES is a cloud-native distributed key management and encryption server designed to secure modern applications at scale.**

- Repo: [github.com/hanzos3/kes](https://github.com/hanzos3/kes)
- Server: [github.com/hanzoai/s3](https://github.com/hanzoai/s3)
- Domain: [s3.hanzo.ai](https://s3.hanzo.ai) / [hanzo.space](https://hanzo.space)

---

- [What is KES?](#what-is-kes)
- [Installation](#install)
- [Quick Start](#quick-start)
- [Documentation](#docs)

## What is KES?

KES (Key Encryption Service) is a distributed key management server that scales horizontally. It can either be run as an edge server close to applications -- reducing latency to and load on a central key management system (KMS) -- or as a central key management service. KES nodes are self-contained stateless instances that can be scaled up and down automatically.

<p align="center">
  <img src='.github/arch.png?sanitize=true' width='70%'>
</p>

## Install

The KES server and CLI is available as a single binary, container image, or can be built from source.

<details open="true"><summary><b><a name="docker">Docker</a></b></summary>

Pull the latest release via:
```
docker pull ghcr.io/hanzos3/kes
```
</details>

<details><summary><b><a name="binary-releases">Binary Releases</a></b></summary>

| OS      | ARCH    | Binary                                                                                       |
|:-------:|:-------:|:--------------------------------------------------------------------------------------------:|
| linux   | amd64   | [linux-amd64](https://github.com/hanzos3/kes/releases/latest/download/kes-linux-amd64)       |
| linux   | arm64   | [linux-arm64](https://github.com/hanzos3/kes/releases/latest/download/kes-linux-arm64)       |
| darwin  | arm64   | [darwin-arm64](https://github.com/hanzos3/kes/releases/latest/download/kes-darwin-arm64)     |
| windows | amd64   | [windows-amd64](https://github.com/hanzos3/kes/releases/latest/download/kes-windows-amd64.exe) |

Download the binary via `curl` but replace `<OS>` and `<ARCH>` with your operating system and CPU architecture.
```
curl -sSL --tlsv1.2 'https://github.com/hanzos3/kes/releases/latest/download/kes-<OS>-<ARCH>' -o ./kes
```
```
chmod +x ./kes
```

You can also verify the binary with [minisign](https://jedisct1.github.io/minisign/) by downloading the corresponding [`.minisig`](https://github.com/hanzos3/kes/releases/latest) signature file.
Run:
```
curl -sSL --tlsv1.2 'https://github.com/hanzos3/kes/releases/latest/download/kes-<OS>-<ARCH>.minisig' -o ./kes.minisig
```
```
minisign -Vm ./kes -P RWTx5Zr1tiHQLwG9keckT0c45M3AGeHD6IvimQHpyRywVWGbP1aVSGav
```
</details>

<details><summary><b><a name="build-from-source">Build from source</a></b></summary>

Download and install the binary via your Go toolchain:

```sh
go install github.com/minio/kes/cmd/kes@latest
```

> Note: The Go module path remains `github.com/minio/kes` for upstream compatibility.
</details>

## Quick Start

You can get started by setting up your own KES server in less than five minutes.

<details><summary><b>First steps</b></summary>

#### 1. Start a dev server
```sh
kes server --dev
```

#### 2. Configure CLI
Point the KES CLI to the KES server:
```sh
export KES_SERVER=https://127.0.0.1:7373
export KES_API_KEY=kes:v1:AD9E7FSYWrMD+VjhI6q545cYT9YOyFxZb7UnjEepYDRc
```

#### 3. Create a Key
Create a new root encryption key -- e.g. `my-key`.
```
kes key create my-key
```
> Note that creating a new key will fail with `key already exist` if it already exists.

#### 4. Generate a DEK
Derive a new data encryption key (DEK).
```sh
kes key dek my-key
```
The plaintext part of the DEK would be used by an application to encrypt some data.
The ciphertext part of the DEK would be stored alongside the encrypted data for future
decryption.

</details>

## Docs

- [Integration Guides](https://github.com/hanzos3/kes/wiki#supported-kms-targets)
- [Server API](https://github.com/hanzos3/kes/wiki/Server-API)
- [Go SDK](https://pkg.go.dev/github.com/minio/kes-go)

### Monitoring

KES servers provide an API endpoint `/v1/metrics` that observability tools, like [Prometheus](https://prometheus.io/), can scrape.

For a graphical Grafana dashboard refer to the following [example](examples/grafana/dashboard.json).

![](.github/grafana-dashboard.png)

## FAQs

<details><summary><b>I have received an <code>insufficient permissions</code> error</b></summary>

This means that you are using a KES identity that is not allowed to perform a specific operation, like creating or listing keys.

The KES [admin identity](https://github.com/hanzos3/kes/blob/master/server-config.yaml#L8)
can perform any general purpose API operation. You should never experience a `not authorized: insufficient permissions`
error when performing general purpose API operations using the admin identity.

In addition to the admin identity, KES supports a [policy-based](https://github.com/hanzos3/kes/blob/master/server-config.yaml#L77) access control model.
You will receive a `not authorized: insufficient permissions` error in the following two cases:
1. **You are using a KES identity that is not assigned to any policy. KES rejects requests issued by unknown identities.**

   This can be fixed by assigning a policy to the identity. Checkout the [examples](https://github.com/hanzos3/kes/blob/master/server-config.yaml#L79-L88).
2. **You are using a KES identity that is assigned to a policy but the policy either not allows or even denies the API call.**

   In this case, you have to grant the API permission in the policy assigned to the identity. Checkout the [list of APIs](https://github.com/hanzos3/kes/wiki/Server-API#api-overview).
   For example, when you want to create a key you should allow the `/v1/key/create/<key-name>`. The `<key-name>` can either be a
   specific key name, like `my-key-1` or a pattern allowing arbitrary key names, like `my-key*`.

   Also note that deny rules take precedence over allow rules. Hence, you have to make sure that any deny pattern does not
   accidentally matches your API request.

</details>

---

## License
Use of `KES` is governed by the AGPLv3 license that can be found in the [LICENSE](./LICENSE) file.

Copyright 2015-2026 Hanzo AI, Inc.
