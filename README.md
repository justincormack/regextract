A small utility to extract files from a Docker hub image and output as a tarball.

```
go get -u github.com/justincormack/regextract
```

It defaults to the latest layer in the image, specify `-layer 0` for the bottom layer.
Negative numbers count from the top.

```
regextract library/alpine 2>/dev/null | tar tf -
```
This will list all the files in the top layer of the alpine library image.

You can specify the files you want to extract eg
```
regextract library/alpine bin/busybox | tar xf -
```
Will extract busybox from Alpine (note lack of leading `/`).
