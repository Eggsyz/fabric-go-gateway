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
 * @File:  hyperController
 * @Version: 1.0.0
 * @Date: 12/7/19 10:07 上午
 */

type ChannelController struct {
}

// 创建通道
func (c *ChannelController) PostCreate(ctx iris.Context) *models.Message {
	var request models.CreateChannelRequest
	if err := ctx.ReadJSON(&request); err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	err = client.CreateChannel(request.ChannelId, request.ChannelConfigPath, request.OrdererEndpoint)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk("ok")
}

// 加入通道
func (c *ChannelController) PostJoin(ctx iris.Context) *models.Message {
	var request models.JoinChannelRequest
	if err := ctx.ReadJSON(&request); err != nil {
		return ReturnError(err)
	}
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	err = client.JoinChannel(request.ChannelId, request.OrdererEndpoint, request.PeerEndpoint)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk("ok")
}

// 查询通道锚节点信息
func (c *ChannelController) GetAnchorpeer(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	response, err := client.QueryAnchorPeersFromOrderer(channelId)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response)
}

// 查询通道peer节点信息
func (c *ChannelController) GetPeer(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	response, err := client.QueryPeers(channelId)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response)
}

// 查询通道排序节点信息
func (c *ChannelController) GetOrderer(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	response, err := client.QueryOrderersFromOrderer(channelId)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response)
}

// 查询通道实例化链码信息
func (c *ChannelController) GetInstantiatedchaincodes(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	response, err := client.QueryInstantiatedChaincodes(channelId, "")
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response)
}
