# Norynth Chain
Norynth Chain – A Cosmos SDK-based blockchain for decentralized cognition and data integrity.

## Get Started

To get started with Norynth Chain, run the following command:

ignite chain serve

The `serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured using the `config.yml` file. For more information, refer to the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI offers scaffolding for both Vue and React frontends:

- For a Vue frontend, run: `ignite scaffold vue`
- For a React frontend, run: `ignite scaffold react`

These commands can be executed within your scaffolded blockchain project.

## Release

To release a new version of your blockchain, create and push a new tag with the `v` prefix. This will automatically create a draft release with the configured targets.

git tag v0.1 git push origin v0.1

Once the draft release is created, you can make any final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command:

curl https://get.ignite.com/username/norynth@latest! | sudo bash

Replace `username/norynth` with your actual `username` and `repo_name` from GitHub. For more details, refer to the [installation process](https://github.com/allinbits/starport-installer).

## Learn More

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
