package lib

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func Atoi64(s string) int64 {
	integer, _ := strconv.Atoi(s)
	return int64(integer)
}

func Atoi32(s string) int32 {
	integer, _ := strconv.Atoi(s)
	return int32(integer)
}

func Atoi(s string) int {
	integer, _ := strconv.Atoi(s)
	return integer
}

func Itoa64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Itoa32(i int32) string {
	return strconv.Itoa(int(i))
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func ReadUInt16(data []byte) (ret uint16) {
	ret = binary.BigEndian.Uint16(data)
	return
}

// http://stackoverflow.com/questions/16888357/convert-an-integer-to-a-byte-array
func ReadInt32(data []byte) (ret int32) {
	ret = int32(binary.BigEndian.Uint32(data)) // fastest convert method, do not use "binary.Read"
	return
}

// After benchmarking the "encoding/binary" way, it takes almost 4 times longer than int -> string -> byte
func WriteInt32(n int32) (buf []byte) {
	buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n)) // fastest convert method, do not use "binary.Write"
	return
}

func WriteUInt16BE(n int) (buf []byte) {
	buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(n)) // fastest convert method, do not use "binary.Write"
	return
}

func RandInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandInt32(min int32, max int32) int32 {
	return min + rand.Int31n(max-min)
}

func RandInt(min int, max int) int {
	diff := int32(max - min)
	if diff <= 0 {
		return 0
	}
	return int(int32(min) + rand.Int31n(diff))
}

func GetNow() time.Time {
	t := time.Now()
	return t
}

func GetUnixTime() int64 {
	t := GetNow().Unix()
	return t
}

func GetMicroTime() int64 {
	//t := GetNow().UnixNano()
	return GetNow().UnixNano() / int64(time.Millisecond)
}

func GetDateTime() string {
	t := GetNow()
	//t.Format("2016-01-01 23:59:59") timezone 안 없어짐
	y, m, d := t.Date()
	h, i, s := t.Clock()
	ymdhis := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", y, m, d, h, i, s)
	return ymdhis
}

type TimerInfo struct {
	CallBack func()
	Delay    time.Duration
	Info     interface{}
}

func SetInterval(callback func(), delay int64) {
	td := time.Duration(delay) * time.Millisecond
	timer := time.NewTicker(td)
	quit := make(chan struct{})
	go func() {
		for {
			//time.Sleep(time.Millisecond * 100)

			select {
			case <-timer.C:
				callback()
			case <-quit:
				timer.Stop()
				return
			}
		}
	}()
}

func SetTimeout(callback func(), delay time.Duration) {
	time.AfterFunc(time.Duration(delay), callback)
}

type MutexInt struct {
	mu sync.RWMutex
	v  int
}

func (mi *MutexInt) Get() int {
	mi.mu.RLock()
	defer mi.mu.RUnlock()
	return mi.v
}

func (mi *MutexInt) Set(v int) {
	mi.mu.Lock()
	mi.v = v
	defer mi.mu.Unlock()
}

func (mi *MutexInt) Increment(i int) {
	mi.mu.Lock()
	mi.v += i
	defer mi.mu.Unlock()
}

func (mi *MutexInt) Decrement(i int) {
	mi.mu.Lock()
	mi.v -= i
	defer mi.mu.Unlock()
}

type MutexInt64 struct {
	mu sync.RWMutex
	v  int64
}

func (mi *MutexInt64) Get() int64 {
	mi.mu.RLock()
	defer mi.mu.RUnlock()
	return mi.v
}

func (mi *MutexInt64) Set(v int64) {
	mi.mu.Lock()
	mi.v = v
	defer mi.mu.Unlock()
}

func (mi *MutexInt64) Increment(i int64) {
	mi.mu.Lock()
	mi.v += i
	defer mi.mu.Unlock()
}

func (mi *MutexInt64) Decrement(i int64) {
	mi.mu.Lock()
	mi.v -= i
	defer mi.mu.Unlock()
}
