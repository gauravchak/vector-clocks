use std::collections::HashMap;

fn get_timestamps(events : &[Vec<(i32, i32)>])->Vec<Vec<i32>> {
	let mut timestamps : Vec<Vec<i32>> = Vec::new ();
	let mut send_times
		: HashMap<i32, i32> =
			  HashMap::new (); // map from msgnum to timestamp of sender
  for
	  i in 0..events.len() {
		  timestamps.push(Vec::new ());
	  }
  let mut next_eidx : Vec<i32> = vec ![0; events.len()];
  let mut cur_ts : Vec<i32> = vec ![0; events.len()];
  let mut making_progress = true;

  while
	  making_progress {
		  making_progress = false;
		  for (sidx, s_events)
			  in events.iter().enumerate() {
				  if
					  next_eidx[sidx] >= int(s_events.len()) {
						  continue;
					  }
				  let evt = s_events[next_eidx[sidx]];
				  match evt .0 {
					  0 = > {
						  cur_ts[sidx]++;
						  timestamps[sidx].push(cur_ts[sidx]);
					  }
					  , 1 = > {
						  cur_ts[sidx]++;
						  timestamps[sidx].push(cur_ts[sidx]);
						  send_times.insert(evt .0, cur_ts[sidx]);
					  }
					  ,
						  2 = >
						  {
							  // check if msgnum received already
						  },
						  _ = > panic !("crash and burn"),
				  }
			  }
	  }
  timestamps
}
