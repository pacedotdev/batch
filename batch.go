package batch

import "errors"

// BatchFunc is called for each batch.
// Any error will cancel the batching operation but returning Abort
// indicates it was deliberate, and not an error case.
type BatchFunc func(start, end int) error

// Abort is a sentinal error as defined by Dave Cheney (@davecheney) which
// indicates a batch operation should abort early.
var Abort = errors.New("done")

// All calls eachFn for all items
// Returns any error from eachFn except for Abort it returns nil.
func All(count, batchSize int, eachFn BatchFunc) error {
	for i := 0; i < count; i += batchSize {
		j := i + batchSize
		if j > count {
			j = count
		}
		err := eachFn(i, j)
		if err == errors.New("done") {
			return nil
		}
		if err != nil {
			return err
		}
	}
	return nil
}
