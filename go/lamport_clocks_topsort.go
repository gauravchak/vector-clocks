package lamport_clocks

func GetTimestampsTopSort(input [][]Event) [][]int {
	// the list of timestamps corresponding to server events.
	timestamps := make([][]int, len(input))
	for i := range timestamps {
		timestamps[i] = make([]int, len(input[i]))
	}

	type TimedEvent struct {
		event     Event         // original event
		s_idx     int           // index of the server this event is from
		e_idx     int           // index in the server's events
		timestamp int           // the timestamp of this event
		blockers  int           // number of events blocking this
		blocking  []*TimedEvent // the events this one is blocking
	}

	msg_to_send_evt := make(map[int]*TimedEvent) // map from msgnum to sender
	msg_to_recv_evt := make(map[int]*TimedEvent) // map from msgnum to recver

	// create a list of unblocked_events.
	unblocked_events := make([]*TimedEvent, 0)
	// For every timedevent of the first event of a server
	// if it is a local event or a send event add it to this list.

	// Loop over all events and create a TimedEvent node for each.
	for i := range input {
		var prev_te *TimedEvent // this is the TE of the previous event in the server
		for j := range input[i] {
			te := new(TimedEvent)
			te.event = input[i][j]
			te.s_idx = i
			te.e_idx = j
			te.timestamp = 0
			te.blockers = 0                      // set it to 0 and then change
			te.blocking = make([]*TimedEvent, 0) // empty initially

			if j > 0 {
				// set blockers and prev_te's blocking
				te.blockers++
				prev_te.blocking = append(prev_te.blocking, te)
			}
			if te.event.evt_type == 1 {
				// if receiver seen already then add that as blocking
				// and increment receiver's blockers
				if recver, seen := msg_to_recv_evt[te.event.msg_num]; seen {
					te.blocking = append(te.blocking, recver)
				} else {
					msg_to_send_evt[te.event.msg_num] = te
				}
			}
			if te.event.evt_type == 2 {
				if sender, seen := msg_to_send_evt[te.event.msg_num]; seen {
					sender.blocking = append(sender.blocking, te)
					te.blockers++
				} else {
					msg_to_recv_evt[te.event.msg_num] = te
					te.blockers++
				}
			}

			if te.blockers == 0 {
				unblocked_events = append(unblocked_events, te)
			}

			prev_te = te
		}
	}

	// start with unblocked events process them and then unblock whoever this one clears up.
	for len(unblocked_events) > 0 {
		te := unblocked_events[0]
		unblocked_events = unblocked_events[1:] // drop first TimedEvent
		te.timestamp++
		timestamps[te.s_idx][te.e_idx] = te.timestamp
		for _, blocked_te := range te.blocking {

			if blocked_te.timestamp < te.timestamp {
				blocked_te.timestamp = te.timestamp
			}
			blocked_te.blockers--

			if blocked_te.blockers == 0 {
				unblocked_events = append(unblocked_events, blocked_te)
			}
		}
	}
	return timestamps
}
