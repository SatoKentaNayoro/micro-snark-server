use filecoin_proofs::caches::get_post_params;
use filecoin_proofs::{get_partitions_for_window_post, PoStConfig, with_shape};
use storage_proofs_core::{compound_proof, error::Result, merkle::MerkleTreeTrait};
use crate::types::TaskInfo;
use std::slice::from_raw_parts;
use filecoin_proofs::parameters::window_post_setup_params;
use serde_json::Value;
use storage_proofs_core::compound_proof::CompoundProof;
use storage_proofs_post::fallback::{FallbackPoSt, FallbackPoStCompound};

pub unsafe fn run_snark(task_info: &TaskInfo) -> Result<Vec<u8>> {
    let post_config = get_post_config(from_raw_parts(task_info.post_config_u8.ptr, task_info.post_config_u8.len).to_vec().as_ref())?;
    let vanilla_v = serde_json::from_slice(from_raw_parts(task_info.vanilla_proof_u8.ptr, task_info.vanilla_proof_u8.len).to_vec().as_ref())?;
    let pub_in_v = serde_json::from_slice(from_raw_parts(task_info.pub_in_u8.ptr, task_info.pub_in_u8.len).to_vec().as_ref())?;
    with_shape!(post_config.sector_size.0,proof_snark,&post_config,pub_in_v,vanilla_v,task_info.replicas_len)
}

fn proof_snark<Tree: 'static + MerkleTreeTrait>(post_config: &PoStConfig, pub_in_v: Value, vanilla_proofs_v: Value, replicas_len: usize) -> Result<Vec<u8>> {
    let vanilla_params = window_post_setup_params(&post_config);
    let partitions = get_partitions_for_window_post(replicas_len, &post_config);
    let setup_params = compound_proof::SetupParams {
        vanilla_params,
        partitions,
        priority: post_config.priority,
    };

    let pub_params: compound_proof::PublicParams<'_, FallbackPoSt<'_, Tree>> =
        FallbackPoStCompound::setup(&setup_params)?;
    let groth_params = get_post_params::<Tree>(&post_config)?;
    let proof = FallbackPoStCompound::prove_with_vanilla_by_snark_server(
        &pub_params,
        pub_in_v,
        vanilla_proofs_v,
        &groth_params,
    )?;
    proof.to_vec()
}


fn get_post_config(post_config_u8: &Vec<u8>) -> Result<PoStConfig> {
    let post_config_v = serde_json::from_slice(post_config_u8)?;
    let post_config = serde_json::from_value::<PoStConfig>(post_config_v)?;
    Ok(post_config)
}