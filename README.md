duck [![GoDoc](https://godoc.org/github.com/go-zoo/duck?status.png)](http://godoc.org/github.com/go-zoo/claw) [![Build Status](https://travis-ci.org/go-zoo/duck.svg?branch=master)](https://travis-ci.org/go-zoo/duck)
=======

## What is duck ?

Duck pass context values to all handler for a request,
and clear the context after the request ends.

![alt tag](http://i.dailymail.co.uk/i/pix/2007/06_03/WimbDuck2R2906_468x287.jpg)

## Features

- Use `SetContext()` and `GetContext()`.
- Provide a Middleware to delete context when request ends.

## Example
```go
package main

import (
    "github.com/go-zoo/duck"
)

func main() {
    mux := http.NewServeMux()

    mux.Handle("/", TestHandler)

	http.ListenAndServe(":8080", duck.UseContext(mux))
}

func TestHandler(rw http.ResponseWriter, req *http.Request) {
    duck.SetContext(req, "id", "test")
    rw.Write([]byte(duck.GetContext(req, "id")))
}

```

## TODO
- DOC
- Refactoring
- Debugging

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

## License
MIT

## Links

Lightning Fast HTTP Mux : [Bone](https://github.com/go-zoo/bone)
