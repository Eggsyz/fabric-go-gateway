package services

import (
	"github.com/fabric-go-gateway/logging"
	"github.com/fabric-go-gateway/models"
	"github.com/fabric-go-gateway/utils"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	promsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"os"
	"sync"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  fabclient
 * @Version: 1.0.0
 * @Date: 2020-03-27 15:04
 */

type Client struct {
	rcClient        *resmgmt.Client
	mspClient       *msp.Client
	chClientMap     map[string]*ChClient
	ledgerClientMap map[string]*LedgerClient
	sdk             *fabsdk.FabricSDK
	orgName         string
	userName        string
	mu              sync.Mutex
}

var client *Client
var mu sync.Mutex
var logger = logging.MustGetLogger("services")

func NewClient(sdkCfgPath, orgName, userName string) (*Client, error) {
	if client == nil {
		mu.Lock()
		defer mu.Unlock()
		if client == nil {
			sdk, err := fabsdk.New(config.FromFile(sdkCfgPath))
			if err != nil {
				logger.Errorf("failed to create sdk. err: %v", err)
				return nil, err

			}
			context := sdk.Context(fabsdk.WithUser(userName), fabsdk.WithOrg(orgName))
			if context == nil {
				logger.Errorf("failed to create sdk context. err: %v", err)
				return nil, err
			}
			rsClient, err := resmgmt.New(context)
			if err != nil {
				logger.Errorf("failed to create resmgmt client. err: %v", err)
				return nil, err
			}
			mspClient, err := msp.New(sdk.Context(), msp.WithOrg(orgName))
			if err != nil {
				logger.Errorf("failed to create msp client. err: %v", err)
				return nil, err
			}
			client = &Client{
				sdk:             sdk,
				mspClient:       mspClient,
				rcClient:        rsClient,
				userName:        userName,
				orgName:         orgName,
				chClientMap:     make(map[string]*ChClient),
				ledgerClientMap: make(map[string]*LedgerClient),
			}
		}
	}
	logger.Info("Successfully created client.")
	return client, nil
}

// 创建通道
func (client *Client) CreateChannel(channelID, channelConfigPath, ordererEndpoint string) error {
	adminIdentity, err := client.mspClient.GetSigningIdentity(client.userName)
	if err != nil {
		logger.Errorf("failed to get identity: %v", err)
		return err
	}
	channelReq := resmgmt.SaveChannelRequest{
		ChannelID:         channelID,
		ChannelConfigPath: channelConfigPath,
		SigningIdentities: []promsp.SigningIdentity{adminIdentity},
	}
	var opts []resmgmt.RequestOption
	if len(ordererEndpoint) != 0 {
		opts = append(opts, resmgmt.WithOrdererEndpoint(ordererEndpoint))
	}
	_, err = client.rcClient.SaveChannel(channelReq, opts...)
	if err != nil {
		logger.Errorf("failed to create channel. channelId: %s, err: %v", channelID, err)
		return err
	}
	return nil
}

// 加入通道
func (client *Client) JoinChannel(channelID, ordererEndpoint string, peerEndpoint string) error {
	var opts []resmgmt.RequestOption
	opts = append(opts, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if len(ordererEndpoint) != 0 {
		opts = append(opts, resmgmt.WithOrdererEndpoint(ordererEndpoint))
	}
	if len(peerEndpoint) != 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(peerEndpoint))
	}
	err := client.rcClient.JoinChannel(channelID, opts...)
	if err != nil {
		logger.Errorf("failed to join channel. channelId: %s, err: %v", channelID, err)
		return err
	}
	return nil
}

// 安装链码
func (client *Client) InstallChaincode(chaincodeId, chaincodeVersion, chaincodePath string, peerEndpoint string) error {
	goPath := os.Getenv("GOPATH")
	ccPkg, err := gopackager.NewCCPackage(chaincodePath, goPath)
	if err != nil {
		logger.Errorf("failed to create chaincode package. chaincodeId: %s, err: %v", chaincodeId, err)
		return err
	}
	req := resmgmt.InstallCCRequest{
		Name:    chaincodeId,
		Path:    chaincodePath,
		Version: chaincodeVersion,
		Package: ccPkg,
	}
	var opts []resmgmt.RequestOption
	if len(peerEndpoint) != 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(peerEndpoint))
	}
	_, err = client.rcClient.InstallCC(req, opts...)
	if err != nil {
		logger.Errorf("failed to install chaincode. chaincodeId: %s, err: %v", chaincodeId, err)
		return err
	}
	return nil
}

