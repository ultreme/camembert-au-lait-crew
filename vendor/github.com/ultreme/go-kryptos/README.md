# go-kryptos

[![Build Status](https://travis-ci.org/ultreme/go-kryptos.svg)](https://travis-ci.org/ultreme/go-kryptos)
[![Coverage Status](https://coveralls.io/repos/ultreme/go-kryptos/badge.svg?branch=master&service=github)](https://coveralls.io/github/ultreme/go-kryptos?branch=master)
[![GoDoc](https://godoc.org/github.com/ultreme/go-kryptos?status.svg)](https://godoc.org/github.com/ultreme/go-kryptos)

Encrypt/Decrypt using the awesome [Kryptos](http://www.salutcestcool.com/quatre/kryptos/) algorithm

## Usage

```command
$ kryptos hello world \!
¨¥®®© ¸©³®¤ !
```

```command
$ ➜  ~  echo hello world \! | kryptos
¨¥®®© ¸©³®¤ !
```

```command
$ echo '¨¥®®© ¸©³®¤ !' | kryptos --decrypt
hello world !
```

```command
$ kryptos --decrypt '¨¥®®© ¸©³®¤ !'
hello world !
```

```command
$ echo hello world \! | kryptos | kryptos --decrypt
hello world !
```

## API

```golang
import "github.com/ultreme/go-kryptos"
kryptos.Encrypt("Hello world !")
kryptos.Decrypt("¨¥®®© ¸©³®¤ !")
```

See [GoDoc](https://godoc.org/github.com/ultreme/go-kryptos) for more examples

## Install

```bash
go get github.com/ultreme/go-kryptos/...
```

## License

MIT
