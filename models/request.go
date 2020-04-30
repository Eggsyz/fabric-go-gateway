package models

/**
 * @Author: eggsy
 * @Description:
 * @File:  request
 * @Version: 1.0.0
 * @Date: 12/7/19 10:06 上午
 */

type CreateChannelRequest struct {
	ChannelId         string `json:"channelId"`
	ChannelConfigPath string `json:"channelConfigPath"`
	OrdererEndpoint   string `json:"ordererEndpoint"`
}

type JoinChannelRequest struct {
	ChannelId       string `json:"channelId"`
	OrdererEndpoint string `json:"ordererEndpoint"`
	PeerEndpoint    string `json:"peerEndpoint"`
}

type InstallChaincodeRequest struct {
	ChaincodeId      string `json:"chaincodeId"`
	ChaincodeVersion string `json:"chaincodeVersion"`
	ChaincodePath    string `json:"chaincodePath"`
	PeerEndpoint     string `json:"peerEndpoint"`
}

type QueryOrInvokeChaincodeRequest struct {
	ChaincodeId   string            `json:"chaincodeId"`
	ChannelId     string            `json:"channelId"`
	Fcn           string            `json:"fcn"`
	Args          []string          `json:"args"`
	TransientMap  map[string][]byte `json:"transientMap"`
	PeerEndpoints []string          `json:"peerEndpoints"`
}

type InstantiateOrUpgradeChaincodeRequest struct {
	ChannelId         string              `json:"channelId"`
	ChaincodeId       string              `json:"chaincodeId"`
	ChaincodeVersion  string              `json:"chaincodeVersion"`
	ChaincodePath     string              `json:"chaincodePath"`
	Args              []string            `json:"args"`
	Policy            string              `json:"policy"`
	CollectionConfigs []*CollectionConfig `json:"collectionConfigs"`
	PeerEndpoint      string              `json:"peerEndpoint"`
	OrdererEndpoint   string              `json:"ordererEndpoint"`
}

type QueryChaincodeRequest struct {
	ChannelId    string   `json:"channelId"`
	ChaincodeId  string   `json:"chaincodeId"`
	Args         []string `json:"args"`
	PeerEndpoint string   `json:"peerEndpoint"`
}

type InvokeChaincodeRequest struct {
	ChannelId     string            `json:"channelId"`
	ChaincodeId   string            `json:"chaincodeId"`
	Args          []string          `json:"args"`
	PeerEndpoints []string          `json:"peerEndpoints"`
	TransientMap  map[string][]byte `json:"transientMap"`
}

type QueryRequest struct {
	ChannelId string `json:"channelId"`
}

type QueryBlockRequest struct {
	ChannelId string `json:"channelId"`
	BlockNum  uint64 `json:"blockNum"`
	BlockHash []byte `json:"blockHash"`
	TxID      string `json:"txId"`
}

type QueryTransactionRequest struct {
	ChannelId string `json:"channelId"`
	TxID      string `json:"txId"`
}

type QueryEndorsersRequest struct {
	ChannelId   string `json:"channelId"`
	ChaincodeId string `json:"chaincodeId"`
}

type QueryChannelsRequest struct {
	PeerEndpoint string `json:"peerEndpoint"`
}
