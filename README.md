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
# packagefile github.com/dhermes/go-build-remote/cmd/hello=.../go-build-remote/tmp01/2b/2b56fd4022e225d4bd4e47e954e462c99e397a2cd0901d9921015b3e0e8bf19b-d
# ...
# packagefile github.com/spf13/cobra=.../go-build-remote/tmp01/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d
# ...
# packagefile github.com/spf13/pflag=.../go-build-remote/tmp01/51/514856a8a5d0365908ed001af14f2593d836e2428284124e1f89aa70de57b5c5-d
# ...
# packagefile net=.../go-build-remote/tmp01/bb/bbcf2688e2ba626c11d9029bdd9bc4e5c0ab1503f73cdad3b4f30a77a1f65884-d
# ...
# packagefile runtime/cgo=.../go-build-remote/tmp01/c6/c6cd3ba650073f07bdbe15b710c29acbb3c7eac77a469070d51d8de2ec5871a3-d
# ...
```

Also, whoops `runtime/cgo` I definitely should have built with
`CGO_ENABLED=0`.

## Can haz `GODEBUG`?

From the docs for `GOCACHE`

> `GODEBUG=gocachehash=1` causes the `go` command to print the inputs
> for all of the content hashes it uses to construct cache lookup keys.
> The output is voluminous but can be useful for debugging the cache.

Using it:

```
TMP_GOCACHE1="$(pwd)/tmp01"
rm -fr "${TMP_GOCACHE1}" && mkdir -p "${TMP_GOCACHE1}"

GOCACHE="${TMP_GOCACHE1}" go run ./cmd/hello/ --anything goes  # Warm up first
# c = main.Config{Anything:"goes"}

GODEBUG=gocachehash=1 GOCACHE="${TMP_GOCACHE1}" go run ./cmd/hello/ --anything goes > /dev/null 2> stderr.txt
cat stderr.txt
# HASH[build internal/unsafeheader]
# HASH[build internal/unsafeheader]: "go1.17.2"
# HASH[build internal/unsafeheader]: "compile\n"
# HASH[build internal/unsafeheader]: "goos darwin goarch amd64\n"
# HASH[build internal/unsafeheader]: "import \"internal/unsafeheader\"\n"
# HASH[build internal/unsafeheader]: "omitdebug false standard true local false prefix \"\"\n"
# ...
# HASH[link github.com/dhermes/go-build-remote/cmd/hello]: "packagefile runtime/cgo=L8quj4hTlXXpseOf3wLt\n"
# HASH[link github.com/dhermes/go-build-remote/cmd/hello]: 624cc1e80dd6bc021fa4268777eb6c65faaaa4887a278d50348af933eedd583f
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "link-stdout" = 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5
```

For now, let's just see if we can find the hash that showed up before via
`git diff`

```
cat stderr.txt | grep 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "link-stdout" = 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5

cat stderr.txt | grep 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7
# HASH[build github.com/dhermes/go-build-remote/cmd/hello]: 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "stdout" = 6d4526b54d5176299b160e53e82b0e82ad741b5ca92e12b995c2ab9262c1c960
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "link-stdout" = 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5

cat stderr.txt | grep '^HASH\[link github.com/dhermes/go-build-remote/cmd/hello\]:'
# HASH[link github.com/dhermes/go-build-remote/cmd/hello]: "go1.17.2"
# ...
# HASH[link github.com/dhermes/go-build-remote/cmd/hello]: "packagefile runtime/cgo=L8quj4hTlXXpseOf3wLt\n"
# HASH[link github.com/dhermes/go-build-remote/cmd/hello]: 624cc1e80dd6bc021fa4268777eb6c65faaaa4887a278d50348af933eedd583f
```

This somewhat goes nowhere, but we can chase down the hashes in `GOCACHE`:

```
cat "${TMP_GOCACHE1}/63/630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5-a"
# v1 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5 e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855                    0  1637391075702695000

file "${TMP_GOCACHE1}/e3/e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-d"
# .../go-build-remote/tmp01/e3/e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855-d: empty

grep -r -l e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855 "${TMP_GOCACHE1}" | sort -u
# .../go-build-remote/tmp01/63/630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5-a
# .../go-build-remote/tmp01/65/6575b56bb5b1cd7df5539e4cbd49b8117483287a87a91d706c195f0c7ad1a6c2-a
# .../go-build-remote/tmp01/6d/6d4526b54d5176299b160e53e82b0e82ad741b5ca92e12b995c2ab9262c1c960-a
# .../go-build-remote/tmp01/a7/a76f57f6bdc530963467bb7d502ab7e28d1ae361a6c7de554181369290e20a6e-a
# .../go-build-remote/tmp01/c1/c1d8a401d031b93967e9ea298d94b5db10cf7f14ae843ac8d681c5dcf8802840-a
# .../go-build-remote/tmp01/f3/f3a7dfa77e01a34da25bd3538564f82d648d001b4383519eff36835f3119cd49-a

