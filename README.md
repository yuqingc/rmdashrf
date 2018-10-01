# rmdashrf (WIP)

**RMDASHRF** serves mounted directory with highly security and efficiently

Table of Contents
=================

<!--ts-->
   * [rmdashrf (WIP)](#rmdashrf-wip)
   * [Table of Contents](#table-of-contents)
      * [Get Started](#get-started)
         * [Start by running binaries](#start-by-running-binaries)
         * [Run a Docker image](#run-a-docker-image)
         * [Build from source](#build-from-source)
      * [API Documentation](#api-documentation)
      * [Roadmap](#roadmap)
      * [Contributing](#contributing)
         * [Tools](#tools)
         * [Dependencies](#dependencies)
         * [Debug](#debug)

<!-- Added by: matt, at:  -->

<!--te-->

## Get Started

### Start by running binaries

1. Download binaries in the [release page](https://github.com/yuqingc/rmdashrf/releases)

2. Start

```
$ ./rmdashrf -port=8080 -volume=/data
```

The `-port` flag is optional, `8080` by default. Use `./rmdashrf -h` or `./rmdashrf -help` for more information

### Run a Docker image

TODO

### Build from source

TODO

## API Documentation

- [apis](https://github.com/yuqingc/rmdashrf/blob/master/docs/apis.md)

## Roadmap

- [roadmap](https://github.com/yuqingc/rmdashrf/blob/master/docs/roadmap.md)


## Contributing

### Tools

|Name|Version|
|-|-|
|go|1.11|
|dep|0.5|
|docker-ce|18.03|


### Dependencies

```
$ dep ensure
```

### Debug

- The entry main package is in `cmd/rmdashrf`
