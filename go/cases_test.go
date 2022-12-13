package lamport_clocks

type timeTest struct {
	description string
	input       [][]Event
	expected    [][]int
}

var timeTests = []timeTest{
	{
		description: "zero servers",
		input:       [][]Event{},
		expected:    [][]int{},
	},
	{
		description: "four servers with local events",
		input: [][]Event{
			{
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
			},
			{ // (Local Event), (Send Msg1), (Recv Msg2), (Send, Msg4)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
			},
			{ // (Local Event), (Recv Msg4)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
			},
			{ // (Local Event), (Recv Msg4)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 0, msg_num: 0},
			},
		},
		expected: [][]int{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
	},
	{
		description: "three servers 4 messages",
		input: [][]Event{
			{ // (Recv Msg1), (Send Msg2)
				Event{evt_type: 2, msg_num: 1},
				Event{evt_type: 1, msg_num: 2},
			},
			{ // (Local Event), (Send Msg1), (Recv Msg2), (Send, Msg4)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 1, msg_num: 1},
				Event{evt_type: 2, msg_num: 2},
				Event{evt_type: 1, msg_num: 4},
			},
			{ // (Local Event), (Recv Msg4)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 4},
			},
		},
		expected: [][]int{
			{3, 4},
			{1, 2, 5, 6},
			{1, 7},
		},
	},
	{
		description: "ten servers 4 messages",
		input: [][]Event{
			{ // (Recv Msg1), (Send Msg2)
				Event{evt_type: 2, msg_num: 1},
				Event{evt_type: 1, msg_num: 2},
			},
			{ // (Local Event), (Send Msg1), (Recv Msg2), (Send, Msg4)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 1, msg_num: 1},
				Event{evt_type: 2, msg_num: 2},
				Event{evt_type: 1, msg_num: 4},
			},
			{ // (Local Event), (Recv Msg4), (Send 5)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 4},
				Event{evt_type: 1, msg_num: 5},
			},
			{ // (Local Event), (Recv Msg5), (Send 6)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 5},
				Event{evt_type: 1, msg_num: 6},
			},
			{ // (Local Event), (Recv Msg6), (Send 7)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 6},
				Event{evt_type: 1, msg_num: 7},
			},
			{ // (Local Event), (Recv Msg7), (Send 8)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 7},
				Event{evt_type: 1, msg_num: 8},
			},
			{ // (Local Event), (Recv Msg8), (Send 9)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 8},
				Event{evt_type: 1, msg_num: 9},
			},
			{ // (Local Event), (Recv Msg9), (Send 10)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 9},
				Event{evt_type: 1, msg_num: 10},
			},
			{ // (Local Event), (Recv Msg10), (Send 11)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 10},
				Event{evt_type: 1, msg_num: 11},
			},
			{ // (Local Event), (Recv Msg11)
				Event{evt_type: 0, msg_num: 0},
				Event{evt_type: 2, msg_num: 11},
			},
		},
		expected: [][]int{
			{3, 4},
			{1, 2, 5, 6},
			{1, 7, 8},
			{1, 9, 10},
			{1, 11, 12},
			{1, 13, 14},
			{1, 15, 16},
			{1, 17, 18},
			{1, 19, 20},
			{1, 21},
		},
	},
}
