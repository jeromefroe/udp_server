A short example that when implementing a UDP server, the server
must guarantee that its buffer is large enough to read the biggest
message that could be sent. Otherwise, if the buffer is not large
enough, the UDP message will be truncated to the size of the buffer.