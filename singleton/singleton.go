package singleton

// singleton is the structure that will keep our *shared* state
type singleton struct {
	count int
}

// instance is the PRIVATE (non-exposed, first letter lowercase) so that no one
// outside this package can access it
var instance *singleton

// GetInstance has the sole purpose of exposing the private instance in our own
// terms: in this case, all callers get the one SINGLEton instance we want.
// In other words, as long as there's already one, every caller gets that same
// one.
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

// AddOne is the way to interact with the PRIVATE `count` field of the singleton
// instance.
func (s *singleton) AddOne() int {
	s.count ++
	return s.count
}
