export FABRIC_CFG_PATH=$PWD

rm -rf artifacts/*
rm -rf crypto-config/

docker rm -f $(docker ps -aq)

./bin/cryptogen generate --config=./crypto-config.yaml

./bin/configtxgen -profile OrdererGenesis -channelID mainchannel -outputBlock ./channel-artifacts/genesis.block


./bin/configtxgen -profile MainChannel -outputCreateChannelTx ./channel-artifacts/MainChannel.tx -channelID mainchannel


./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/SalesMSPanchors.tx -channelID mainchannel -asOrg SalesMSP

./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/ResourcingMSPanchors.tx -channelID mainchannel -asOrg ResourcingMSP

./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/EngManagementMSPanchors.tx -channelID mainchannel -asOrg EngManagementMSP

./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/UpManagementMSPanchors.tx -channelID mainchannel -asOrg UpManagementMSP



export IMAGE_TAG=latest

docker-compose -f docker-compose-cli.yaml -f docker-compose-etcdraft2.yaml up -d

# docker-compose -f docker-compose-cli.yaml -f docker-compose-couch.yaml -f docker-compose-etcdraft2.yaml up -d

docker ps -a

docker exec -it cli bash


