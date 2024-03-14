## Note

distributed log with hashicorp raft algorithm


```
Current State:
- store read write data ok
- raft ok
```

```
TODO:
- clean from dead code
- refactor /internal/log
- create a new segment when the maximum size is reached
- make the readme file
```

```
changes:
- added grpc client and server
- added event 
- event cleanup
```