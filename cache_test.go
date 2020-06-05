package gocache

import (
	"github.com/muesli/cache2go"
	. "github.com/smartystreets/goconvey/convey"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestCache_Add(t *testing.T) {
	Convey("test add", t, func() {
		DefaultCache.Add("1", "11", time.Duration(5)*time.Second)
		v, exist := DefaultCache.Get("1")
		So(exist, ShouldBeTrue)
		So(v.(string), ShouldEqual, "11")
		time.Sleep(6 * time.Second)
		_, exist = DefaultCache.Get("1")
		So(exist, ShouldBeFalse)
	})
}

//BenchmarkCache_Add-4   	 5775886	       222 ns/op
func BenchmarkCache_Add(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.RunParallel(func(pb *testing.PB) {
		var i int
		for pb.Next() {
			DefaultCache.Add(strconv.Itoa(i), "111", time.Duration(5)*time.Second)
			DefaultCache.Get(strconv.Itoa(i))
		}
		i++
	})
}

//BenchmarkCache2go-4   	  887784	      1549 ns/op
func BenchmarkCache2go(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	table := cache2go.Cache("test")
	b.RunParallel(func(pb *testing.PB) {
		var i int
		for pb.Next() {
			table.Add(i, time.Duration(5)*time.Second, i)
			table.Value(i)
			i++
		}
	})
}
