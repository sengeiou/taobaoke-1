// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathTBKPing = "/demo.service.v1.TBK/Ping"
var PathTBKKeyConvert = "/demo.service.v1.TBK/KeyConvert"
var PathTBKWithDraw = "/demo.service.v1.TBK/WithDraw"

var PathWechatMatchedTemplateMsgSend = "/demo.service.v1.Wechat/MatchedTemplateMsgSend"
var PathWechatBalanceTemplateMsgSend = "/demo.service.v1.Wechat/BalanceTemplateMsgSend"

// TBKBMServer is the server API for TBK service.
type TBKBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	KeyConvert(ctx context.Context, req *KeyConvertReq) (resp *KeyConvertResp, err error)

	WithDraw(ctx context.Context, req *WithDrawReq) (resp *WithDrawResp, err error)
}

var TBKSvc TBKBMServer

func tBKPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := TBKSvc.Ping(c, p)
	c.JSON(resp, err)
}

func tBKKeyConvert(c *bm.Context) {
	p := new(KeyConvertReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := TBKSvc.KeyConvert(c, p)
	c.JSON(resp, err)
}

func tBKWithDraw(c *bm.Context) {
	p := new(WithDrawReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := TBKSvc.WithDraw(c, p)
	c.JSON(resp, err)
}

// RegisterTBKBMServer Register the blademaster route
func RegisterTBKBMServer(e *bm.Engine, server TBKBMServer) {
	TBKSvc = server
	e.GET("/demo.service.v1.TBK/Ping", tBKPing)
	e.GET("/demo.service.v1.TBK/KeyConvert", tBKKeyConvert)
	e.GET("/demo.service.v1.TBK/WithDraw", tBKWithDraw)
}

// WechatBMServer is the server API for Wechat service.
type WechatBMServer interface {
	MatchedTemplateMsgSend(ctx context.Context, req *MatchedTemplateMsgSendReq) (resp *google_protobuf1.Empty, err error)

	BalanceTemplateMsgSend(ctx context.Context, req *BalanceTemplateMsgSendReq) (resp *google_protobuf1.Empty, err error)
}

var WechatSvc WechatBMServer

func wechatMatchedTemplateMsgSend(c *bm.Context) {
	p := new(MatchedTemplateMsgSendReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WechatSvc.MatchedTemplateMsgSend(c, p)
	c.JSON(resp, err)
}

func wechatBalanceTemplateMsgSend(c *bm.Context) {
	p := new(BalanceTemplateMsgSendReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := WechatSvc.BalanceTemplateMsgSend(c, p)
	c.JSON(resp, err)
}

// RegisterWechatBMServer Register the blademaster route
func RegisterWechatBMServer(e *bm.Engine, server WechatBMServer) {
	WechatSvc = server
	e.GET("/demo.service.v1.Wechat/MatchedTemplateMsgSend", wechatMatchedTemplateMsgSend)
	e.GET("/demo.service.v1.Wechat/BalanceTemplateMsgSend", wechatBalanceTemplateMsgSend)
}
