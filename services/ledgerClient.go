package services

import (
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  ledgerClient
 * @Version: 1.0.0
 * @Date: 2019-10-23 15:47
 */
type LedgerClient struct {
	*ledger.Client
}

func (c *LedgerClient) QueryBlockByNum(blockNumber uint64) (*common.Block, error) {
	block, err := c.Client.QueryBlock(blockNumber)
	if err != nil {
		logger.Errorf("failed to query block. err: %v", err)
		return nil, err
	}
	return block, nil
}

func (c *LedgerClient) QueryBlockByHash(blockHash []byte) (*common.Block, error) {
	block, err := c.Client.QueryBlockByHash(blockHash)
	if err != nil {
		logger.Errorf("failed to query block by hash. err: %v", err)
		return nil, err
	}
	return block, nil
}

func (c *LedgerClient) QueryBlockByTxID(txID string) (*common.Block, error) {
	block, err := c.Client.QueryBlockByTxID(fab.TransactionID(txID))
	if err != nil {
		logger.Errorf("failed to query block by transaction ID. err: %v", err)
		return nil, err
	}
	return block, nil
}

func (c *LedgerClient) QueryTransaction(txID string) (*peer.ProcessedTransaction, error) {
	transaction, err := c.Client.QueryTransaction(fab.TransactionID(txID))
	if err != nil {
		logger.Errorf("failed to query transaction by transaction ID. err: %v", err)
		return nil, err
	}
	return transaction, nil
}

func (c *LedgerClient) QueryInfo() (*fab.BlockchainInfoResponse, error) {
	info, err := c.Client.QueryInfo()
	if err != nil {
		logger.Errorf("failed to query for blockchain info. err: %v", err)
		return nil, err
	}
	return info, nil
}

func (c *LedgerClient) QueryConfigBlock() (*common.Block, error) {
	configBlock, err := c.Client.QueryConfigBlock()
	if err != nil {
		logger.Errorf("failed to query configBlock. err: %v", err)
		return nil, err
	}
	return configBlock, nil
}
