package lamport_clocks

type Event struct {
	evt_type int // Local = 0, Send = 1, Recv = 2
	msg_num  int // If evt_type is Local, this field should be ignored.
}
