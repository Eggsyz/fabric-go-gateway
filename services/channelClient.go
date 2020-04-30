package services

import (
	"errors"
	"github.com/fabric-go-gateway/utils"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

type ChClient struct {
	*channel.Client
}

// 调用链码
func (c *ChClient) InvokeOrQuery(invoke bool, chaincodeId, fcn string, args []string, peerEndpoints []string, transientMap map[string][]byte) (*channel.Response, error) {
	req := channel.Request{
		ChaincodeID: chaincodeId,
		Fcn:         fcn,
	}
	if args == nil {
		return nil, errors.New("the args can't be nil")
	}
	req.Args = utils.ConvertArgs(args)
	if transientMap != nil {
		req.TransientMap = transientMap
	}
	var opts []channel.RequestOption
	if peerEndpoints != nil {
		opts = append(opts, channel.WithTargetEndpoints(peerEndpoints...))
	}
	var response channel.Response
	var err error
	if invoke {
		logger.Infof("start invoke chaincode. chaincodeId: %s", chaincodeId)
		response, err = c.Client.Execute(req, opts...)
	} else {
		logger.Infof("start query chaincode. chaincodeId: %s", chaincodeId)
		response, err = c.Client.Query(req, opts...)
	}
	if err != nil {
		logger.Errorf("failed to invoke chaincode, err: %v", err)
		return nil, err
	}
	return &response, nil
}
