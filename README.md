# ccwc

Coding challenge

[https://codingchallenges.fyi/challenges/challenge-wc/](https://codingchallenges.fyi/challenges/challenge-wc/)

From today's [LinkedIn post](https://www.linkedin.com/feed/update/urn:li:activity:7126124184690139138/) by the Coding Challenge owner.

Not sure if it is fair (no explicit denial in [the intro](https://codingchallenges.fyi/challenges/intro/)), AI ChatGPT assisted:

* [Main conversation](https://chat.openai.com/share/e7d174a3-3ac3-4e45-bcc3-58a1d35ee937)
  * [Sidecar conversation 1](https://chat.openai.com/share/5941b48c-ea88-4068-9037-1451fbfaf74b)
  * [Sidecar converstaion 2
  ](https://chat.openai.com/share/ebbd21c4-922b-4ef5-98ee-5fcd7985d054)
  * [Sidecar converstaion 3
  ](https://chat.openai.com/share/34fef07d-5b15-4a8a-91f9-1f6371c5419f)
  * [Sidecar converstaion 4
  ](https://chat.openai.com/share/71a92e8b-89de-484f-9984-5d93fda1e872)

The steps are not in order.

In Golang that I am actively practicing now.

## Prepare for test

```
wget https://www.gutenberg.org/cache/epub/132/pg132.txt -O test.txt
```

```
--2023-11-03 11:47:18--  https://www.gutenberg.org/cache/epub/132/pg132.txt
Resolving www.gutenberg.org (www.gutenberg.org)... 152.19.134.47, 2610:28:3090:3000:0:bad:cafe:47
Connecting to www.gutenberg.org (www.gutenberg.org)|152.19.134.47|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 342190 (334K) [text/plain]
Saving to: ‘test.txt’

test.txt                          100%[============================================================>] 334.17K   490KB/s    in 0.7s    

2023-11-03 11:47:20 (490 KB/s) - ‘test.txt’ saved [342190/342190]
```

Gold output:

```sh
cat test.txt | wc
cat test.txt | wc -lwc
cat test.txt | wc -m
cat test.txt | wc -lwcm
```

```log
   7145   58164  342190
   7145   58164  342190
339292
   7145   58164  339292  342190
```

## Build

```sh
go build ccwc.go
```

## Run

### Build & test

```
go build ccwc.go
./ccwc -c -w -l -m test.txt
cat test.txt | ./ccwc -c -w -l -m
cat test.txt | wc -lwcm
```

```
   7145   58164  339292  342190 test.txt
   7145   58164  339292  342190 
   7145   58164  339292  34219
```

### Flags:

```sh
./ccwc -c test.txt
./ccwc -l test.txt
./ccwc -w test.txt
./ccwc -m test.txt
```

```log
  342190 test.txt
   7145 test.txt
   58164 test.txt
  339292 test.txt
```

### Filename input:

```sh
./ccwc test.txt
```

```
   7145   58164  342190 test.txt
```

### Piped input:

```sh
cat test.txt | ./ccwc
```

```
   7145   58164  342190
```
