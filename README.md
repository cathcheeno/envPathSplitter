TODO: set build status of CircleCI

# envPathSplitter

## Overview
`envPathSplitter` is a CLI tool to list up env paths with split style.

## Installation
Make sure your Go environment is properly setup.
If not setup yet, see the [install instruction](http://golang.org/doc/install.html).

To install `envPathSplitter`, run the following command.
```
$ go get github.com/cathcheeno/envPathSplitter
```

## Usage
This command will list up all envs related to path.
```
$ eps
```

You can specify the env name like this.
```
$ eps --path PATH
```

Also you can specify the path splitter.
```
$ eps --splitter :
```

For more detail, refer to the help.
```
$ eps --help
```

## Limitation
`envPathSplitter` is tested with only an UNIX-Based system. So `envPathSplitter` doesn't work with other systems.

## LICENSE
This software is released under the MIT License, see [LICENSE](./LICENSE).