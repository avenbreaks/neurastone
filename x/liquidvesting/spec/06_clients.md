<!--
order: 6
-->

# Clients

## CLI

Find below a list of  `neurad` commands added with the  `x/liquidvesting` module. You can obtain the full list by using the `neurad -h` command. A CLI command can look like this:

```bash
neurad query liquidvesting params
```

### Queries

| Command                 | Subcommand | Description                    |
|-------------------------|------------|--------------------------------|
| `query` `liquidvesting` | `denom`    | Get liquid token               |
| `query` `liquidvesting` | `denoms`   | Get all existing liquid tokens |

### Transactions

| Command              | Subcommand  | Description                                       |
|----------------------|-------------|---------------------------------------------------|
| `tx` `liquidvesting` | `liquidate` | Liquidates arbitrary amount of locked ISLM tokens |
| `tx` `liquidvesting` | `redeem`    | Redeem liquid token to ISLM                       |

## gRPC

### Queries

| Verb   | Method                              | Description                    |
| ------ |-------------------------------------| ------------------------------ |
| `gRPC` | `neura.liquidvesting.v1.Query/Denom` | Get liquid token               |
| `gRPC` | `neura.liquidvesting.v1.Query/Denoms`| Get all existing liquid tokens |
| `GET`  | `/neura/liquidvesting/v1/denom`      | Get liquid token               |
| `GET`  | `/neura/liquidvesting/v1/denom`      | Get all existing liquid tokens |

### Transactions

| Verb   | Method                                | Description                                       |
|--------|---------------------------------------| ------------------------------------------------- |
| `gRPC` | `neura.liquidvesting.v1.Msg/Liquidate` | Liquidates arbitrary amount of locked ISLM tokens |
| `gRPC` | `neura.liquidvesting.v1.Msg/Redeem`    | Redeem liquid token to ISLM                       |
| `POST` | `/neura/liquidvesting/v1/tx/liquidate` | Liquidates arbitrary amount of locked ISLM tokens |
| `POST` | `/neura/liquidvesting/v1/tx/redeem`    | Redeem liquid token to ISLM                       |