// 实例化/升级链码
func (client *Client) InstantiateOrUpgradeChaincode(upgrade bool, channelId, chaincodeId, chaincodeVersion, chaincodePath, policy string, collections []*models.CollectionConfig, args []string, peerEndpoint string, ordererEndpoint string) error {
	ccPolicy := cauthdsl.AcceptAllPolicy
	if policy != "" {
		var err error
		ccPolicy, err = cauthdsl.FromString(policy)
		if err != nil {
			logger.Errorf("failed to get endorserment policy, err: %v ", err)
			return err
		}
	}
	var collConfig = make([]*common.CollectionConfig, len(collections))
	if len(collections) != 0 {
		for i := 0; i < len(collections); i++ {
			var err error
			collConfig[i], err = utils.NewCollectionConfig(collections[i].Name, collections[i].Policy, collections[i].RequiredPeerCount,
				collections[i].MaximumPeerCount, collections[i].BlockToLive, collections[i].MemberOnlyRead, collections[i].MemberOnlyWrite)
			if err != nil {
				logger.Errorf("failed to create collection config, err: %v ", err)
				return err
			}
		}
	}
	var req interface{}
	if upgrade {
		req = resmgmt.UpgradeCCRequest{
			Name:       chaincodeId,
			Path:       chaincodePath,
			Version:    chaincodeVersion,
			Policy:     ccPolicy,
			Args:       utils.ConvertArgs(args),
			CollConfig: collConfig,
		}
	} else {
		req = resmgmt.InstantiateCCRequest{
			Name:       chaincodeId,
			Path:       chaincodePath,
			Version:    chaincodeVersion,
			Policy:     ccPolicy,
			Args:       utils.ConvertArgs(args),
			CollConfig: collConfig,
		}
	}

	var opts []resmgmt.RequestOption
	if len(ordererEndpoint) != 0 {
		opts = append(opts, resmgmt.WithOrdererEndpoint(ordererEndpoint))
	}
	if len(peerEndpoint) != 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(peerEndpoint))
	}
	if upgrade {
		_, err := client.rcClient.UpgradeCC(channelId, req.(resmgmt.UpgradeCCRequest), opts...)
		if err != nil {
			logger.Errorf("failed to upgrade chaincode. err : %v", err)
			return err
		}
	} else {
		_, err := client.rcClient.InstantiateCC(channelId, req.(resmgmt.InstantiateCCRequest), opts...)
		if err != nil {
			logger.Errorf("failed to instantiate chaincode. err : %v", err)
			return err
		}
	}
	return nil
}

// 查询peer加入的通道
func (client *Client) QueryChannels(peerEndpoint string) (*peer.ChannelQueryResponse, error) {
	var opts []resmgmt.RequestOption
	if len(peerEndpoint) != 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(peerEndpoint))
	}
	response, err := client.rcClient.QueryChannels(opts...)
	if err != nil {
		logger.Errorf("failed to query channels. err: %v", err)
		return nil, err
	}
	return response, nil
}

// 查询peer已经安装的链码
func (client *Client) QueryInstalledChaincodes(peerEndpoint string) (*peer.ChaincodeQueryResponse, error) {
	var opts []resmgmt.RequestOption
	if len(peerEndpoint) != 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(peerEndpoint))
	}
	response, err := client.rcClient.QueryInstalledChaincodes(opts...)
	if err != nil {
		logger.Errorf("failed to query InstalledChaincodes. err: %v", err)
		return nil, err
	}
	return response, nil
}

// 查询通道实例化的链码
func (client *Client) QueryInstantiatedChaincodes(channelId string, peerEndpoint string) (*peer.ChaincodeQueryResponse, error) {
	var opts []resmgmt.RequestOption
	if len(peerEndpoint) != 0 {
		opts = append(opts, resmgmt.WithTargetEndpoints(peerEndpoint))
	}
	response, err := client.rcClient.QueryInstantiatedChaincodes(channelId, opts...)
	if err != nil {
		logger.Errorf("failed to query InstantiatedChaincodes. err: %v", err)
		return nil, err
	}
	return response, nil
}

