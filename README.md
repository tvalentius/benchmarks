
## Benchmarked Language + Framework

- Python (2.7.10) - [Docs](https://docs.python.org/2/)
- Python + Flask (0.12.2) - [Docs](http://flask.pocoo.org/docs/0.12/)
- NodeJS (6.11.0) - [Docs](https://nodejs.org/dist/latest-v6.x/docs/api/)
- NodeJS + Express (4.15) - [Docs](http://expressjs.com/en/4x/api.html)
- Go (1.9) - [Docs](https://golang.org/doc/)
- Go + Iris (8.3.4) - [Docs](https://iris-go.com/)
- PHP (5.5.27) - [Docs](https://secure.php.net/docs.php)
- PHP + Symfony (3.3.8) - [Docs](https://symfony.com/doc/current/index.html)
- .Net Core ?
- Ruby ?
- Rust ?
- Dart ?
- C++ ?
- Java ?
- HaXe ?
- Erlang ?

## Machine

- MacBook Pro (Early 2015)
- 2,7 GHz Intel Core i5
- Memory 16 GB 1867 MHz DDR3

## Method

### Simple HTTP Server
- Threads 12, Durations 10s, Connections 1000
- Python VS Node VS Go VS PHP VS...

### CPU Intensive Task
- TBA

## Tools

WRK : [NPM](https://www.npmjs.com/package/wrk)

```
wrk -t12 -d10s -c1000 http://localhost:1337

```
## Todos

- create a minimalist visualisations ( D3JS ? )
- create a blog/medium post
- Database Integration ?
- ???

## Results

### Simple HTTP Server

#### Python 2.7.10 (Bare)

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

#### Python 2.7.10 + Flask 0.12.2

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

#### NodeJS 6.11.0 (Bare)

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

#### NodeJS 6.11.0 + Express 4.15

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

#### Go 1.9 (Bare)

```
12 threads and 1000 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.17ms    4.95ms  90.52ms   73.40%
    Req/Sec     3.85k     1.77k   13.73k    80.32%

  459583 requests in 10.05s, 54.79MB read
  Socket errors: connect 0, read 882, write 1, timeout 0
Requests/sec:  45733.95
Transfer/sec:      5.45MB
```


#### Go 1.9 + Iris 8.3.4

```
12 threads and 1000 connections
Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    31.91ms   25.18ms 270.52ms   74.02%
    Req/Sec     2.57k   541.47     6.11k    75.92%
  306795 requests in 10.09s, 37.16MB read
  Socket errors: connect 0, read 525, write 41, timeout 0
Requests/sec:  30391.73
Transfer/sec:      3.68MB
```

#### PHP 5.5.27 (Bare)

```
12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.18ms    5.36ms 156.90ms   92.89%
    Req/Sec   393.43    481.07     2.00k    83.67%
  9235 requests in 10.07s, 1.06MB read
  Socket errors: connect 757, read 9803, write 0, timeout 0
Requests/sec:    916.65
Transfer/sec:    107.42KB
```

#### PHP 5.5.27 + Symfony 3.3.8

```
12 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.05s   569.22ms   1.91s    63.64%
    Req/Sec     5.88      2.70    10.00     67.24%
  58 requests in 10.08s, 16.37KB read
  Socket errors: connect 757, read 234, write 0, timeout 47
Requests/sec:      5.75
Transfer/sec:      1.62KB
```
