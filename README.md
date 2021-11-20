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

## Try Again, With `SOURCE_DATE_EPOCH`

```
TMP_GOCACHE1="$(pwd)/tmp01"
rm -fr "${TMP_GOCACHE1}" && mkdir -p "${TMP_GOCACHE1}"
SOURCE_DATE_EPOCH=0 GOCACHE="${TMP_GOCACHE1}" go run ./cmd/hello/ --anything goes
# c = main.Config{Anything:"goes"}

TMP_GOCACHE2="$(pwd)/tmp02"
rm -fr "${TMP_GOCACHE2}" && mkdir -p "${TMP_GOCACHE2}"
SOURCE_DATE_EPOCH=0 GOCACHE="${TMP_GOCACHE2}" go run ./cmd/hello/ --anything goes
# c = main.Config{Anything:"goes"}

diff "${TMP_GOCACHE1}" "${TMP_GOCACHE2}"
# diff -Nru .../go-build-remote/tmp01/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a .../go-build-remote/tmp02/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a
# --- .../go-build-remote/tmp01/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a 2021-11-19 23:57:32.000000000 -0600
# +++ .../go-build-remote/tmp02/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a 2021-11-19 23:57:53.000000000 -0600
# @@ -1 +1 @@
# -v1 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75              1180034  1637387852815852000
# +v1 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75              1180034  1637387873696378000
# diff -Nru .../go-build-remote/tmp01/3d/3d1130644e3e25a14e87b8f923426da26bcb533aa9d00cad50767e9f37cebad9-a .../go-build-remote/tmp02/3d/3d1130644e3e25a14e87b8f923426da26bcb533aa9d00cad50767e9f37cebad9-a
# --- .../go-build-remote/tmp01/3d/3d1130644e3e25a14e87b8f923426da26bcb533aa9d00cad50767e9f37cebad9-a 2021-11-19 23:57:32.000000000 -0600
# +++ .../go-build-remote/tmp02/3d/3d1130644e3e25a14e87b8f923426da26bcb533aa9d00cad50767e9f37cebad9-a 2021-11-19 23:57:53.000000000 -0600
# ...
```

OK but what is **actually** different. Looking at the files in `GOCACHE`,
most of them are just text files which point at **other** files in the
`GOCACHE` (magic secrets!):

```
find "${TMP_GOCACHE1}" -type f | sort -u | xargs file
# .../go-build-remote/tmp01/07/075f0acb7f7ab2048b3a40344e8bc945af0bfc921e5dc2ffc5da463319baa3ec-d: c program text, ASCII text
# .../go-build-remote/tmp01/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d: current ar archive
# .../go-build-remote/tmp01/0e/0e4c02c3e4c7616da3b53532a5ddd22048fc34c721a6a7052f2e2d64fae56984-d: ASCII text
# .../go-build-remote/tmp01/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d: ASCII text
# .../go-build-remote/tmp01/2b/2b56fd4022e225d4bd4e47e954e462c99e397a2cd0901d9921015b3e0e8bf19b-d: current ar archive
# ...

cat .../go-build-remote/tmp01/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d
# ./addrselect.go
# ./conf.go
# ./dial.go
# ./dnsclient.go
# ...
```

However **most** of the text files **don't** differ and none of the static
archives differ:

```
diff --report-identical-files \
  "${TMP_GOCACHE1}/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d" \
  "${TMP_GOCACHE2}/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d"
# Files .../go-build-remote/tmp01/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d and .../go-build-remote/tmp02/1e/1e28f1f995cdba394c4c1ce40336db2bdaf6c469eda948c4c3ac0775a2da0c7c-d are identical

diff --report-identical-files \
  "${TMP_GOCACHE1}/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d" \
  "${TMP_GOCACHE2}/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d"
Files .../go-build-remote/tmp01/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d and .../go-build-remote/tmp02/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d are identical
```

## Good Enough?

Filippo Valsorda has a wonder [post][1] on byte-for-byte identical **binaries**
but that doesn't have the same bearing on `GOCACHE`. Let's quickly check if
the volatile information (timestamps in ASCII text files) **changes** on
repeated build invocations

```
TMP_GOCACHE3="$(pwd)/tmp03"
rm -fr "${TMP_GOCACHE3}" && mkdir -p "${TMP_GOCACHE3}"
GOCACHE="${TMP_GOCACHE3}" go run ./cmd/hello/ --anything goes
# c = main.Config{Anything:"goes"}

git add "${TMP_GOCACHE3}"
git status
# On branch main
# Your branch is up to date with 'origin/main'.
#
# Changes to be committed:
#   (use "git restore --staged <file>..." to unstage)
#         new file:   tmp03/07/075f0acb7f7ab2048b3a40344e8bc945af0bfc921e5dc2ffc5da463319baa3ec-d
#         new file:   tmp03/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d
# ...
#         new file:   tmp03/fd/fdf20d0b7a6c237e4d5fc44092ce1ac9f66a3e4b6f2df1ef30cfd68f34d4bb72-a
#         new file:   tmp03/trim.txt
#

GOCACHE="${TMP_GOCACHE3}" go run ./cmd/hello/ --anything goes
# c = main.Config{Anything:"goes"}

git diff
# diff --git a/tmp03/63/630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5-a b/tmp03/63/630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5-a
# index 55413ba..16ea6de 100644
# --- a/tmp03/63/630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5-a
# +++ b/tmp03/63/630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5-a
# @@ -1 +1 @@
# -v1 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5 e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855                    0  1637389232498829000
# +v1 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5 e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855                    0  1637389324568857000

git rm -fr "${TMP_GOCACHE3}"
```

Building with `-x` gives us a little bit of insight to what these hashes
correspond to

```
TMP_GOCACHE1="$(pwd)/tmp01"
rm -fr "${TMP_GOCACHE1}" && mkdir -p "${TMP_GOCACHE1}"
GOCACHE="${TMP_GOCACHE1}" go run -x ./cmd/hello/ --anything goes
# WORK=/var/folders/61/6rn9bhys12l9qngs88ygnp7c0000gp/T/go-build2639010854
# ...
# packagefile github.com/dhermes/go-build-remote/cmd/hello=.../go-build-remote/tmp02/2b/2b56fd4022e225d4bd4e47e954e462c99e397a2cd0901d9921015b3e0e8bf19b-d
# ...
# packagefile github.com/spf13/cobra=.../go-build-remote/tmp02/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d
# ...
# packagefile github.com/spf13/pflag=.../go-build-remote/tmp02/51/514856a8a5d0365908ed001af14f2593d836e2428284124e1f89aa70de57b5c5-d
# ...
# packagefile net=.../go-build-remote/tmp02/bb/bbcf2688e2ba626c11d9029bdd9bc4e5c0ab1503f73cdad3b4f30a77a1f65884-d
# ...
# packagefile runtime/cgo=.../go-build-remote/tmp02/c6/c6cd3ba650073f07bdbe15b710c29acbb3c7eac77a469070d51d8de2ec5871a3-d
# ...
```

Also, whoops `runtime/cgo` I definitely should have built with
`CGO_ENABLED=0`.

[1]: https://blog.filippo.io/reproducing-go-binaries-byte-by-byte/
