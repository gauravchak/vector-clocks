package lamport_clocks

func GetTimestampsSimple(input [][]Event) [][]int {
	// the list of timestamps corresponding to server events.
	timestamps := make([][]int, len(input))

	// the index of the next event being processed from that server/event-stream
	next_s_idx := make([]int, len(input))
	current_ts := make([]int, len(input))

	// the send time map of a message. send_ts[123] is the timestamp when
	// message 123 was sent.
	send_ts := make(map[int]int)

	making_progress := true
	for making_progress {
		making_progress = false
		for s_idx := 0; s_idx < len(input); s_idx++ {
			if next_s_idx[s_idx] < len(input[s_idx]) {
				// events are still unprocessed for this server.
				next_evt := input[s_idx][next_s_idx[s_idx]]
				switch next_evt.evt_type {
				case 0:
					{ // Local Event
						current_ts[s_idx]++
						timestamps[s_idx] = append(timestamps[s_idx], current_ts[s_idx])
						next_s_idx[s_idx]++
						making_progress = true
					}
				case 1:
					{ // Send a message
						current_ts[s_idx]++
						timestamps[s_idx] = append(timestamps[s_idx], current_ts[s_idx])
						send_ts[next_evt.msg_num] = current_ts[s_idx]
						next_s_idx[s_idx]++
						making_progress = true
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
							making_progress = true
						}
					}
				}
			}
		}
	}
	return timestamps
}
