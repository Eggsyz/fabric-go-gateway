package controllers

import (
	"encoding/base64"
	"github.com/fabric-go-gateway/config"
	"github.com/fabric-go-gateway/models"
	"github.com/fabric-go-gateway/services"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/kataras/iris/v12"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  ledger
 * @Version: 1.0.0
 * @Date: 12/7/19 1:03 下午
 */
type LedgerController struct {
}

// 根据区块信息
func (c *LedgerController) GetBlock(ctx iris.Context) *models.Message {
	var block *common.Block
	var err error
	channelId := ctx.URLParam("channelId")
	txId := ctx.URLParam("txId")
	blockNum := ctx.URLParamInt64Default("blockNum", 0)
	blockHash := ctx.URLParam("blockHash")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	ledgerClient, err := client.GetLedgerClient(channelId)
	if err != nil {
		return ReturnError(err)
	}
	if len(txId) != 0 {
		block, err = ledgerClient.QueryBlockByTxID(txId)
	} else if len(blockHash) != 0 {
		hash, err := base64.StdEncoding.DecodeString(blockHash)
		if err != nil {
			return ReturnError(err)
		}
		block, err = ledgerClient.QueryBlockByHash(hash)
	} else {
		block, err = ledgerClient.QueryBlockByNum(uint64(blockNum))
	}
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(block.Header)
}

// 查询交易信息
func (c *LedgerController) GetTransaction(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	txId := ctx.FormValue("txId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	ledgerClient, err := client.GetLedgerClient(channelId)
	if err != nil {
		return ReturnError(err)
	}
	tx, err := ledgerClient.QueryTransaction(txId)
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(peer.TxValidationCode_name[tx.ValidationCode])
}

// 查询通道信息
func (c *LedgerController) GetInfo(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	ledgerClient, err := client.GetLedgerClient(channelId)
	if err != nil {
		return ReturnError(err)
	}
	info, err := ledgerClient.QueryInfo()
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(info.BCI)
}

func (c *LedgerController) GetConfigblock(ctx iris.Context) *models.Message {
	channelId := ctx.FormValue("channelId")
	client, err := services.NewClient(config.Conf.SdkCfgPath, config.Conf.OrgName, config.Conf.UserName)
	if err != nil {
		return ReturnError(err)
	}
	ledgerClient, err := client.GetLedgerClient(channelId)
	if err != nil {
		return ReturnError(err)
	}
	block, err := ledgerClient.QueryConfigBlock()
	if err != nil {
		return ReturnError(err)
	}
	return ReturnOk(block.Header)
}
