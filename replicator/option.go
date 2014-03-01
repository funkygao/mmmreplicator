package replicator

type Options struct {
	After  TimestampGenerator
	Filter OpFilter
}
