// package types 请求和响应体的定义

// Filename: ping.go
// Description: ping响应体

package types

type PingRequest struct {
}

type PingResponse struct {
	Version string `json:"version"` // 当前版本
}
