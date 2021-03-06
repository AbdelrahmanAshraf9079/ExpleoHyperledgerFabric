const { Gateway, Wallets, TxEventHandler, GatewayOptions, DefaultEventHandlerStrategies, TxEventHandlerFactory } = require('fabric-network');
const fs = require('fs');
const path = require("path")
const log4js = require('log4js');
const logger = log4js.getLogger('BasicNetwork');
const util = require('util')


const helper = require('./helper')

const invokeTransaction = async (channelName, chaincodeName, fcn, args, username, orgName) => {
    
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

        const connectOptions = {
            wallet, identity: username, discovery: { enabled: true, asLocalhost: true },
            eventHandlerOptions: {
                commitTimeout: 100,
                strategy: DefaultEventHandlerStrategies.NETWORK_SCOPE_ALLFORTX
            },
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, connectOptions);

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork(channelName);
        const contract = network.getContract(chaincodeName);

        let result
        let message;
        if (fcn === "createDocument") {
            result = await contract.submitTransaction(fcn, args[0], args[1], args[2], args[3], args[4], args[5], args[6]);
            console.log(result.toString())

            message = `Successfully added the document asset with key ${args[0]}`

        } else if (fcn === "modifiyDocument") {
            result = await contract.submitTransaction(fcn, args[0], args[2], args[3], args[1]);
            message = `Successfully modified document with key ${args[0]}`
        } else if (fcn == "modifiyOfferState") {
            result = await contract.submitTransaction(fcn, args[0], args[1]);
            message = `Successfully modified offer state with key ${args[0]}`
        } else if (fcn == "addComment") {
            result = await contract.submitTransaction(fcn, args[0], args[1],args[2]);
            message = `Successfully added comment to document state with key ${args[0]}`
        }
        else {
            return `Invocation require either createDocument or modifiyDocument or modifiyOfferState or addComment as function but got ${fcn}`
        }

        await gateway.disconnect();
        
        //result = JSON.parse(result.toString());

        let response = {
            message: message,
            //result
        }
        return response;


    } catch (error) {

        console.log(`Getting error: ${error}`)
        return error.message

    }
}

exports.invokeTransaction = invokeTransaction;