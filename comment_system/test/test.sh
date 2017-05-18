############################
# curl
############################
curl localhost:8007/get\?eventId=1\&pageIndex=1\&pageSize=50

curl -X POST -d 'pid=0&fpid=0&eid=1&uid=1&comment=this is a comment' localhost:8007/add



############################
# wrk
############################
wrk -t50 -c100 -d30  http://localhost:8007/get\?eventId\=1\&pageIndex\=1\&pageSize\=50

➜  test git:(master) ✗ wrk -t50 -c100 -d30  http://localhost:8007/get\?eventId\=1\&pageIndex\=1\&pageSize\=50
Running 30s test @ http://localhost:8007/get?eventId=1&pageIndex=1&pageSize=50
  50 threads and 100 connections

  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.46ms   10.81ms 150.09ms   71.98%
    Req/Sec   112.26     26.61   474.00     69.20%
  168268 requests in 30.10s, 271.20MB read
Requests/sec:   5590.42
Transfer/sec:      9.01MB


#Trouble shoot
#Can't create more than max_prepared_stmt_count statements (current value: 16382)
#set global max_prepared_stmt_count=1000000;
wrk -t50 -c100 -d 30s -s ./add-comments.lua http://localhost:8007/add

➜  test git:(master) ✗ wrk -t5 -c10 -d 30s -s ./add-comments.lua http://localhost:8007/add
Running 30s test @ http://localhost:8007/add
  5 threads and 10 connections

  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    36.59ms  128.16ms   1.14s    95.09%
    Req/Sec   228.35     56.52   323.00     78.90%
  32052 requests in 30.26s, 4.25MB read
Requests/sec:   1059.21
Transfer/sec:    143.78KB
Response count:     0