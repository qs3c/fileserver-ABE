package mq

import (
	cmn "fileserver_enc/common"
)

// TransferData : 将要写到rabbitmq的数据的结构体
type TransferData struct {
	FileHash      string
	CurLocation   string
	DestLocation  string
	DestStoreType cmn.StoreType
}

//type TransferDataEnc struct {
//	FileHash           string
//	CurEncFileLocation string
//	CurAbeKeyLocation  string
//	DestLocation       string
//	DestStoreType      cmn.StoreType
//}
