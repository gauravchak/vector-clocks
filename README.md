# vector-clocks

We are given a set of events per server. For example:

```py
[
    [(2, 1), (1, 2)],
    [(0, 0), (1, 1), (2, 2), (1, 4)],
    [(0, 0), (2, 4)]
]
```

Timing rules:

- All server timestamps start with 0
- Local Event: The timestamp of the server increments
- Send : The timestamp of the server is incremented and a message is sent with that timestamp.
- Recv : The timestamp of the server is set to 1 plus higher of the timestamp of the server before this event and the timestamp of the received message.

The output is meant to be the event timestamps on that server. For the input above, the correct output would be: `{{3, 4}, {1 , 2 , 5 , 6}, {1 , 7}}`. 
Here, the first vector `{3, 4}` are the timestamps of the two events of S0 for instance.
