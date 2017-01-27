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

Binary artifacts for amd64 architectures are built by CI:

- [MacOS](https://circleci.com/api/v1/project/justincormack/regextract/latest/artifacts/0/$CIRCLE_ARTIFACTS/darwin/amd64/regextract)
- [Linux](https://circleci.com/api/v1/project/justincormack/regextract/latest/artifacts/0/$CIRCLE_ARTIFACTS/linux/amd64/regextract)
- [Windows](https://circleci.com/api/v1/project/justincormack/regextract/latest/artifacts/0/$CIRCLE_ARTIFACTS/windows/amd64/regextract.exe)
