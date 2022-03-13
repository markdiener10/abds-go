# abds-go

Agile binary data structure for Golang

This library is the basis for a multi repo set of golang libraries to assist in interfacing to cloud services and some golang constructs.  It has characteristics of JSON but has the ability for stronger type control and extensibility

## usage

import "github.com/markdiener10/abds-go"

go := abds.New() //Initialize
ga := abds.New(ARRAY)

## developer notes

After interfacing with a number of software development teams over the last few years, I had built up a number of code libraries that would allow me to smooth out some of the rough edges and outright code pattern discipline debt.  This repo seeks to make that experience public to share with other devs on other teams.

This library serves as a flexible data payload container to provide a common and less brittle alternative to many common Encode/Decode, Marshal/Unmarshal, Pack/Unpack, Send/Receive, Store/ interfaces seen in Microservice architectures.

While the payload container is specifically targeted to golang, other repos are under construction that will have matching payload container support for Dart/Flutter, Rust, and Javascript.  Maybe even WebAssembly later.

One observed pain point is that of using protobuf and the protoc compiler to take an external description file and "compile" a language specific interface for a given data structure.  This library is designed 180 degrees out of phase to that notion. We want to allow for multiple version support with lower release impedence and for rolling upgrades to absorb and process version mismatches and to survive those nightly build/release events.

Once the base code settles out a little more and more of the other service interfaces are complete, I will loop back around and work on code optimization and performance along with adjustements for harmonizing developer usage across golang,dart,rust and javascript. (Maybe PHP and webassembly later)

