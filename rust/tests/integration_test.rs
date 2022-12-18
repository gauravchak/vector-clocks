use vector_clocks;

#[test]
fn empty_events() {
	assert_eq !(0, vector_clocks::get_timestamps([]));
}
