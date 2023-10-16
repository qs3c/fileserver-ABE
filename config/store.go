package config

import (
	cmn "fileserver_enc/common"
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "/data/fileserver/"
	//TempLocalRootDir = "C:\\Users\\27499\\GolandProjects\\fileserver_enc\\temp\\"
	// TempPartRootDir : 分块文件在本地临时存储地址的路径
	TempPartRootDir = "/data/fileserver/"
	//TempPartRootDir = "C:\\Users\\27499\\GolandProjects\\fileserver_enc\\temp\\"
	// CephRootDir : Ceph的存储路径prefix
	CephRootDir = "/ceph"
	// OSSRootDir : OSS的存储路径prefix
	OSSRootDir = "abe-cloud/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = cmn.StoreOSS
)
