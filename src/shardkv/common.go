package shardkv

//
// Sharded key/value server.
// Lots of replica groups, each running Raft.
// Shardctrler decides which group serves each shard.
// Shardctrler may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK             = "OK"
	ErrFuture      = "ErrFuture"
	ErrNoKey       = "ErrNoKey"
	ErrTimeout     = "ErrTimeout"
	ErrWrongGroup  = "ErrWrongGroup"
	ErrWrongLeader = "ErrWrongLeader"
)

type Err string

type OperationArgs struct {
	Key       string
	Value     string
	Method    string
	ClientId  int64
	MessageId int
}

type OperationReply struct {
	Err   Err
	Value string
}

type PullShardArgs struct {
	Num       int
	ShardList []int
}

type PullShardReply struct {
	Err    Err
	Num    int
	State  map[int]Shard
	Client map[int64]int
}

type DeleteShardArgs struct {
	Num       int
	ShardList []int
}

type DeleteShardReply struct {
	Err Err
}
