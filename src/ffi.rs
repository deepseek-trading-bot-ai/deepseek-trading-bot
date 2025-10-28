#[no_mangle]
pub extern "C" fn init_core() {
    println!("Rust core bridge active via FFI!");
}
