# abds-go

Agile binary data structure (ABDS) for Golang

This library is the basis for a multi repo set of golang libraries to assist in interfacing to cloud services and some golang constructs.  It has its roots in JSON but concentrates on syntactical brevity and wider type support.  JSON was conceived for universal compatibility with many tradeoffs/weaknesses.  ABDS adds developer comfort and interface simplicity while maintaining compatibility. ABDS should enable Schema-On-The-Fly.

## usage

import abds "github.com/markdiener10/abds-go"

`gabds := abds.New() //Create a new ABDS instance`

### Operations

- Create an abds instance  
`g := abds.New()`

- Add data (S() function are to create and assign values to elements)
```
g.S(1)  # Assign the first element to a value of 1 (int)
g.S("TAGA",2)  # Set the second element to a value of 2 (int) and tag it with "TAGA"
g.S("TAGB",3.1415)  # Set the third element to a value of pi (float64) and tag it with "TABG"
g.S(1,3) # Set the first element value to 3
```
- Modify Values (Pxxx() functions are pointer functions)`
```
*g.Pu(1)++ # Increase the first element to 4
*g.Pu64(1)++ # Change the first element to uint64 and increate the value to 5
*g.Pu(2)++ # Increase the second element to 3
*g.Pf32("FLOAT") += 5 #Create a tagged float32 element and increase its value from 0 to 5
```
- Get formatted values (Does not change internal type)
```
g.Vi32(2) # Retrieve the value of the second element as a int32 value
g.Vf64(2) # Retrieve the value of the second element as a float64 value
g.Vc128(2) # Retrieve the value of the second element as a complex128 value
g.Vs(2) # Retrieve the value of the second element as a string value
```
- Create a new ABDS child 
```
gch = g.Nchild("CHILDTAG")
gch.S(10)
gch.S("TAG",35)
```
- Iterate over values in the ABDS instance
```
it = g.NewIter()
for g.Iter(it) {
  it.Vs() # Get the value of the element as a string value
  *it.Pi()+=10 # Modify the element by increasing it by 10
  it.S(20) # Set the value of the element to 20 (int)
  it.S(uint(20)) # Set the value of the element to 20 (uint)

  if it.IsTag() {
    //Do something if the element is a tagged element
  }
  if it.I() > 3 {
    //Do something if the element index > 3
  }

  it.Delete() # Delete the element from the abds instance
}
```

You can use lambda/closures to scan over an ABDS instance.  Look at
abds-scan_test.go for examples on the syntax

You can pass custom structs, maps, arrays by address but not by value (cut down on allocations)
```
g.S(&mystruct)
g.S(&mymap)
g.S(&myarray)
```

- Errors

Errors in ABDS operations do not need to be constantly checked in code. The library implements
an internal error capture framework so you opt-in on checking for errors at any point in the code.  
```
g.ClearErrors() # Clear the error capture
g.IsErr() # Check for any errors since the last error clearing
g.Errs() # Get all detected errors since the last error clearing.  AbdsErr returned but capture not cleared
g.Errs(true) # Get all detected errors since the last error clearing.  AbdsErr returned and capture cleared
```

## Developer notes  (Mark Diener/TekDevo LLC)

After interfacing with a number of software development teams over the last few years, I had built up a number of code libraries that would allow me to smooth out some of the rough edges and some code pattern discipline debt.  This repo seeks to make that experience public.

This library serves as a flexible data payload container to provide a common and less brittle alternative to many common Encode/Decode, Marshal/Unmarshal, Pack/Unpack, Send/Receive, Store/ interfaces seen in Microservice architectures.

The library interface was affected by the language of implementation so every option of type and parameter flexibility was exercised to make usage simple with different usage patterns.  Sorely missing in golang were concepts like function overloading to allow for access functions to deliver pointers in a more type agnostic way.  Operator overloading would likely have been useful in some fashion.

If the number of declared type structs with marshal/unmarshal annotations drops to zero
in your project, then this library has accomplished a very basic mission.  Schema-On-The-Fly may prove to have value during initial project development as the number of changes is highest.  The number of system panics as downstream microservices should go down as ABDS allows for schema tolerance as code policy.

While this repo is specifically targeted to golang, an abds-rust repo is taking shape. Other repos are under construction that will have payload support include Dart/Flutter, PHP, and Javascript.  Maybe even WebAssembly later.

One observed pain point is that libraries like protobuf/protoc compiler take an external description file and "compile" a language specific interface for a given data structure.  This library is designed 180 degrees out of phase to that notion. We want to allow for multiple schema support with less release impedence.  This primary focus of this library is agility and pliability at the trade off of some CPU cycles.  Each language will determine the ease of use and no pre-compile steps should be necessary.

Once the base code settles out a little more and more of the other service interfaces are complete, I will loop back around and work on code optimization and profiling.  Or just purchase more nodes for the Kubernetes cluster or let AWS Fargate do its thing. Also need to add GODOC style comments.

Also missing delete and insert operations (coming soon)
Injecting tags is also pending.