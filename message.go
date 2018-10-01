package main

type MsgRequestLogin struct {
	Api string `json:"api"`
	Uid int64  `json:"uid"`
}
