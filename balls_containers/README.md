Task https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem


```
$ ./run.sh 2>/dev/null
Possible
Impossible

```

```
$ go test -v
=== RUN   TestArrangeMatrix
=== RUN   TestArrangeMatrix/testpossible.txt
2018/04/27 11:34:25 main.go:57: container 0 has different types of balls
2018/04/27 11:34:25 main.go:57: container 1 has different types of balls
=== RUN   TestArrangeMatrix/testimpossible.txt
2018/04/27 11:34:25 main.go:57: container 1 has different types of balls
2018/04/27 11:34:25 main.go:57: container 1 has different types of balls
2018/04/27 11:34:25 main.go:57: container 0 has different types of balls
2018/04/27 11:34:25 main.go:57: container 0 has different types of balls
--- PASS: TestArrangeMatrix (0.00s)
    --- PASS: TestArrangeMatrix/testpossible.txt (0.00s)
    --- PASS: TestArrangeMatrix/testimpossible.txt (0.00s)
PASS
ok  	github.com/ypapax/hackerrank/balls_containers	0.013s


```
