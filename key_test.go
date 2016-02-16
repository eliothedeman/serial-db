package serial

import (
	"testing"

	"github.com/eliothedeman/randutil"
)

func randKV() *KeyVal {
	k := []byte(randutil.AlphaString(randutil.IntRange(10, 100)))
	v := []byte(randutil.AlphaString(randutil.IntRange(10, 100)))
	return NewKeyval(k, v)
}

func randKVs(n int) []*KeyVal {
	kvs := make([]*KeyVal, n)
	for i := 0; i < len(kvs); i++ {
		kvs[i] = randKV()
	}
	return kvs
}

func TestKeyMarshalUnmarshal(t *testing.T) {
	kv := randKV()

	buff := kv.MarshalTable(nil)
	n := &KeyVal{}
	err := n.UnmarshalTable(buff)
	if err != nil {
		t.Error(err)
	}

	if !kv.Equals(n) {
		t.Fail()
	}
}

func TestKeyMarshalUnmarshalPreAlloc(t *testing.T) {
	kv := randKV()
	buff := make([]byte, 1000)

	buff = kv.MarshalTable(buff)
	n := &KeyVal{}
	err := n.UnmarshalTable(buff)
	if err != nil {
		t.Error(err)
	}

	if !kv.Equals(n) {
		t.Fail()
	}
}
func BenchmarkKeyValMarshalTable(b *testing.B) {

	// make 1000 kvs
	kvs := randKVs(1000)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		kvs[i%len(kvs)].MarshalTable(nil)
	}
}

func BenchmarkKeyValUnMarshalTable(b *testing.B) {

	// make 1000 kvs
	kvs := randKVs(1000)
	buffs := make([][]byte, len(kvs))
	for i, kv := range kvs {
		buffs[i] = kv.MarshalTable(nil)
	}

	kv := &KeyVal{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		kv.UnmarshalTable(buffs[i%len(buffs)])
	}
}

func BenchmarkKeyValMarshalTablePreAlloc(b *testing.B) {

	// make 1000 kvs
	kvs := randKVs(1000)

	b.ResetTimer()
	b.ReportAllocs()
	buff := make([]byte, 1000)
	for i := 0; i < b.N; i++ {
		buff = kvs[i%len(kvs)].MarshalTable(buff)
	}
}

func BenchmarkKeyValUnMarshalTablePreAlloc(b *testing.B) {

	// make 1000 kvs
	kvs := randKVs(1000)
	buffs := make([][]byte, len(kvs))
	buff := make([]byte, 1000)
	for i, kv := range kvs {
		buffs[i] = kv.MarshalTable(buff)
	}

	kv := &KeyVal{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		kv.UnmarshalTable(buffs[i%len(buffs)])
	}
}
