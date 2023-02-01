# Mele Network Upgrade Guide

The network upgrade procedure consists of a coordinated set of steps that replaces the binary versions across nodes as well as execute the upgrade handler on each node’s state.

## Upgrade Preparation
First of all, it is important to agree on a commit hash of the codebase that will be the new version.

Then the linux binaries need to be built running `make build-linux`. The output of that command will be the new `meled` binary located in the `build` folder. This will be the new node binary. This binary needs to be put into a zip file called `meled-upgrade-linux-amd64.zip`.

By running `sha256sum meled-upgrade-linux-amd64.zip`, we will get the sha256 sum of the linux binary that will be included in the governance proposal.

Repeat steps 2-4 for any distribution you wish to automatically upgrade. (For example `windows/amd64` or `darwin/amd64`). The full list of platforms and architectures can be found here https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63 and the build command is the following
`GOOS=linux GOARCH=amd64 make build`

These zip files need to be uploaded to a publicly accessible place. This can either be AWS S3 or a simple EC2 instance with enough data transfer. It is important to keep in mind that all nodes will contact this server simultaneously. The test for this can be if that file can be wget and the sha256 sum matches the generation one.
We have simulated the upgrade on a VPS by installing apache and putting all the zip files in `/var/www/html` and were accessible via http://IP/meled-upgrade-linux.zip.

## Issuing a Proposal

An upgrade proposal JSON file needs to be created in the following form:

```
{
    "title": "Software Upgrade Proposal",
    "description": "Change node software",
    "deposit": [
        {
            "denom": "umelg",
            "amount": "10000000",
        },
    ],
    "plan": {
        "name": "upgrade",
        "height": "20",
        "info": "{\"binaries\":{\"darwin/amd64\":\"http://IP/meled-upgrade-darwin-amd64.zip?checksum=sha256:801a192e1d3c0c64bbf1f85b185e45322f01131f659c6264cb4202e7f52a29a6\",\"linux/amd64\":\"http://207.154.244.153/meled-upgrade-linux-amd64.zip?checksum=sha256:f0c2355397d5352b0bc72fa57b3620d97cb5672e0e3e1e3b09b56f77efda2884\"}}"
    }
}
```

It is very **VERY IMPORTANT** to note that the info field of the governance proposal shouldn’t have any spaces or tabs and that every quote needs to be escaped (\").
The title and the description can be arbitrary values, name should be fixed, and the height needs to be calculated to be sufficiently in the future so that all validators would be able to vote, and all validators that don’t plan to use the automatic upgrade feature have the time to upgrade their nodes.

The proposal can be issued with the following command:

```
meled tx mgov submit-proposal software-upgrade ./proposal_upgrade.json --from <WALLET_NAME> --keyring-backend test --chain-id <CHAIN_ID> -b block -y
```

After that, validators are able to vote for that proposal by querying all proposals, finding the ID and voting for that ID

```
meled q mgov proposals
```

## Voting

```
meled tx mgov vote <PROPOSAL_ID> yes --from <VALIDATOR_WALLET> --keyring-backend test --chain-id <CHAIN_ID> -b block -y
```

## Automatic Upgrade

All validators that have setup their validator software being run through `meled-manager` and have the environment variable `DAEMON_ALLOW_DOWNLOAD_BINARIES` and `DAEMON_RESTART_AFTER_UPGRADE` both set to `"on"` don’t have to take any further steps. When the scheduled upgrade block comes, the upgrade manager will, based on their OS and architecture, download the appropriate binaries and restart the chain with them and execute the upgrade handler.

## Manual Upgrade
For validators and nodes that do not want to enable automatic upgrades they need to follow the following steps:

- Download (wget) or build the binary for their distribution 
- Create the upgrade structure
```
mkdir -p  ~/.meled/upgrade_manager/upgrades/upgrade/bin
```
- Move the binary to the upgrade file structure
```
cp build/meled ~/.meled/upgrade_manager/upgrades/upgrade/bin
```

Note:

This upgrade procedure applies to the `meled-manager` that is being run by the process manager. If one wants to receive a new version of meled that they can invoke from their shell they will need to manually acquire the meled binary from predefined URLs or build it locally with the `make install` command.

