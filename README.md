<!--
parent:
  order: false
-->

<div align="center">
  <h1> neura </h1>
</div>

<!-- TODO: add banner -->
<!-- ![banner](docs/ethermint.jpg) -->

neura is a scalable, high-throughput Proof-of-Stake blockchain that is fully compatible and interoperable with Ethereum. 
It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) which runs on top of [CometBFT](https://github.com/cometbft/cometbft) consensus engine.
Ethereum compatibility allows developers to build applications on neura using the existing Ethereum codebase and toolset,
without rewriting smart contracts that already work on Ethereum or other Ethereum-compatible networks.
Ethereum compatibility is done using modules built by [Tharsis](https://thars.is) for their [Evmos](https://evmos.org) network.

neura Network’s Shariah-compliant native currency is ISLM – [Fatwa]([https://islamiccoin.net/fatwa](https://islamiccoin.net/shariah)).

**Note**: Requires [Go 1.21+](https://golang.org/dl)

## Installation

**Note**: Make sure that you install `Go` (you can follow official guide [Install Go](https://go.dev/doc/install) and `GOPATH` is configured for source directory).

For prerequisites and detailed build instructions please read the [Installation](https://docs.neura.network/quickstart/installation.html) instructions. Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/avenbreaks/neurastone/releases).

## Quick Start

To learn how the neura works from a high-level perspective, go to the [Intro](https://docs.neura.network/learn) section from the documentation. You can also check the instructions to [Run a Node](https://docs.neura.network/user-guides/run-node/mainnet).

<!-- ## Community -->
<!-- ## Contributing -->
