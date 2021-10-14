package email

// send { from, []to, []cc, []bcc, subject, body, attachments } | []{ msgId, type[to, cc, bcc], address, errorMessage }
// status { } | { msgId, type[to, cc, bcc], address, status[delivered,bounced,blocked], message } **set status enum accordingly
// received { } | { ??? } **should this be an event topic?