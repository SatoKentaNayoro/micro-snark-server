//go:build cgo
// +build cgo

package snark_ffi

//#cgo LDFLAGS: -L${SRCDIR} -lsnarkcrypto
//#include "./snarkcrypto.h"
import "C"
import (
	"micro-snark-server/internal/task"
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

func SnarkPost(t *task.Task) ([]byte, error) {
	vp := C.fil_VanillaProof{
		ptr: (*C.uint8_t)((unsafe.Pointer(&t.VanillaProof[0]))),
		len: C.size_t(len(t.VanillaProof)),
	}

	pi := C.fil_PubIn{
		ptr: (*C.uint8_t)((unsafe.Pointer(&t.PubIn[0]))),
		len: C.size_t(len(t.PubIn)),
	}

	p_c := C.fil_PostConfig{
		ptr: (*C.uint8_t)((unsafe.Pointer(&t.PostConfig[0]))),
		len: C.size_t(len(t.PostConfig)),
	}

	__ret := C.fil_snark_post(vp, pi, p_c, C.size_t(t.ReplicasLen))
	runtime.KeepAlive(t)

	myslice := C.GoBytes(unsafe.Pointer(__ret.proofs_ptr), C.int(__ret.proofs_len))

	_ = []byte(myslice)

	return nil, nil
}
