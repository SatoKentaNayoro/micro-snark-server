use libc;
use ffi_toolkit::{free_c_str, code_and_message_impl, CodeAndMessage, FCPResponseStatus};
use drop_struct_macro_derive::DropStructMacro;
use std::ptr;

#[repr(C)]
#[derive(DropStructMacro)]
pub struct fil_SnarkPostResponse {
    pub error_msg: *const libc::c_char,
    pub proofs_len: libc::size_t,
    pub proofs_ptr: *const u8,
    pub status_code: FCPResponseStatus,
}

impl Default for fil_SnarkPostResponse {
    fn default() -> fil_SnarkPostResponse {
        fil_SnarkPostResponse {
            error_msg: ptr::null(),
            proofs_len: 0,
            proofs_ptr: ptr::null(),
            status_code: FCPResponseStatus::FCPNoError,
        }
    }
}

code_and_message_impl!(fil_SnarkPostResponse);


#[repr(C)]
#[derive(Clone)]
pub struct fil_VanillaProof {
    pub ptr: *const u8,
    pub len: libc::size_t,
}

#[repr(C)]
#[derive(Clone)]
pub struct fil_PubIn {
    pub ptr: *const u8,
    pub len: libc::size_t,
}

#[repr(C)]
#[derive(Clone)]
pub struct fil_PostConfig {
    pub ptr: *const u8,
    pub len: libc::size_t,
}

pub struct TaskInfo {
    pub vanilla_proof_u8: fil_VanillaProof,
    pub pub_in_u8: fil_PubIn,
    pub post_config_u8: fil_PostConfig,
    pub replicas_len: libc::size_t,
}