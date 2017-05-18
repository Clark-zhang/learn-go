wrk.method = "POST"
wrk.body = "pid=0&fpid=0&eid=1&uid=1&comment=this is a comment from wrk lua"
wrk.headers["Host"] = "localhost"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"

logfile = io.open("/tmp/clark-learn-go-cs-wrk.log", "w");
local cnt = 0;

response = function(status, header, body)
  logfile:write("status:" .. status .. "\n");
  cnt = cnt + 1;
  logfile:write("status:" .. status .. "\n" .. body .. "\n-------------------------------------------------\n");
end

done = function(summary, latency, requests)
  logfile:write("------------- SUMMARY -------------\n")
  print("Response count: ", cnt)
  logfile.close();
end