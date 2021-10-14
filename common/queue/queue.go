package queue

// push { queue, payload, releaseAt } | { msgId }
// pop { } | { msgId, queue, payload } **should this be an event topic?
// ack { msgId } | { }
// count { queue } | int
// list { } | []string
