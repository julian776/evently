package models

import "context"

/**
 * Defining a new type called `HandlerFunc` which is a function that takes in a `context.Context` and a `string` parameter named `command`, and returns an `interface{}` and an `error`. This new type can be used as a function signature for any function that matches this definition.
 *
 * @typedef
 * @name HandlerFunc
 */
type HandlerFunc func(
	ctx context.Context,
	message Message,
) (interface{}, error)
