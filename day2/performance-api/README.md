# Performance testing with Go
* v1 => net/http package
* V2 => Gin
* V3 => Echo
* V4 => Fiber

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

V3
```
Running 10s test @ http://localhost:8080/hello
  5 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.97ms    1.83ms  53.76ms   93.60%
    Req/Sec    26.42k    18.30k  124.31k    84.80%
  1319394 requests in 10.07s, 163.58MB read
Requests/sec: 131017.30
Transfer/sec:     16.24MB
```

V4
```
Running 10s test @ http://localhost:8080/hello
  5 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.92ms    3.82ms  99.23ms   98.17%
    Req/Sec    32.39k    11.19k  173.56k    88.38%
  1614154 requests in 10.04s, 207.82MB read
Requests/sec: 160766.15
Transfer/sec:     20.70MB
```