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
 * @File:  peer
 * @Version: 1.0.0
 * @Date: 12/8/19 3:44 下午
 */

type PeerController struct {
}

func (c *PeerController) GetJoinedchannel(ctx iris.Context) *models.Message {
	peer := ctx.FormValue("peer")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	response, err := client.QueryChannels(peer)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response)
}

func (c *PeerController) GetInstalledchaincode(ctx iris.Context) *models.Message {
	peer := ctx.FormValue("peer")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	response, err := client.QueryInstalledChaincodes(peer)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(response)
}
