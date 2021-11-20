# `go-build-remote`

> Toying With Strategies for Local / Remote Go Build Caching

## Regular Build

```
$ go run ./cmd/hello/ --anything goes
c = main.Config{Anything:"goes"}
```

## Regular Build

```
TMP_GOCACHE1="$(pwd)/tmp01"
rm -fr "${TMP_GOCACHE1}"
mkdir -p "${TMP_GOCACHE1}"
ls -alFG "${TMP_GOCACHE1}"
# total 0
# drwxr-xr-x   2 dhermes  staff   64 Nov 19 23:32 ./
# drwxr-xr-x  11 dhermes  staff  352 Nov 19 23:32 ../

GOCACHE="${TMP_GOCACHE1}" go run ./cmd/hello/ --anything goes
# c = main.Config{Anything:"goes"}
ls -alFG "${TMP_GOCACHE1}"
# total 16
# drwxr-xr-x  260 dhermes  staff  8320 Nov 19 23:33 ./
# drwxr-xr-x   11 dhermes  staff   352 Nov 19 23:32 ../
# drwxr-xr-x    2 dhermes  staff    64 Nov 19 23:33 00/
# drwxr-xr-x    2 dhermes  staff    64 Nov 19 23:33 01/
# ...

find "${TMP_GOCACHE1}" -type f | sort -u
# .../go-build-remote/tmp01/07/075f0acb7f7ab2048b3a40344e8bc945af0bfc921e5dc2ffc5da463319baa3ec-d
# .../go-build-remote/tmp01/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d
# .../go-build-remote/tmp01/0e/0e4c02c3e4c7616da3b53532a5ddd22048fc34c721a6a7052f2e2d64fae56984-d
# .../go-build-remote/tmp01/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d
# ...
# .../go-build-remote/tmp01/f8/f8697f7cd1d9acbe676e15b587434e55bdf4d5707d9293ffde0bebf0b821f765-a
# .../go-build-remote/tmp01/fa/fa34456d1037fa907ad03253728665646ec110da5ed5650b7f36f73b5828e102-d
# .../go-build-remote/tmp01/fd/fdf20d0b7a6c237e4d5fc44092ce1ac9f66a3e4b6f2df1ef30cfd68f34d4bb72-a
# .../go-build-remote/tmp01/trim.txt

TMP_GOCACHE2="$(pwd)/tmp02"
rm -fr "${TMP_GOCACHE2}" && mkdir -p "${TMP_GOCACHE2}"
GOCACHE="${TMP_GOCACHE2}" go run ./cmd/hello/ --anything goes
# c = main.Config{Anything:"goes"}

diff "${TMP_GOCACHE1}" "${TMP_GOCACHE2}"
# diff -Nru .../go-build-remote/tmp01/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a .../go-build-remote/tmp02/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a
# --- .../go-build-remote/tmp01/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a 2021-11-19 23:33:30.000000000 -0600
# +++ .../go-build-remote/tmp02/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a 2021-11-19 23:37:43.000000000 -0600
# @@ -1 +1 @@
# -v1 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75              1180034  1637386410460111000
# +v1 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75              1180034  1637386663261969000
# diff -Nru .../go-build-remote/tmp01/3d/3d1130644e3e25a14e87b8f923426da26bcb533aa9d00cad50767e9f37cebad9-a .../go-build-remote/tmp02/3d/3d1130644e3e25a14e87b8f923426da26bcb533aa9d00cad50767e9f37cebad9-a
# ...
```
