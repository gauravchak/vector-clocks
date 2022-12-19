use std::collections::HashMap;

pub fn get_timestamps(events : &[Vec<(i32, i32)>])->Vec<Vec<i32>> {
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
						  send_times.insert(evt .0, cur_ts[sidx]);
						  next_eidx[sidx] = next_eidx[sidx] + 1;
						  making_progress = true;
					  }
				  else if
					  evt .0 == 2 {
						  // check if msgnum received already
					  }
			  }
	  }
  timestamps
}
