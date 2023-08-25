# Performance testing with Go
* v1 => net/http package
* V2 => Gin

## Testing
```
$wrk -t 5 -c 100 -d 10s http://localhost:8080/hello
```


## Test results
V1
```
Running 10s test @ http://localhost:8080/hello
  5 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.24ms    3.14ms  94.24ms   93.53%
    Req/Sec    27.24k    18.80k  149.42k    79.36%
  1362901 requests in 10.10s, 170.27MB read
Requests/sec: 134980.61
Transfer/sec:     16.86MB
```

V2
```
Running 10s test @ http://localhost:8080/hello
  5 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.20ms    3.29ms 111.95ms   95.74%
    Req/Sec    25.80k    16.11k  119.53k    79.32%
  1288005 requests in 10.10s, 173.20MB read
Requests/sec: 127587.41
Transfer/sec:     17.16MB
```