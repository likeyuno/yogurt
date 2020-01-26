// https://github.com/shadowsocks/shadowsocks-go/blob/master/LICENSE

package reverse

type LeakyBuf struct {
	bufSize  int
	freeList chan []byte
}

const leakyBufSize = 4096
const maxNBuf = 2048

var BufInstance = NewLeakyBuf(maxNBuf, leakyBufSize)

func NewLeakyBuf(n, bufSize int) *LeakyBuf {
	return &LeakyBuf{
		bufSize:  bufSize,
		freeList: make(chan []byte, n),
	}
}

func (lb *LeakyBuf) Get() (b []byte) {
	select {
	case b = <-lb.freeList:
	default:
		b = make([]byte, lb.bufSize)
	}
	return
}

func (lb *LeakyBuf) Put(b []byte) {
	if len(b) != lb.bufSize {
		panic("invalid buffer size that's put into leaky buffer")
	}
	select {
	case lb.freeList <- b:
	default:
	}
	return
}
