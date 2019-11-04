package imstrmap

import (
	"fmt"
	"github.com/onsi/gomega"
	"runtime"
	"runtime/debug"
	"testing"
)

func newSrcMap() map[string]string {

	m := map[string]string{
		"locality": "vlocality",
	}
	for _, k := range []string{"a", "b", "c", "das", "huhqw", "uyoqw", "y9qw", "juioq", "qqeq", "vqrqasas", "hqw", "asdqw", "asqwqwe"} {
		m[k] = k
	}
	return m
}

var (
	indexerFactory = NewIndexerFactory([]string{"locality"})
)

func TestStrMap(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	im := FromMap(newSrcMap(), indexerFactory)
	v, ok := im.Get("locality")
	g.Expect(v).To(gomega.Equal("vlocality"))
	g.Expect(ok).To(gomega.Equal(true))

	v, ok = im.Get("b")
	g.Expect(v).To(gomega.Equal("b"))
	g.Expect(ok).To(gomega.Equal(true))

	v, ok = im.Get("z")
	g.Expect(ok).To(gomega.Equal(false))
}

func BenchmarkNewImmutableMap(b *testing.B) {
	srcMap := newSrcMap()
	for i := 0; i < b.N; i++ {
		FromMap(srcMap, indexerFactory)
	}
}

func BenchmarkImmutabeStringMapIndexGet(b *testing.B) {
	srcMap := newSrcMap()
	m := FromMap(srcMap, indexerFactory)
	for i := 0; i < b.N; i++ {
		m.Get("locality")
		m.Get("locality")
		m.Get("locality")
	}
}

func BenchmarkImmutabeStringMapNoIndexGet(b *testing.B) {
	srcMap := newSrcMap()
	m := FromMap(srcMap, indexerFactory)
	for i := 0; i < b.N; i++ {
		m.Get("a")
		m.Get("b")
		m.Get("d")
	}
}

func BenchmarkStringMapGet(b *testing.B) {
	srcMap := newSrcMap()
	for i := 0; i < b.N; i++ {
		_ = srcMap["a"]
		_ = srcMap["b"]
		_ = srcMap["d"]
	}
}
func BenchmarkImmutabeStringMap_Range(b *testing.B) {

	srcMap := newSrcMap()
	m := FromMap(srcMap, indexerFactory)
	for i := 0; i < b.N; i++ {
		m.Range(func(s string, s2 string) bool {
			return false
		})
	}
}
func BenchmarkStringMap_Range(b *testing.B) {

	srcMap := newSrcMap()
	for i := 0; i < b.N; i++ {
		for k, v := range srcMap {
			_, _ = k, v
		}
	}
}

func BenchmarkImmutabeStringMap_Map(b *testing.B) {
	srcMap := newSrcMap()
	m := FromMap(srcMap, indexerFactory)
	for i := 0; i < b.N; i++ {
		m.Map()
	}
}

func TestImmutabeStringMap_Memory(t *testing.T) {
	runtime.GC()
	debug.FreeOSMemory()
	before := &runtime.MemStats{}
	runtime.ReadMemStats(before)
	a := make([]*ImmutabeStringMap, 0)
	for i := 0; i < 10000; i++ {
		a = append(a, FromMap(newSrcMap(), indexerFactory))
	}
	runtime.GC()
	debug.FreeOSMemory()
	after := &runtime.MemStats{}
	runtime.ReadMemStats(after)
	fmt.Println("heap", after.HeapInuse-before.HeapInuse)
	fmt.Printf("%p\n", a)
}

func TestStringMap_Memory(t *testing.T) {
	runtime.GC()
	debug.FreeOSMemory()
	before := &runtime.MemStats{}
	runtime.ReadMemStats(before)
	a := make([]map[string]string, 0)
	for i := 0; i < 10000; i++ {
		a = append(a, newSrcMap())
	}
	runtime.GC()
	debug.FreeOSMemory()
	after := &runtime.MemStats{}
	runtime.ReadMemStats(after)
	fmt.Println("heap", after.HeapInuse-before.HeapInuse)
	fmt.Printf("%p\n", a)
}
