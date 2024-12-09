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
---
```
1_seq :  9501 | is sorted:  true
1_par :  2926 | is sorted:  true
Speedup: 3.25x
===
2_seq :  9395 | is sorted:  true
2_par :  5129 | is sorted:  true
Speedup: 1.83x
===
3_seq :  9489 | is sorted:  true
3_par :  3711 | is sorted:  true
Speedup: 2.56x
===
4_seq :  9310 | is sorted:  true
4_par :  3045 | is sorted:  true
Speedup: 3.06x
===
5_seq :  9258 | is sorted:  true
5_par :  2953 | is sorted:  true
Speedup: 3.14x
===
seq average time: 9390, par average time: 3552
Speedup: 2.64x
```

run command `go build main.go utils.go common.go seq.go par.go; ./main`