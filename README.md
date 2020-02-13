# batch

Very simple batching API for Go.

## Install

You might as well just copy the [batch.go](https://github.com/matryer/batch/blob/master/batch.go) file into your own project (and the [batch_test.go](https://github.com/matryer/batch/blob/master/batch_test.go) while you're at it for future generations) rather than adding a dependency.

But it is maintained as a Go module which you can get with:

```bash
go get github.com/matryer/batch
```

## Usage

Import it:

```go
import (
	"github.com/matryer/batch"
)
```

If we wanted to perform an operation in batches of ten, regardless of how many items were in the slice, we could use the `batch.All` function by passing in the total number of items `len(items)`, and the `batchSize`. The function is a callback that gets called for each batch, with `start` and `end` indexes which you can use to re-slice the `items`:

```go
items, err := getAllItemsFromRequest(r)
if err != nil {
	return errors.Wrap(err, "getAllItemsFromRequest")
}
batchSize := 10
err := batch.All(len(items), batchSize, func(start, end int) error {
	batchItems := items[start:end]
	if err := performSomeRemoteThing(ctx, batchItems); err != nil {
		return errors.Wrap(err, "performSomeRemoteThing")
	}
})
if err != nil {
	return err
}
```

In this example, if we got `105` items, the `performSomeRemoteThing` function would get called eleven times, each time with a different page of `10` items (the `batchSize`) except the last time, when it would be a slice of the remaining five items.

The mechanics are [fairly simple](https://github.com/matryer/batch/blob/master/batch.go), but the code is encapsulated and well tested.
