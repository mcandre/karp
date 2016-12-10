# karp - universal app launcher

# EXAMPLES

```
$ karp karp.gif
...

$ karp https://raw.githubusercontent.com/mcandre/karp/master/karp.txt
...

$ karp
Usage: karp <path>

-v Show version information
```

![magikarp](https://raw.githubusercontent.com/mcandre/karp/master/karp.gif)

# ABOUT

`karp` is a cross-operating system binary for easily launching file-based applications, using the skratchdot/open-golang library to route calls to `xdg-open`, etc.

`karp` will use the default application registered for a given path's protocol, mimetype, or file extension.

# DOWNLOAD

https://github.com/mcandre/karp/releases

# DEVELOPMENT REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)
* [pargs](https://github.com/mcandre/pargs)
* [editorconfig-tools](https://www.npmjs.com/package/editorconfig-tools)
* [bashate](https://github.com/openstack-dev/bashate)
* [shlint](https://rubygems.org/gems/shlint)
* [shellcheck](http://hackage.haskell.org/package/ShellCheck)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/karp/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/karp.git $GOPATH/src/github.com/mcandre/karp
$ cd $GOPATH/src/github.com/mcandre/karp
$ git submodule update --init --recursive
$ sh -c 'cd cmd/karp && go install'
```

# PORT

```
$ make port
```

# BUILD AND ARCHIVE PORTS

```
$ make port
...
Archived ports in bin/karp-0.0.1.zip
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.

# LICENSE

FreeBSD

# CREDIT

Powered by [open-golang](https://github.com/skratchdot/open-golang)
