use std::collections::HashMap;

pub fn get_timestamps(events : &[Vec<(i32, i32)>])->Vec<Vec<i32>> {
	return get_timestamps_optimized(events);
}

fn get_timestamps_simple(events : &[Vec<(i32, i32)>])->Vec<Vec<i32>> {
	let mut timestamps : Vec<Vec<i32>> = Vec::new ();
  for
	  _i in 0..events.len() {
		  timestamps.push(Vec::new ());
	  }

  let mut send_times
	  : HashMap<i32, i32> =
			HashMap::new (); // map from msgnum to timestamp of sender

  let mut next_eidx : Vec<usize> = vec ![0; events.len()];
  let mut cur_ts : Vec<i32> = vec ![0; events.len()];
  let mut making_progress = true;

  while
	  making_progress {
		  making_progress = false;
		  for (sidx, s_events)
			  in events.iter().enumerate() {
				  if
					  next_eidx[sidx] >= s_events.len() {
						  continue;
					  }

				  let eidx = next_eidx[sidx];
				  let evt : (i32, i32) = s_events[eidx];

				  if
					  evt .0 == 0 {
						  cur_ts[sidx] = cur_ts[sidx] + 1;
						  timestamps[sidx].push(cur_ts[sidx]);
						  next_eidx[sidx] = next_eidx[sidx] + 1;
						  making_progress = true;
					  }
				  else if
					  evt .0 == 1 {
						  cur_ts[sidx] = cur_ts[sidx] + 1;
						  timestamps[sidx].push(cur_ts[sidx]);
						  send_times.insert(evt .1, cur_ts[sidx]);
						  next_eidx[sidx] = next_eidx[sidx] + 1;
						  making_progress = true;
					  }
				  else if
					  evt .0 == 2 { // recv event
						  if
							  let Some(ts) = send_times.get(&evt .1) {
								  if
									  cur_ts[sidx] < *ts {
										  cur_ts[sidx] = *ts;
									  }
								  cur_ts[sidx] = cur_ts[sidx] + 1;
								  timestamps[sidx].push(cur_ts[sidx]);
								  next_eidx[sidx] = next_eidx[sidx] + 1;
								  making_progress = true;
							  }
						  else {
							  println !("Recv blocked on msg {}", evt .0);
						  }
					  }
			  }
	  }
  timestamps
}

fn get_timestamps_optimized(input : &[Vec<(i32, i32)>])->Vec<Vec<i32>> {
	let mut timestamps = Vec::new ();
	for
		_ in 0..input.len() {
			timestamps.push(Vec::new ());
		}

	let mut s_idx_to_proc = Vec::new ();
	for
		i in 0..input.len() {
			s_idx_to_proc.push(i);
		}

	let mut next_s_idx = vec ![0; input.len()];
	let mut current_ts = vec ![0; input.len()];
	let mut send_ts : HashMap<i32, i32> = HashMap::new ();
	let mut blocked_server : HashMap<i32, usize> = HashMap::new ();
	let mut s_idx_to_proc_idx = 0;

	while
		s_idx_to_proc_idx < s_idx_to_proc.len() {
			let s_idx = s_idx_to_proc[s_idx_to_proc_idx];
			s_idx_to_proc_idx += 1;

			if
				next_s_idx[s_idx] >= input[s_idx].len() {
					continue;
				}

			let next_evt = input[s_idx][next_s_idx[s_idx]];
			match next_evt .0 {
				0 => {
					current_ts[s_idx] += 1;
					timestamps[s_idx].push(current_ts[s_idx]);
					next_s_idx[s_idx] += 1;

					s_idx_to_proc_idx -= 1;
					s_idx_to_proc[s_idx_to_proc_idx] = s_idx;
				}
				1 => {
					current_ts[s_idx] += 1;
					timestamps[s_idx].push(current_ts[s_idx]);
					send_ts.insert(next_evt .1, current_ts[s_idx]);
					next_s_idx[s_idx] += 1;

					let block_sidx = blocked_server.get(&next_evt .1);
					if
						let Some(block_sidx) = block_sidx {
							s_idx_to_proc_idx -= 1;
							s_idx_to_proc[s_idx_to_proc_idx] = *block_sidx;
							s_idx_to_proc.push(s_idx);
						}
					else {
						s_idx_to_proc_idx -= 1;
						s_idx_to_proc[s_idx_to_proc_idx] = s_idx;
					}
				}
				2 => {
					let msg_send_ts = send_ts.get(&next_evt .1);
					if
						let Some(msg_send_ts) = msg_send_ts {
							if
								*msg_send_ts > current_ts[s_idx] {
									current_ts[s_idx] = *msg_send_ts;
								}
							current_ts[s_idx] += 1;
							timestamps[s_idx].push(current_ts[s_idx]);
							next_s_idx[s_idx] += 1;

							s_idx_to_proc_idx -= 1;
							s_idx_to_proc[s_idx_to_proc_idx] = s_idx;
						}
					else {
						blocked_server.insert(next_evt .1, s_idx);
					}
				}
        _ => {}
			}
		}
	timestamps
}