package model

type KdcProxyMessage struct {
	KerbMessage   []byte `asn1:"tag:0,explicit"`
	TargetDomain  string `asn1:"tag:1,optional"`
	DcLocatorHint int    `asn1:"tag:2,optional"`
}
