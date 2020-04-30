package utils

import (
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  utils
 * @Version: 1.0.0
 * @Date: 12/3/19 4:59 下午
 */
// generate private data CollectionConfig
func NewCollectionConfig(colName, policy string, reqPeerCount, maxPeerCount int32, blockToLive uint64, memberOnlyRead, MemberOnlyWrite bool) (*cb.CollectionConfig, error) {
	p, err := cauthdsl.FromString(policy)
	if err != nil {
		return nil, err
	}
	cpc := &cb.CollectionPolicyConfig{
		Payload: &cb.CollectionPolicyConfig_SignaturePolicy{
			SignaturePolicy: p,
		},
	}
	return &cb.CollectionConfig{
		Payload: &cb.CollectionConfig_StaticCollectionConfig{
			StaticCollectionConfig: &cb.StaticCollectionConfig{
				Name:              colName,
				MemberOrgsPolicy:  cpc,
				RequiredPeerCount: reqPeerCount,
				MaximumPeerCount:  maxPeerCount,
				BlockToLive:       blockToLive,
				MemberOnlyRead:    memberOnlyRead,
				MemberOnlyWrite:   MemberOnlyWrite,
			},
		},
	}, nil
}

func ConvertArgs(args []string) [][]byte {
	var ccArgs [][]byte
	for i := 0; i < len(args); i++ {
		ccArgs = append(ccArgs, []byte(args[i]))
	}
	return ccArgs
}
