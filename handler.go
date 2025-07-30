package main

import "sync"

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET" : set,
	"GET" : get,
	"HSET": hset,
	"HGET": hget,
}

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}

func set(args []Value) Value {
	if len(args) != 2{
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}
	key := args[0].bulk
	value := args[1].bulk

	SETsMu.Lock()
	SETs[key] = value
	SETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1{
		return Value{typ: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}

}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}
	return Value{typ: "string", str: args[0].bulk}
}

func hset(args []Value) Value {
	if len(args) != 3{
		return Value{typ: "error", str: "ERR wrong number of the arguments 'hset' command"}
	}
	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSETsMu.Lock()
	if _, ok:= HSETs[hash]; !ok{
		HSETs[hash] = map[string]string{}
	}

	HSETs[hash][key] = value
	HSETsMu.Unlock()

	return Value{typ: "string", str: "OK"}
}

func hget(args []Value) Value{
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of argument for 'hget' command."}
	}

	hash := args[0].bulk
	key  := args[1].bulk

	HSETsMu.Lock()
	value, ok := HSETs[hash][key]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk" , bulk: value}
}