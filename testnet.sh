#!/bin/bash
rm -r ~/.bcflappcli
rm -r ~/.bcflappd

bcflappd init mynode --chain-id scavenge

bcflappcli config keyring-backend test
bcflappcli config chain-id scavenge
bcflappcli config output json
bcflappcli config indent true
bcflappcli config trust-node true

bcflappcli keys add user1
bcflappcli keys add user2
bcflappd add-genesis-account $(bcflappcli keys show user1 -a) 1000token,100000000stake
bcflappd add-genesis-account $(bcflappcli keys show user2 -a) 500token

bcflappd gentx --name user1 --keyring-backend test

bcflappd collect-gentxs 

echo -e "\nUse this command to start your test net."

echo -e "\nbcflappd start"
