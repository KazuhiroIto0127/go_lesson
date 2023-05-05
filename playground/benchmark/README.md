ベンチマークを実行するには下記。
ポインターのほうがわずかにパフォーマンスが良いことがわかる。

```shell
root@18aaf90768c6:/app/benchmark# go test -bench=.
goos: linux
goarch: amd64
pkg: pointer.go
cpu: 13th Gen Intel(R) Core(TM) i5-13600KF
BenchmarkModifyStructByValue-20         1000000000               0.1366 ns/op
BenchmarkModifyStructByPointer-20       1000000000               0.1237 ns/op
```
