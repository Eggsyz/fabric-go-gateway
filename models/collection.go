package models

/**
 * @Author: eggsy
 * @Description:
 * @File:  collection
 * @Version: 1.0.0
 * @Date: 2020-03-27 17:38
 */
type CollectionConfig struct {
	Name              string `json:"name"`
	Policy            string `json:"policy"`
	RequiredPeerCount int32  `json:"requiredPeerCount"`
	MaximumPeerCount  int32  `json:"maximumPeerCount"`
	BlockToLive       uint64 `json:"blockToLive"`
	MemberOnlyRead    bool   `json:"memberOnlyRead"`
	MemberOnlyWrite   bool   `json:"memberOnlyWrite"`
}
