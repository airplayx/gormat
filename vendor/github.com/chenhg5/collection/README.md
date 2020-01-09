# Go Collection

<a href="https://goreportcard.com/report/github.com/chenhg5/collection"><img alt="golang" src="https://img.shields.io/badge/awesome-golang-blue.svg"></a>
<a href="https://godoc.org/github.com/chenhg5/collection" rel="nofollow"><img src="https://camo.githubusercontent.com/a9a286d43bdfff9fb41b88b25b35ea8edd2634fc/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f646572656b7061726b65722f64656c76653f7374617475732e737667" alt="GoDoc" data-canonical-src="https://godoc.org/github.com/derekparker/delve?status.svg" style="max-width:100%;"></a>
<a href="https://raw.githubusercontent.com/chenhg5/collection/master/LICENSE" rel="nofollow"><img src="https://camo.githubusercontent.com/e0d5267d60ee425acfe1a1f2d6e6d92a465dcd8f/687474703a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d626c75652e737667" alt="license" data-canonical-src="http://img.shields.io/badge/license-MIT-blue.svg" style="max-width:100%;"></a>
[![Build Status](https://api.travis-ci.org/chenhg5/collection.svg?branch=master)](https://api.travis-ci.org/chenhg5/collection)

Collection provides a fluent, convenient wrapper for working with arrays of data.

You can easily convert a map or an array into a Collection with the method ```Collect()```.
And then you can use the powerful and graceful apis of Collection to process the data.

In general, Collection are immutable, meaning every Collection method returns an entirely new Collection instance

## doc

[here](https://godoc.org/github.com/chenhg5/collection#Collection)

## example

```golang
a := []int{2,3,4,5,6,7}

Collect(a).Each(func(item, value interface{}) (interface{}, bool) {
    return value.(decimal.Decimal).IntPart() + 2, false
}).ToIntArray()

// []int{4,5,6,7,8,9}

b := []map[string]interface{}{
    {"name": "Jack", "sex": 0},
    {"name": "Mary", "sex": 1},
    {"name": "Jane", "sex": 1},
}

Collect(b).Where("name", "Jack").ToMapArray()[0]

// map[string]interface{}{"name": "Jack", "sex": 0}

``` 

[more examples](https://godoc.org/github.com/chenhg5/collection#pkg-examples)

## contribution

pr is very welcome. 