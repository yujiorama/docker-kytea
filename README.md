# docker-kytea

container image for [Kytea](https://github.com/neubig/kytea).

[hub.docker.com](https://hub.docker.com/r/yujiorama/docker-kytea)

## Build

```bash
$ docker build -t docker-kytea -f Dockerfile .
$ docker images | grep docker-kytea
docker-kytea                      latest           5b8bab0e0388   5 seconds ago       202MB
```

## Run

see [KyTea](http://www.phontron.com/kytea/index-ja.html).

### from stdin.

```bash
$ echo コーパスの文です。 | docker run --rm -i yujiorama/docker-kytea kytea
コーパス/名詞/こーぱす の/助詞/の 文/名詞/ぶん で/助動詞/で す/語尾/す 。/補助記号/。
```

### from text file within volume mount directory.

```bash
$ echo コーパスの文です。 > test.raw
$ docker run --rm -i -v $(pwd):/app yujiorama/docker-kytea kytea /app/test.raw
コーパス/名詞/こーぱす の/助詞/の 文/名詞/ぶん で/助動詞/で す/語尾/す 。/補助記号/。
```

## License

[MIT](./LICENSE)
