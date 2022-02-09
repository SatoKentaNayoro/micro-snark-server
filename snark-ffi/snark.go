package snark_ffi

//#cgo LDFLAGS: -L${SRCDIR}/rust/target/debug -lsnarkcrypto
//#include "./snarkcrypto.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type FCPResponseStatus int32

const (
	FCPResponseStatusFCPNoError           FCPResponseStatus = C.FCPResponseStatus_FCPNoError
	FCPResponseStatusFCPUnclassifiedError FCPResponseStatus = C.FCPResponseStatus_FCPUnclassifiedError
	FCPResponseStatusFCPCallerError       FCPResponseStatus = C.FCPResponseStatus_FCPCallerError
	FCPResponseStatusFCPReceiverError     FCPResponseStatus = C.FCPResponseStatus_FCPReceiverError
)

type FilSnarkPostResponse struct {
	ErrorMsg   string
	ProofsLen  uint
	ProofsPtr  []byte
	StatusCode FCPResponseStatus
	refResp    *C.fil_SnarkPostResponse
}

func SnarkPost() {
	v := []uint8{1, 2, 3}
	p := []uint8{1, 2, 3}
	pc := []uint8{1, 2, 3}
	rl := 3

	vp := C.fil_VanillaProof{
		ptr: (*C.uint8_t)((unsafe.Pointer(&v))),
		len: C.size_t(len(v)),
	}

	pi := C.fil_PubIn{
		ptr: (*C.uint8_t)((unsafe.Pointer(&p))),
		len: C.size_t(len(p)),
	}

	p_c := C.fil_PostConfig{
		ptr: (*C.uint8_t)((unsafe.Pointer(&pc))),
		len: C.size_t(len(pc)),
	}
	runtime.KeepAlive(vp)
	runtime.KeepAlive(pi)
	runtime.KeepAlive(p_c)
	runtime.KeepAlive(rl)
	__ret := C.fil_snark_post(vp, pi, p_c, C.size_t(rl))
	fmt.Println(__ret)
}
