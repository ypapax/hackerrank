Task https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem

Check if number is prime for time less than square root of n.

```
$ go install && cat in.txt | prime_number 2>/dev/null
Not prime
Prime
Prime
```

```
$ go test -v

--- PASS: TestLargestHourGlassSum (0.00s)
    --- PASS: TestLargestHourGlassSum/test3 (0.00s)
    --- PASS: TestLargestHourGlassSum/test12 (0.00s)
    --- PASS: TestLargestHourGlassSum/test5 (0.00s)
    --- PASS: TestLargestHourGlassSum/test7 (0.00s)
    --- PASS: TestLargestHourGlassSum/test104729 (0.00s)
    --- PASS: TestLargestHourGlassSum/test104730 (0.00s)
PASS
ok  	github.com/ypapax/hackerrank/prime_number	0.009s
```
