## Fabric-go-gateway RESTful API

### QueryBlockByBlockHash

+ URL: x.x.x.x:8080/api/blockchain/block?channelId=XXX&&blockHash=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "number": 2,
      "previous_hash": "HSWWWFloxRQSVzKDm2ZsX3386tQlITSSmKbbjkdAPMg=",
      "data_hash": "EgwotM+ahuqdlgTxZ4D9R+j9ZbMtYXsOESmZUon6Ufg="
    }
  }
  ```

### QueryBlockByBlockNum

+ URL: x.x.x.x:8080/api/blockchain/block?channelId=XXX&&blockNum=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "number": 1,
      "previous_hash": "EZCiVEtuCXQGkYyOIxnRnISNKPDBF7jWFaw0JqG457g=",
      "data_hash": "h+tVNATnvhD9rfzTVSVUQB2wdlj7Y/spJSIbQGymFMc="
    }
  }
  ```

### QueryBlockByTxid

+ URL: x.x.x.x:8080/api/blockchain/block?channelId=XXX&&txId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "number": 4,
      "previous_hash": "Hy72LLyvR0w7eI6hOPc+SvsAR/xVOaYhC4clVv9ga6U=",
      "data_hash": "tx5yRCsZo8rGswrEpsFv0kVdY1yiCbV1wT0e26YEjMM="
    }
  }
  ```

### QueryTransactionByTxid

+ URL: x.x.x.x:8080/api/blockchain/transaction?channelId=XXX&&txId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "VALID"
  }
  ```

### QueryBlockChainInfo

+ URL: x.x.x.x:8080/api/blockchain/info?channelId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "height": 5006,
      "currentBlockHash": "nn0tsjPENnfj5Mx+4A9yHfJ67BaH2hcZ33ln4ezBdS4=",
      "previousBlockHash": "eA/BIp6PCixlOOzcT0ArdN1imubmbu3+Y+FOB1HJSY8="
    }
  }
  ```

### QueryBlockChainConfigBlock

+ URL: x.x.x.x:8080/api/blockchain/configblock?channelId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "number": 2,
      "previous_hash": "HSWWWFloxRQSVzKDm2ZsX3386tQlITSSmKbbjkdAPMg=",
      "data_hash": "EgwotM+ahuqdlgTxZ4D9R+j9ZbMtYXsOESmZUon6Ufg="
    }
  }
  ```

### QueryChannelPeers

+ URL: x.x.x.x:8080/api/channel/peers?channelId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": [
      "grpcs://192.168.9.96:7051"
    ]
  }
  ```

### QueryChannelAnchorPeers

+ URL: x.x.x.x:8080/api/channel/anchorpeers?channelId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": [
      {
        "Org": "Org2MSP",
        "Host": "peer0.org2.example.com",
        "Port": 9051
      },
      {
        "Org": "Org1MSP",
        "Host": "peer0.org1.example.com",
        "Port": 7051
      }
    ]
  }
  ```

### QueryChannelOrderers

+ URL: x.x.x.x:8080/api/channel/orderers?channelId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": [
      "orderer.example.com:7050"
    ]
  }
  ```

### QueryChaincodeEndorser

+ URL: x.x.x.x:8080/api/chaincode/endorser?channelId=XXX&&chaincodeId=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": [
      "grpcs://192.168.9.96:7051"
    ]
  }
  ```

### QueryPeerInstalledchaincode

+ URL: x.x.x.x:8080/api/peer/installedchaincode?peer=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "chaincodes": [
        {
          "name": "mycc",
          "version": "1.0",
          "path": "github.com/chaincode/chaincode_example02/go/",
          "id": "R2/KGpSSdAAZcfHsKDbLCTIfC3Emizdi1okxyT8hgTQ="
        },
        {
          "name": "mycc",
          "version": "1.0.1",
          "path": "github.com/hyperledger/fabric/fabric-samples/chaincode/chaincode_example02/go",
          "id": "LI9IrAdBPz46YWel7fpliXHHekN4kb1O6O60dMHDfwk="
        },
        {
          "name": "tscc",
          "version": "1.0.1",
          "path": "github.com/hyperledger/fabric/fabric-samples/chaincode/chaincode_example02/go",
          "id": "pkK2Qc1F7nQo/UdiFbcNpR7JKWYdDtR/kJ7ra1fARVA="
        }
      ]
    }
  }
  ```

### QueryPeerJoinedChannel

+ URL: x.x.x.x:8080/api/peer/joinedchannel?peer=XXX

+ Method: Get

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": {
      "channels": [
        {
          "channel_id": "mychannel"
        },
        {
          "channel_id": "test"
        }
      ]
    }
  }
  ```

### CreateChannel

+ URL: x.x.x.x:8080/api/channel/create

+ Method: Post

+ Request

  ```
  {
      "ChannelConfigPath":"/Users/eggsy/channel.tx", 
      "channelId":"test"
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "ok"
  }
  ```

### JoinChannel

+ URL: x.x.x.x:8080/api/channel/join

+ Method: Post

+ Request

  ```
  {
      "channelId":"test"
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "ok"
  }
  ```

### InstallChaincode

+ URL: x.x.x.x:8080/api/chaincode/install

+ Method: Post

+ Request

  ```
  {
  	"chaincodeVersion":"1.0.1",
  	"chaincodeId":"tscc",
  	"chaincodePath":"github.com/hyperledger/fabric/fabric-samples/chaincode/chaincode_example02/go"
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "ok"
  }
  ```

### InstantiateChaincode

+ URL: x.x.x.x:8080/api/chaincode/instantiate

+ Method: Post

+ Request

  ```
  {
  	"channelId":"mychannel",
  	"chaincodeVersion":"1.0.1",
  	"chaincodeId":"tscc",
  	"chaincodePath":"github.com/hyperledger/fabric/fabric-samples/chaincode/chaincode_example02/go",
  	"args":["init","a","500","b","500"],
  	"policy":"OR ('Org1MSP.peer','Org2MSP.peer')"
  	
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "ok"
  }
  ```

### QueryChaincode

+ URL: x.x.x.x:8080/api/chaincode/query

+ Method: Post

+ Request

  ```
  {
  	"channelId":"mychannel",
  	"chaincodeId":"tscc",
  	"fcn":"query",
  	"args":["a"]
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "500"
  }
  ```



### InvokeChaincode

+ URL: x.x.x.x:8080/api/chaincode/invoke

+ Method: Post

+ Request

  ```
  {
  	"channelId":"mychannel",
  	"chaincodeId":"mycc",
  	"fcn":"invoke",
  	"args":["a","b","1"]
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": null
  }
  ```

### UpgradeChaincode

+ URL: x.x.x.x:8080/api/chaincode/upgrade

+ Method: Post

+ Request

  ```
  {
  	"channelId":"mychannel",
  	"chaincodeVersion":"1.0.2",
  	"chaincodeId":"tscc",
  	"chaincodePath":"github.com/hyperledger/fabric/fabric-samples/chaincode/chaincode_example02/go",
  	"args":["init","a","500","b","500"],
  	"policy":"OR ('Org1MSP.peer','Org2MSP.peer')"
  	
  }
  ```

+ Response

  ```
  {
    "code": "200",
    "message": "Success",
    "data": "ok"
  }
  ```