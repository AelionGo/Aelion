// Filename: ping.go
// Description: ping响应体

// Package types 请求和响应体的定义
package types

type PingRequest struct {
}

type PingResponse struct {
	Version string `json:"version"` // 当前版本
}
