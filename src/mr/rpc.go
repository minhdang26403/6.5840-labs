package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

// Add your RPC definitions here.
type MapArgs struct {
	TaskNumber        int
	IntermediateFiles []string
}

type MapReply struct {
	Filename   string
	NReduce    int
	TaskNumber int
}

type ReduceArgs struct {
	TaskNumber int
}

type ReduceReply struct {
	TaskNumber        int
	IntermediateFiles []string
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
