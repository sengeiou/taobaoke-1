package service

import "context"

type Context = context.Context

const (
	itemInfoGetCode            = 11 // 淘宝客商品详情查询（简版）
	tbkTpwdCreateCode          = 12 //淘宝客-公用-淘口令生成
	tbkDgMaterialOptionalCode  = 13 //淘宝客-推广者-物料搜索
	tbkOrderDetailsGetCode     = 14 // 淘宝客-推广者-所有订单查询
	tbkScInvitecodeGetCode     = 15 // 淘宝客邀请码生成-社交
	tbkScPublisherInfoSaveCode = 16 //淘宝客信息备案
	tbkScPublisherInfoGetCode  = 17 // 淘宝客信息查询
)
