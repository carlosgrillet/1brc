# 1 billion rows challenge in Go

A few days ago I discovered this challenge by Gunnar Morling [original](https://github.com/gunnarmorling/1brc)
and I decided to give it a try.

At first glance it may look like the challenge is just to read avery big file fast
but once you start woriking on it, you realize that you have to have a well versed
understanding of the programing language that you are using and how the system works
to achieve faster times.

On my first attempt this was my time. Very simple code, no fancy stuff, without concurrency

```
$ time ./onebrc
./onebrc  140.11s user 3.62s system 128% cpu 1:51.94 total
```
