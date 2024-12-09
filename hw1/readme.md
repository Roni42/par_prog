common - common sort parts \
seq - sequentian version \
par - parallel version \
utils - helpers \
main - tests \

---
```
results:
1:  9421 true
2:  3117 true
Speedup: 3.02x
===
1:  9418 true
2:  4688 true
Speedup: 2.01x
===
1:  9371 true
2:  4520 true
Speedup: 2.07x
===
1:  9467 true
2:  3085 true
Speedup: 3.07x
===
1:  9330 true
2:  3047 true
Speedup: 3.06x
===
seq average time: 9401, par average time: 3691
Speedup: 2.55x
```

run command `go build main.go utils.go common.go seq.go par.go; ./main`