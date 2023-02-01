#!/usr/bin/env bash

# Remove existing chain data
rm -r ~/.meled

# make folder structure for upgrades
mkdir -p ~/.meled/upgrade_manager/genesis/bin
mkdir -p ~/.meled/upgrade_manager/upgrades

# symlink genesis binary to upgrade
ln -s $(which meled) ~/.meled/upgrade_manager/genesis/bin/meled

# Initialize validator
meled init validator --chain-id test

# Validator
# mele17zc58s96rxj79jtqqsnzt3wtx3tern6ad2rn0h
echo "pet apart myth reflect stuff force attract taste caught fit exact ice slide sheriff state since unusual gaze practice course mesh magnet ozone purchase" | meled keys add validator --keyring-backend test --recover

# Validator1
# mele14u53eghrurpeyx5cm47vm3qwugtmhcpn7ckqm5
echo "bottom soccer blue sniff use improve rough use amateur senior transfer quarter" | meled keys add validator1 --keyring-backend test --recover

# Validator2
# mele1y95lvkndyxd99keazuzudq5xyzuvhvwsz7sms9
echo "wreck layer draw very fame person frown essence approve lyrics sustain spoon" | meled keys add validator2 --keyring-backend test --recover

# Validator3
# mele1cr00glwjvf7fzv72tfdj4x3ctusq6uu3v674zu
echo "exotic merit wrestle sad bundle age purity ability collect immense place tone" | meled keys add validator3 --keyring-backend test --recover

# Validator4
# mele1nknm73uvfjwslmnwayh7n2c0fv9vwujnplwmcy
echo "faculty head please solid picnic benefit hurt gloom flag transfer thrive zebra" | meled keys add validator4 --keyring-backend test --recover

# Validator5
# mele1jsll0r65zynv86vnt7ryhtmr6eqh98zl9yzcr3
echo "amateur napkin price catch burger void bid more inner retire club cram" | meled keys add validator5 --keyring-backend test --recover

# Test 1
# mele1dfjns5lk748pzrd79z4zp9k22mrchm2a5t2f6u
echo "betray theory cargo way left cricket doll room donkey wire reunion fall left surprise hamster corn village happy bulb token artist twelve whisper expire" | meled keys add test1 --keyring-backend test --recover

# Test 2
# mele1c7nn5mt43m37t0zmqwh6rslrgcr3gd4pxqutpj
echo "toss sense candy point cost rookie jealous snow ankle electric sauce forward oblige tourist stairs horror grunt tenant afford master violin final genre reason" | meled keys add test2 --keyring-backend test --recover

# Operator 1
# mele1zsnyxrqkdt8e8sa8kjug8ftemn2qvu5x79hh9y
echo "find diamond example tooth need impact document total enrich hobby axis bicycle more oak junk because blade alley mesh electric evolve duty attack once" | meled keys add operator1 --keyring-backend test --recover

# Operator 2
# mele1sys3te2w3n23v80smprt9whuncjl4scytfy3af
echo "van hungry victory version major maple scan era buddy exact scheme again mention plastic clutch motor aware easily early zone tiger flavor shell bright" | meled keys add operator2 --keyring-backend test --recover

# Manager
# mele12rmu8657nunpgnd5ufwqphnwlspzcwl29ejqpa
echo "lawn cup spawn stay amazing stuff marble egg north measure survey until divorce ridge hat whip okay home solar brave soft nut kitchen lady" | meled keys add manager --keyring-backend test --recover

# Delegator
# mele1e2k26e6yz3kwysp896mwpf5u5r0nh8zx0yvtsz
echo "around fire birth cradle assault equal risk dune goat recycle torch hole control pluck cry math noble crystal language uncover leave ski dust answer" | meled keys add delegator --keyring-backend test --recover

# Add genesis accounts
meled add-genesis-account $(meled keys show validator -a --keyring-backend test) 100000000000000umelg,300000000000000umelc
meled add-genesis-account $(meled keys show validator1 -a --keyring-backend test) 110000000000000umelg,300000000000000umelc
meled add-genesis-account $(meled keys show validator2 -a --keyring-backend test) 120000000000000umelg,300000000000000umelc
meled add-genesis-account $(meled keys show validator3 -a --keyring-backend test) 130000000000000umelg,300000000000000umelc
meled add-genesis-account $(meled keys show validator4 -a --keyring-backend test) 140000000000000umelg,300000000000000umelc
meled add-genesis-account $(meled keys show validator5 -a --keyring-backend test) 150000000000000umelg,300000000000000umelc
meled add-genesis-account $(meled keys show test1 -a --keyring-backend test) 10000000000000umelg,20000000000000umelc
meled add-genesis-account $(meled keys show test2 -a --keyring-backend test) 10000000000000umelg,20000000000000umelc
meled add-genesis-account $(meled keys show delegator -a --keyring-backend test) 10000000000000umelg,20000000000000umelc
meled add-genesis-account $(meled keys show operator1 -a --keyring-backend test) 20000000000000umelc
meled add-genesis-account $(meled keys show operator2 -a --keyring-backend test) 20000000000000umelc
meled add-genesis-account $(meled keys show manager -a --keyring-backend test) 20000000000000umelc

# Generate CreateValidator signed transaction
meled gentx validator 90000000000000umelg --keyring-backend test --chain-id test

# Collect genesis transactions
meled collect-gentxs

# Fix genesis
sed -i 's/stake/umelg/g' ~/.meled/config/genesis.json