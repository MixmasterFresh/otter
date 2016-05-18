#Otter (the easy Collaboration engine)
[![Appveyor Build Status](https://ci.appveyor.com/api/projects/status/dt7in40jvcl28950/branch/master?svg=true)](https://ci.appveyor.com/project/TheAustinSeven/otter/branch/master)
[![Travis CI Build Status](https://travis-ci.org/TheAustinSeven/otter.svg?branch=master)](https://travis-ci.org/TheAustinSeven/otter)

###Note: Otter is not yet ready for use and is still in development. Any API methods currently implemented are open to change at any time.
This program is intended to make real-time collaboration easy. This is done with Operation Transformation, a technique through which operations are "transformed" according to the ones before them. This means that if we have two operations `delete(from_index, to_index)`and `insert(index, contents)`, then if we have `delete(2,3)` and `insert(3,'a')` they will be transformed according to what is run before them. In this case when `insert` is run first, `delete` is unchanged, but if `delete` is run first, then `insert` must be changed so that it inserts one earlier.

Otter is being built in Go so that it runs, as the professionals put it, "super-duper" fast. This will allow the wrappers to simply make a call to the optimized Go, and everything can run nice and fast.

##API
Currently the api is still being planned and is subject to change.

##Contributing
Right now this is a personal project, and I want to bring this to version 1 by myself, but if you are interested in building a language wrapper, or the Javascript client, for Otter, start an issue on this repository and we can start to discuss how this should be done.

## License
Otter is released under the [Apache 2.0 License](http://opensource.org/licenses/Apache-2.0).
