# redis-clone
A simple Redis clone built with Go to learn about networking, concurrency, and the Redis protocol (RESP).

# Description
This project is an educational exercise to implement a subset of Redis's functionality from scratch. The primary goals are to understand how an in-memory database works, how to handle multiple client connections concurrently, and how to parse a network protocol like RESP.

This implementation is inspired by the "Build your own Redis" challenge.

# Table of Contents
Features

Prerequisites

Installation

Usage

License

Acknowledgments

# Features
This Redis clone currently supports the following commands:

[x] PING

[x] SET

[x] GET

[x] HSET

[x] HGET

# Prerequisites
Make sure you have the following tools installed on your system:

Go (version 1.18 or higher)

Git

Redis Client

# Installation
Clone the repository
```
Bash

git clone https://github.com/your-username/go-redis-clone.git
Navigate to the project directory
```
```
Bash

cd go-redis-clone
Initialize Go Modules
```

# Usage

```
# close redis default server
# on macOS
brew services stop redis

# on linux
sudo systemctl stop redis

# run all go files
go run *.go

# run the custom redis server
redis-cli

```

# Testing

```
127.0.0.1:6379> PING
PONG

127.0.0.1:6379> SET message "hello world"
OK

127.0.0.1:6379> GET message
"hello world"

127.0.0.1:6379> HSET user:1 username "alex"
(integer) 1

```

# License
Distributed under the MIT License. See LICENSE.txt for more information.

# Acknowledgments
This project was heavily inspired by the Codecrafters "Build your own Redis" challenge. 
reference: https://www.build-redis-from-scratch.dev/en/introduction
