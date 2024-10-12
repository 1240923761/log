# How to use
```go
func main() {
    log.SetWXAddress("set your wxBOT here.")
	log.Debug(nil, "debug here")
    log.WX(nil, "trace_id")
}
```
it goes like this:

``
FILEBEATS [DEBUG] | time | trace_id | msg
``