cat stderr.txt | grep \
    '\(630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5\|6575b56bb5b1cd7df5539e4cbd49b8117483287a87a91d706c195f0c7ad1a6c2\|6d4526b54d5176299b160e53e82b0e82ad741b5ca92e12b995c2ab9262c1c960\|a76f57f6bdc530963467bb7d502ab7e28d1ae361a6c7de554181369290e20a6e\|c1d8a401d031b93967e9ea298d94b5db10cf7f14ae843ac8d681c5dcf8802840\|f3a7dfa77e01a34da25bd3538564f82d648d001b4383519eff36835f3119cd49\)'
# HASH subkey fdf20d0b7a6c237e4d5fc44092ce1ac9f66a3e4b6f2df1ef30cfd68f34d4bb72 "stdout" = c1d8a401d031b93967e9ea298d94b5db10cf7f14ae843ac8d681c5dcf8802840
# HASH subkey a71032b747bf46ae0d4ee2e86c4e6a2478018fc21e6fb060b359b3dcebd9131d "stdout" = 6575b56bb5b1cd7df5539e4cbd49b8117483287a87a91d706c195f0c7ad1a6c2
# HASH subkey abc8312cb7818ec207f9289891d77e19d290667ae6a7f0c20ff550606486028b "stdout" = a76f57f6bdc530963467bb7d502ab7e28d1ae361a6c7de554181369290e20a6e
# HASH subkey 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 "stdout" = f3a7dfa77e01a34da25bd3538564f82d648d001b4383519eff36835f3119cd49
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "stdout" = 6d4526b54d5176299b160e53e82b0e82ad741b5ca92e12b995c2ab9262c1c960
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "link-stdout" = 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5

cat stderr.txt | grep \
  '\(3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702\|9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7\|a71032b747bf46ae0d4ee2e86c4e6a2478018fc21e6fb060b359b3dcebd9131d\|abc8312cb7818ec207f9289891d77e19d290667ae6a7f0c20ff550606486028b\|fdf20d0b7a6c237e4d5fc44092ce1ac9f66a3e4b6f2df1ef30cfd68f34d4bb72\)'
# HASH[build runtime/cgo]: fdf20d0b7a6c237e4d5fc44092ce1ac9f66a3e4b6f2df1ef30cfd68f34d4bb72
# HASH subkey fdf20d0b7a6c237e4d5fc44092ce1ac9f66a3e4b6f2df1ef30cfd68f34d4bb72 "stdout" = c1d8a401d031b93967e9ea298d94b5db10cf7f14ae843ac8d681c5dcf8802840
# HASH[build net]: a71032b747bf46ae0d4ee2e86c4e6a2478018fc21e6fb060b359b3dcebd9131d
# HASH subkey a71032b747bf46ae0d4ee2e86c4e6a2478018fc21e6fb060b359b3dcebd9131d "stdout" = 6575b56bb5b1cd7df5539e4cbd49b8117483287a87a91d706c195f0c7ad1a6c2
# HASH[build github.com/spf13/pflag]: abc8312cb7818ec207f9289891d77e19d290667ae6a7f0c20ff550606486028b
# HASH subkey abc8312cb7818ec207f9289891d77e19d290667ae6a7f0c20ff550606486028b "stdout" = a76f57f6bdc530963467bb7d502ab7e28d1ae361a6c7de554181369290e20a6e
# HASH[build github.com/spf13/cobra]: 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702
# HASH subkey 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 "stdout" = f3a7dfa77e01a34da25bd3538564f82d648d001b4383519eff36835f3119cd49
# HASH[build github.com/dhermes/go-build-remote/cmd/hello]: 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "stdout" = 6d4526b54d5176299b160e53e82b0e82ad741b5ca92e12b995c2ab9262c1c960
# HASH subkey 9c4088b546b50cdb94399154c0c26d91dbcae49f635513dd1d556c399abe88f7 "link-stdout" = 630a8ca60ef6ab1e5333358bd71bd59049d6767f0248d5f68a0740bfd50b94f5
```

Tracing this back, these correspond to the same 5 `packagefile ...`
packages we saw when using `go build -x`. We can further **check** that
these hashes directly correspond to the static archives for that package.
For example the `github.com/spf13/cobra` package showed up in the
`go build -x` output as:

```
# packagefile github.com/spf13/cobra=.../go-build-remote/tmp01/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d
```

and we can see that exact static archive by following

```
# HASH[build github.com/spf13/cobra]: 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702
```

within `GOCACHE`:

```
cat "${TMP_GOCACHE1}/35/3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702-a"
# v1 3535f3659d3278773cdab967c7d276a77177b38b0522bb9909d960f8f5278702 094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75              1180034  1637391071999507000

file "${TMP_GOCACHE1}/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d"
# .../go-build-remote/tmp01/09/094944797386974e2e1fe52311b04d613c3d62b5b919c9ae8725cf194e28aa75-d: current ar archive
```

[1]: https://blog.filippo.io/reproducing-go-binaries-byte-by-byte/
