package models

/**
 * @Author: eggsy
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 12/7/19 10:43 上午
 */
type Message struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
