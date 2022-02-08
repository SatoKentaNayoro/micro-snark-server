package snark_ffi

//#cgo LDFLAGS: -L${SRCDIR}/rust/target/debug -lsnarkcrypto
//#include "./snarkcrypto.h"
import "C"

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

func FilSnarkPost() {
	C.fil_snark_post()
}
