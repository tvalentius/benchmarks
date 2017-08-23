## Python VS Node VS ...(TBA)


```
wrk -t12 -d10s -c1000 http://localhost:1337

```


### Python 2.7.10 (Bare)

```
12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.94ms  657.67us   7.13ms   85.41%
    Req/Sec    66.00    107.82     1.12k    90.76%
  4358 requests in 10.10s, 540.49KB read
  Socket errors: connect 757, read 5321, write 7, timeout 0
Requests/sec:    431.67
Transfer/sec:     53.54KB
```

### Python 2.7.10 + Flask 0.12.2

```
12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   157.87ms   56.94ms 619.02ms   94.05%
    Req/Sec    72.60     54.47   240.00     67.87%
  8069 requests in 10.09s, 1.27MB read
  Socket errors: connect 757, read 342, write 0, timeout 0
Requests/sec:    799.68
Transfer/sec:    128.89KB
```

### NodeJS 6.11.0 (Bare)

```
12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    15.41ms    3.48ms 351.07ms   94.75%
    Req/Sec     1.30k     0.95k    3.34k    62.83%
  155065 requests in 10.04s, 22.63MB read
  Socket errors: connect 757, read 49, write 0, timeout 0
Requests/sec:  15451.29
Transfer/sec:      2.25MB
```

### NodeJS 6.11.0 + Express 4.15

```
12 threads and 1000 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    44.12ms   16.56ms 257.37ms   88.55%
    Req/Sec   441.63    252.31     1.16k    70.79%
  52473 requests in 10.10s, 10.71MB read
  Socket errors: connect 757, read 165, write 5, timeout 0
Requests/sec:   5196.13
Transfer/sec:      1.06MB
```
