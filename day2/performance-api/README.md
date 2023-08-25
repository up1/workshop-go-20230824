# Performance testing with Go
* v1 => net/http package

## Testing
```
$wrk -t 5 -c 100 -d 10s http://localhost:8080/hello
```