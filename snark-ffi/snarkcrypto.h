#include <stdint.h>
#include <libc.h>

typedef enum {
  FCPResponseStatus_FCPNoError = 0,
  FCPResponseStatus_FCPUnclassifiedError = 1,
  FCPResponseStatus_FCPCallerError = 2,
  FCPResponseStatus_FCPReceiverError = 3,
} FCPResponseStatus;

typedef struct {
    const char *error_msg;
    size_t proofs_len;
    const uint8_t proofs_ptr;
    FCPResponseStatus status_code;
} fil_SnarkPostResponse;

typedef struct {
    const uint8_t ptr;
    size_t len;
} fil_VanillaProof;

typedef struct {
    const uint8_t ptr;
    size_t len;
} fil_PubIn;

typedef struct {
    const uint8_t ptr;
    size_t len;
} fil_PostConfig;

fil_SnarkPostResponse *fil_snark_post(fil_VanillaProof vanilla_proof,
                                    fil_PubIn p_i,
                                    fil_PostConfig p_c,
                                    size_t replicas_len);