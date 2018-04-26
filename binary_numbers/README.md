[Task](https://www.hackerrank.com/challenges/30-binary-numbers/problem)

Given a base- integer, , convert it to binary (base-). Then find and print the base- integer denoting the maximum number of consecutive 's in 's binary representation.

```
$ go test -v
=== RUN   TestToBinaryAndOneCount
=== RUN   TestToBinaryAndOneCount/test13
=== RUN   TestToBinaryAndOneCount/test5
=== RUN   TestToBinaryAndOneCount/test15
--- PASS: TestToBinaryAndOneCount (0.00s)
    --- PASS: TestToBinaryAndOneCount/test13 (0.00s)
    --- PASS: TestToBinaryAndOneCount/test5 (0.00s)
    --- PASS: TestToBinaryAndOneCount/test15 (0.00s)
PASS
ok  	github.com/ypapax/hackerrank/binary_numbers	0.008s

```

```
$ echo 5 | go run main.go
1

```