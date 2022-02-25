//go:build cgo
// +build cgo

package snark_ffi

//#cgo LDFLAGS: -L${SRCDIR} -lsnarkcrypto
//#include "./snarkcrypto.h"
import "C"
import (
	"fmt"
	"micro-snark-server/internal/task"
	"reflect"
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
	ProofsPtr  *C.char
	StatusCode FCPResponseStatus
	refResp    *C.fil_SnarkPostResponse
}

func (x *FilSnarkPostResponse) Deref() {
	if x.refResp == nil {
		return
	}
	x.ErrorMsg = packPCharString(x.refResp.error_msg)
	x.ProofsPtr = x.refResp.proofs_ptr
	x.StatusCode = (FCPResponseStatus)(x.refResp.status_code)
	x.ProofsLen = uint(x.refResp.proofs_len)
}

type stringHeader struct {
	Data unsafe.Pointer
	Len  int
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = unsafe.Pointer(p)
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - uintptr(h.Data))
	}
	return
}

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
	ref := NewFilSnarkPostResponseRef(unsafe.Pointer(__ret))
	ref.Deref()
	fmt.Println("ptr:", ref.ProofsPtr)
	fmt.Println("err:", ref.ErrorMsg)
	fmt.Println("code:", ref.StatusCode)
	fmt.Println("len:", ref.ProofsLen)
	var result []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&result)))
	sliceHeader.Cap = int(ref.ProofsLen)
	sliceHeader.Len = int(ref.ProofsLen)
	sliceHeader.Data = uintptr(unsafe.Pointer(ref.ProofsPtr))
	fmt.Println(result)
	return nil, nil
}

func NewFilSnarkPostResponseRef(ref unsafe.Pointer) *FilSnarkPostResponse {
	if ref == nil {
		return nil
	}
	obj := new(FilSnarkPostResponse)
	obj.refResp = (*C.fil_SnarkPostResponse)(ref)
	return obj
}
