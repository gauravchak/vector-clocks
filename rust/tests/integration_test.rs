use vector_clocks;

#[test]
fn local_events() {
	let events = vec ![
		vec ![ (0, 0), (0, 0), (0, 0) ],
		vec ![ (0, 0), (0, 0) ],
		vec ![(0, 0)]
	];
	let res = vector_clocks::get_timestamps(&events);
	let expected = vec ![ vec ![ 1, 2, 3 ], vec ![ 1, 2 ], vec ![1] ];
	assert_eq !(res, expected);
}

#[test]
fn local_events_and_sends() {
	let events = vec ![
		vec ![ (0, 0), (1, 2), (1, 3) ],
		vec ![ (0, 0), (1, 4) ],
		vec ![(1, 7)]
	];
	let res = vector_clocks::get_timestamps(&events);
	let expected = vec ![ vec ![ 1, 2, 3 ], vec ![ 1, 2 ], vec ![1] ];
	assert_eq !(res, expected);
}

#[test]
fn all_events_no_stuck() {
	let events = vec ![
		vec ![ (0, 0), (1, 2), (2, 2) ],
		vec ![ (1, 4), (2, 4) ],
		vec ![ (1, 7), (2, 7) ]
	];
	let res = vector_clocks::get_timestamps(&events);
	let expected = vec ![ vec ![ 1, 2, 3 ], vec ![ 1, 2 ], vec ![ 1, 2 ] ];
	assert_eq !(res, expected);
}

#[test]
fn all_events_some_stuck() {
	let events = vec ![
		vec ![ (2, 1), (1, 2) ],
		vec ![ (0, 0), (1, 1), (2, 2), (1, 4) ],
		vec ![ (0, 0), (2, 4) ]
	];
	let res = vector_clocks::get_timestamps(&events);
	let expected = vec ![ vec ![ 3, 4 ], vec ![ 1, 2, 5, 6 ], vec ![ 1, 7 ] ];
	assert_eq !(res, expected);
}
