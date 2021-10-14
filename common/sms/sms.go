package sms

// send { from, []to, text } | { msgIds }
// status { } | { msgId, mobile, status[delivered,bounced,blocked], message } **set status enum accordingly
// received { } | { ??? } **should this be an event topic?