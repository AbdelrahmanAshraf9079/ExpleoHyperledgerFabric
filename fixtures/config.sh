export FABRIC_CFG_PATH=$PWD

rm -rf artifacts/*
rm -rf crypto-config/

./bin/cryptogen generate --config=./crypto-config.yaml

./bin/configtxgen -profile OrdererGenesis -channelID Orderer-channel -outputBlock ./artifacts/genesis.block


./bin/configtxgen -profile MainChannel -outputCreateChannelTx ./artifacts/MainChannel.tx -channelID MainChannel


./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./artifacts/SalesMSPanchors.tx -channelID MainChannel -asOrg SalesMSP

./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./artifacts/ResourcingMSPanchors.tx -channelID MainChannel -asOrg ResourcingMSP

./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./artifacts/EngManagementMSPanchors.tx -channelID MainChannel -asOrg EngManagementMSP

./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./artifacts/UpManagementMSPanchors.tx -channelID MainChannel -asOrg UpManagementMSP
