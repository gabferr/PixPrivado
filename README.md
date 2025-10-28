# Pix Privado (Private Pix)

A high-performance, permissioned payment system inspired by Brazil's Pix, built on an EVM blockchain (Solidity) with a Go (Golang) backend.

This project was developed as a portfolio piece to demonstrate a complete, professional-grade backend and blockchain solution.

## Core Features

* Instant, internal B2B transfers.
* RESTful API for integration with ERPs and other systems.
* Immutable transaction ledger on a private EVM blockchain.
* Administrative panel for user and account management.

## Tech Stack

* **Backend:** Go (Golang)
* **Blockchain:** Solidity (EVM-compatible)
* **Database:** PostgreSQL
* **Cache:** Redis (for locking and idempotency)
* **API:** REST (using `go-ethereum` to interact with the chain)
* **Deployment:** Docker / Docker Compose

## Architecture

(Aqui você pode colocar uma versão simples do diagrama de arquitetura do seu PDF)

```mermaid
flowchart LR
    subgraph FRONTEND
        U[Users / ERPs]
    end
    API [API Gateway]
    subgraph BACKEND
        B[Backend Go]
        R[Redis]
        DB[(Postgres)]
    end
    subgraph CHAIN
        N1 [EVM Node 1]
        SC[Smart Contract (Solidity)]
    end
    U --> API --> B
    B --> DB
    B --> R
    B --> N1
    N1 --> SC
```

## How to Run

1.  Clone the repository:
    ```bash
    git clone [https://github.com/gabferr/PixPrivado.git](https://github.com/gabferr/PixPrivado.git)
    cd PixPrivado
    ```
2.  Start the environment:
    ```bash
    docker-compose up -d
    ```
