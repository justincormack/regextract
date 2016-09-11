A small utility to extract files from a Docker hub image and output as a tarball.

It assumes that the files you are interested in are in the top layer. It would be
better to work out which layer they are in of course.

```
./regextract library/alpine 2>/dev/null | tar tf -
```
This will list all the files in the top layer of the alpine library image.

You can specify the files you want to extract eg
```
./regextract library/alpine bin/busybox | tar xf -
```
Will extract busybox from Alpine (note lack of leading `/`).
