# Roadmap

## Overview

- Serves mounted directory which is either local or a remote volume

- Security: Prevent API callers from operating anything out of the mounted directory

- Provides both RESTful APIs and GraphQL APIs

- Provides Docker images and binaries which can be run out of the box (CLI args for convenient use of binary file)

- There will be a frontend app for simply using of `rmdashrf`. The app is written with [React.js](https://reactjs.org/) or [Angular](https://angular.io/). (_Angular_ is preferred)

## Features plan

### P0

- Create file, directory

- Remove file, directory

- Upload file, directory

- Download file, directory as a tar archived package

### P1

- Preview file

- Trashcan

- gRPC

### Future plan

- multipul mouted volumes for different accounts