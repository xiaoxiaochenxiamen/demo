package proto_pb

import (
	"bytes"
	proto "github.com/golang/protobuf/proto"
	"proto_go"
	"utils"
)

const (
	PacketHeadLen = 8
)

func ReadTestPBRequstMatch(buff []byte) (requst *testPB.RequestMatch, err error) {
	requst = &testPB.RequestMatch{}
	err = proto.Unmarshal(buff, requst)
	return
}

func write(data []byte) []byte {
	dataLen := int32(len(data))
	buffLen := utils.EncodeInt32(dataLen + int32(PacketHeadLen))
	packetId := utils.EncodeInt32(int32(testPB.MsgId_MATCH_RESULT))
	var buff bytes.Buffer
	buff.Write(buffLen)
	buff.Write(packetId)
	buff.Write(data)
	return buff.Bytes()
}

func WriteTestPBMatchResult(result *testPB.MatchResult) ([]byte, error) {
	data, err := proto.Marshal(result)
	if nil != err {
		return nil, err
	} else {
		return write(data), nil
	}
}

func WriteTestPBRequstMatch(userId int32, score int32) ([]byte, error) {
	requst := &testPB.RequestMatch{
		UId:   &userId,
		Score: &score,
	}
	data, err := proto.Marshal(requst)
	if nil != err {
		return nil, err
	} else {
		return write(data), nil
	}
}
