package lamport_clocks

func GetTimestampsOptimized(input [][]Event) [][]int {
	// the list of timestamps corresponding to server events.
	timestamps := make([][]int, len(input))

	// creating an arrat from 0 to num_servers - 1 as the order in which to process servers.

	var s_idx_to_proc []int
	for i := 0; i < len(input); i++ {
		s_idx_to_proc = append(s_idx_to_proc, i)
	}

	// the index of the next event being processed from that server/event-stream
	next_s_idx := make([]int, len(input))
	current_ts := make([]int, len(input))

	// the send time map of a message. send_ts[123] is the timestamp when
	// message 123 was sent.
	send_ts := make(map[int]int)

	// the server that is blocked on someone sending this message
	blocked_server := make(map[int]int)

	s_idx_to_proc_idx := 0
	for s_idx_to_proc_idx < len(s_idx_to_proc) {
		// take the next server to process events and increment index
		s_idx := s_idx_to_proc[s_idx_to_proc_idx]
		s_idx_to_proc_idx++

		if next_s_idx[s_idx] >= len(input[s_idx]) {
			// all events for this server have been processed
			continue
		}

		next_evt := input[s_idx][next_s_idx[s_idx]]
		switch next_evt.evt_type {
		case 0:
			{ // Local Event
				current_ts[s_idx]++
				timestamps[s_idx] = append(timestamps[s_idx], current_ts[s_idx])
				next_s_idx[s_idx]++

				// redo this server ... kind of like depth first ...
				// other approach would have been to add s_idx at the end.
				s_idx_to_proc_idx--
				s_idx_to_proc[s_idx_to_proc_idx] = s_idx
			}
		case 1:
			{ // Send a message
				current_ts[s_idx]++
				timestamps[s_idx] = append(timestamps[s_idx], current_ts[s_idx])
				send_ts[next_evt.msg_num] = current_ts[s_idx]
				next_s_idx[s_idx]++

				block_sidx, ok := blocked_server[next_evt.msg_num]
				if ok {
					s_idx_to_proc_idx--
					s_idx_to_proc[s_idx_to_proc_idx] = block_sidx // switch to blocked server
					s_idx_to_proc = append(s_idx_to_proc, s_idx)  // and add this server at the end.
				} else {
					// redo this server ... kind of like depth first ...
					// other approach would have been to add s_idx at the end.
					s_idx_to_proc_idx--
					s_idx_to_proc[s_idx_to_proc_idx] = s_idx
				}
			}
		case 2:
			{
				// TODO: Recv message
				msg_send_ts, msg_sent_already := send_ts[next_evt.msg_num]
				if msg_sent_already {
					if msg_send_ts > current_ts[s_idx] {
						current_ts[s_idx] = msg_send_ts
					}
					current_ts[s_idx]++
					timestamps[s_idx] = append(timestamps[s_idx], current_ts[s_idx])
					next_s_idx[s_idx]++
					// redo this index
					s_idx_to_proc_idx--
					s_idx_to_proc[s_idx_to_proc_idx] = s_idx
				} else {
					// blocked
					blocked_server[next_evt.msg_num] = s_idx
				}
			}

		}
	}
	return timestamps
}
