package controllers

import (
	"github.com/fabric-go-gateway/config"
	"github.com/fabric-go-gateway/models"
	"github.com/fabric-go-gateway/services"
	"github.com/kataras/iris/v12"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  chaincode
 * @Version: 1.0.0
 * @Date: 12/7/19 1:01 下午
 */
type ChaincodeController struct {
}

// 安装链码
func (c *ChaincodeController) PostInstall(ctx iris.Context) *models.Message {
	var request models.InstallChaincodeRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	err = client.InstallChaincode(request.ChaincodeId, request.ChaincodeVersion, request.ChaincodePath, request.PeerEndpoint)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk("ok")
}

// 实例化链码
func (c *ChaincodeController) PostInstantiate(ctx iris.Context) *models.Message {
	var request models.InstantiateOrUpgradeChaincodeRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	err = client.InstantiateOrUpgradeChaincode(false, request.ChannelId, request.ChaincodeId, request.ChaincodeVersion, request.ChaincodePath, request.Policy, request.CollectionConfigs, request.Args, request.PeerEndpoint, request.OrdererEndpoint)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk("ok")
}

// 查询链码
func (c *ChaincodeController) PostQuery(ctx iris.Context) *models.Message {
	var request models.QueryOrInvokeChaincodeRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	channelClient, err := client.GetChannelClient(request.ChannelId)
	if err != nil {
		return ReturnError(err)
	}
	response, err := channelClient.InvokeOrQuery(false, request.ChaincodeId, request.Fcn, request.Args, request.PeerEndpoints, nil)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(string(response.Payload))
}

// 调用链码
func (c *ChaincodeController) PostInvoke(ctx iris.Context) *models.Message {
	var request models.QueryOrInvokeChaincodeRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	channelClient, err := client.GetChannelClient(request.ChannelId)
	if err != nil {
		return ReturnError(err)
	}
	response, err := channelClient.InvokeOrQuery(true, request.ChaincodeId, request.Fcn, request.Args, request.PeerEndpoints, request.TransientMap)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response.Payload)
}

// 升级链码
func (c *ChaincodeController) PostUpgrade(ctx iris.Context) *models.Message {
	var request models.InstantiateOrUpgradeChaincodeRequest
	err := ctx.ReadJSON(&request)
	if err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	err = client.InstantiateOrUpgradeChaincode(true, request.ChannelId, request.ChaincodeId, request.ChaincodeVersion, request.ChaincodePath, request.Policy, request.CollectionConfigs, request.Args, request.PeerEndpoint, request.OrdererEndpoint)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk("ok")
}

// 查询链码背书节点信息
func (c *ChaincodeController) GetEndorser(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	chaincodeId := ctx.FormValue("chaincodeId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	res, err := client.QueryEndorsers(channelId, chaincodeId)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(res)
}
