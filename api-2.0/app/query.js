const { Gateway, Wallets, } = require('fabric-network');
const fs = require('fs');
const path = require("path")
const log4js = require('log4js');
const logger = log4js.getLogger('BasicNetwork');
const util = require('util')


const helper = require('./helper')
const query = async (channelName, chaincodeName, args, fcn, username, orgName) => {

    let orgPath = `connection-${orgName}.json`;
    let orgCa = `ca.${orgName}.expleoFabric.com`;

    let ccpPath = path.resolve(__dirname, '..', 'config', orgPath);
    let ccpJSON = fs.readFileSync(ccpPath, 'utf8');
    let ccp = JSON.parse(ccpJSON);

    orgName = orgName.charAt(0).toUpperCase()+orgName.slice(1);
    let orgMSP =`${orgName}MSP`;
    let walletName = `wallet${orgName}`;


    try {

        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', 'config', orgPath);
        const ccpJSON = fs.readFileSync(ccpPath, 'utf8')
        const ccp = JSON.parse(ccpJSON);

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), walletName);
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        let identity = await wallet.get(username);
        if (!identity) {
            console.log(`An identity for the user ${username} does not exist in the wallet, so registering user`);
            await helper.getRegisteredUser(username, org_name, true)
            identity = await wallet.get(username);
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, {
            wallet, identity: username, discovery: { enabled: true, asLocalhost: true }
        });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);

        // Get the contract from the network.
        const contract = network.getContract(chaincodeName);
        if (fcn == "readDocument"){

        let result = await contract.evaluateTransaction(fcn, args[0]);

        result = JSON.parse(result.toString());
        return result
        
        } else if ( fcn == "getDocumentHistory"){

        let result = await contract.evaluateTransaction(fcn, args[0]);
        result = JSON.parse(result.toString());
        return result

        } else if ( fcn == "getAllDocuments" ){

        let result = await contract.evaluateTransaction(fcn);
        result = JSON.parse(result.toString());
        return result

        }

    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        return error.message

    }
}

exports.query = query