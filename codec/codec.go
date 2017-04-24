package codec

import (
	"encoding/asn1"
	"encoding/hex"
	"github.com/kblin/go-kkdcp/model"
	"log"
)

type KkdcpRequest struct {
	Message []byte
	Domain  string
}

func Decode(data []byte) (req KkdcpRequest, err error) {
	msg := model.KdcProxyMessage{}
	_, err = asn1.Unmarshal(data, &msg)
	if err != nil {
		log.Printf("%s", hex.Dump(data))
		log.Fatalf("%s", err)
		return KkdcpRequest{}, err
	}
	return KkdcpRequest{Message: msg.KerbMessage, Domain: msg.TargetDomain}, nil
}

func Encode(data []byte) (message []byte, err error) {
	msg := model.KdcProxyMessage{KerbMessage: data}
	message, err = asn1.Marshal(msg)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	return message, nil
}
