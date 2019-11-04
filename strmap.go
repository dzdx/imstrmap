package imstrmap

import (
	"bytes"
	"encoding/binary"
)

func NewIndexerFactory(keys []string) func() *indexer {
	keyIndex := make(map[string]int, len(keys))
	for i, k := range keys {
		keyIndex[k] = i
	}
	return func() *indexer {
		offsets := make([]int, len(keys))
		for i := range offsets {
			offsets[i] = -1
		}
		return &indexer{
			keyIndex: keyIndex,
			offsets:  offsets,
		}
	}
}

type indexer struct {
	offsets  []int
	keyIndex map[string]int
}

func (i *indexer) getOffset(name string) int {
	index, ok := i.keyIndex[name]
	if !ok {
		return -1
	}
	return i.offsets[index]
}

func (i *indexer) setOffset(name string, offset int) {
	index, ok := i.keyIndex[name]
	if ok {
		i.offsets[index] = offset
	}
}

type ImmutabeStringMap struct {
	count   int
	indexer indexer
	data    []byte
}

func (m *ImmutabeStringMap) Get(k string) (string, bool) {
	offset := m.indexer.getOffset(k)
	if offset >= 0 {
		vLen := int(binary.BigEndian.Uint16(m.data[offset : offset+2]))
		return string(m.data[offset+2 : offset+2+vLen]), true
	}

	v := ""
	found := false
	kbytes := []byte(k)
	m.iter(func(ak []byte, av []byte) bool {
		if bytes.Equal(kbytes, ak) {
			v = string(av)
			found = true
			return true
		}
		return false
	})
	return v, found
}

func (m *ImmutabeStringMap) Range(f func(string, string) bool) {
	m.iter(func(s []byte, s2 []byte) bool {
		return f(string(s), string(s2))
	})
}

func (m *ImmutabeStringMap) Map() map[string]string {
	ret := make(map[string]string, m.count)
	m.iter(func(i []byte, i2 []byte) bool {
		ret[string(i)] = string(i2)
		return false
	})
	return ret
}

func (m *ImmutabeStringMap) iter(f func([]byte, []byte) bool) {
	offset := 0
	dataLen := len(m.data)
	for offset < dataLen {
		kLen := int(binary.BigEndian.Uint16(m.data[offset : offset+2]))
		k := m.data[offset+2 : offset+2+kLen]
		vLen := int(binary.BigEndian.Uint16(m.data[offset+2+kLen : offset+4+kLen]))
		v := m.data[offset+4+kLen : offset+4+kLen+vLen]
		if stop := f(k, v); stop {
			break
		}
		offset += 4 + kLen + vLen
	}
}

func FromMap(src map[string]string, indexerFactory func() *indexer) *ImmutabeStringMap {
	indexer := indexerFactory()
	var buf bytes.Buffer
	var offset int
	for k, v := range src {
		kLen := len(k)
		vLen := len(v)
		kLenBytes := make([]byte, 2)
		vLenBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(kLenBytes, uint16(kLen))
		binary.BigEndian.PutUint16(vLenBytes, uint16(vLen))

		buf.Write(kLenBytes)
		buf.WriteString(k)
		buf.Write(vLenBytes)
		buf.WriteString(v)
		indexer.setOffset(k, offset+2+kLen)
		offset += 4 + kLen + vLen
	}
	m := &ImmutabeStringMap{
		indexer: *indexer,
		data:    buf.Bytes(),
		count:   len(src),
	}
	return m
}
