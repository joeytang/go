package test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/thrift-iterator/go"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func Test_skip_list_of_string(t *testing.T) {
	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	proto := thrift.NewTBinaryProtocol(buf, true, true)
	proto.WriteListBegin(thrift.STRING, 3)
	proto.WriteString("a")
	proto.WriteString("b")
	proto.WriteString("c")
	proto.WriteListEnd()
	iter := thrifter.NewIterator(buf.Bytes())
	should.Equal(buf.Bytes(), iter.SkipList())
}

func Test_decode_list_of_string(t *testing.T) {
	should := require.New(t)
	buf := thrift.NewTMemoryBuffer()
	proto := thrift.NewTBinaryProtocol(buf, true, true)
	proto.WriteListBegin(thrift.STRING, 3)
	proto.WriteString("a")
	proto.WriteString("b")
	proto.WriteString("c")
	proto.WriteListEnd()
	iter := thrifter.NewIterator(buf.Bytes())
	should.Equal([]interface{}{"a", "b", "c"}, iter.ReadList())
}