// 查询通道锚节点信息
func (client *Client) QueryAnchorPeersFromOrderer(channelId string) ([]*fab.OrgAnchorPeer, error) {
	channelCfg, err := client.rcClient.QueryConfigFromOrderer(channelId)
	if err != nil {
		logger.Errorf("failed to query channel anchorpeers. err: %v", err)
		return nil, err
	}
	anchorPeers := channelCfg.AnchorPeers()
	return anchorPeers, nil
}

// 查询通道排序节点信息
func (client *Client) QueryOrderersFromOrderer(channelId string) ([]string, error) {
	channelCfg, err := client.rcClient.QueryConfigFromOrderer(channelId)
	if err != nil {
		logger.Errorf("failed to query channel orderers. err: %v", err)
		return nil, err
	}
	orderers := channelCfg.Orderers()
	return orderers, nil
}

// 查询通道peer节点信息
func (client *Client) QueryPeers(channelId string) ([]string, error) {
	clientChannelContext := client.sdk.ChannelContext(channelId, fabsdk.WithUser(client.userName), fabsdk.WithOrg(client.orgName))
	channelContext, err := clientChannelContext()
	if err != nil {
		logger.Errorf("failed to get channelContext. err: %v", err)
		return nil, err
	}
	discover, err := channelContext.ChannelService().Discovery()
	if err != nil {
		logger.Errorf("failed to get discovery. err: %v", err)
		return nil, err
	}
	peers, err := discover.GetPeers()
	if err != nil {
		logger.Errorf("failed to get peer. err: %v", err)
		return nil, nil
	}
	var peerUrls []string
	for _, peer := range peers {
		peerUrls = append(peerUrls, peer.URL())
	}
	return peerUrls, nil
}

// 查询通道链码背书节点信息
func (client *Client) QueryEndorsers(channelId, chaincodeId string) ([]string, error) {
	clientChannelContext := client.sdk.ChannelContext(channelId, fabsdk.WithUser(client.userName), fabsdk.WithOrg(client.orgName))
	channelContext, err := clientChannelContext()
	if err != nil {
		logger.Errorf("failed to get channelContext. err: %v", err)
		return nil, err
	}
	selection, err := channelContext.ChannelService().Selection()
	if err != nil {
		logger.Errorf("failed to get discovery. err: %v", err)
		return nil, err
	}
	endorsers, err := selection.GetEndorsersForChaincode([]*fab.ChaincodeCall{&fab.ChaincodeCall{ID: chaincodeId}})
	if err != nil {
		logger.Errorf("failed to get endorsers. err: %v", err)
		return nil, err
	}
	var peerUrls []string
	for _, peer := range endorsers {
		peerUrls = append(peerUrls, peer.URL())
	}
	return peerUrls, nil
}

// 获取账本客户端，查询区块或者交易
func (client *Client) GetLedgerClient(channelId string) (*LedgerClient, error) {
	if _, ok := client.ledgerClientMap[channelId]; !ok {
		client.mu.Lock()
		defer client.mu.Unlock()
		if _, ok = client.ledgerClientMap[channelId]; !ok {
			// create ledgerClient
			clientChannelContext := client.sdk.ChannelContext(channelId, fabsdk.WithUser(client.userName), fabsdk.WithOrg(client.orgName))
			ledCli, err := ledger.New(clientChannelContext)
			if err != nil {
				logger.Errorf("failed to create ledger client. err: %v", err)
				return nil, err
			}
			client.ledgerClientMap[channelId] = &LedgerClient{ledCli}
			return client.ledgerClientMap[channelId], nil
		}
	}
	return client.ledgerClientMap[channelId], nil
}

// 获取通道客户端，查询/调用链码
func (client *Client) GetChannelClient(channelId string) (*ChClient, error) {
	if _, ok := client.chClientMap[channelId]; !ok {
		client.mu.Lock()
		defer client.mu.Unlock()
		if _, ok = client.ledgerClientMap[channelId]; !ok {
			// create ledgerClient
			clientChannelContext := client.sdk.ChannelContext(channelId, fabsdk.WithUser(client.userName), fabsdk.WithOrg(client.orgName))
			chCli, err := channel.New(clientChannelContext)
			if err != nil {
				logger.Errorf("failed to create channel client. err: %v", err)
				return nil, err
			}
			client.chClientMap[channelId] = &ChClient{chCli}
			return client.chClientMap[channelId], nil
		}
	}
	return client.chClientMap[channelId], nil
}
