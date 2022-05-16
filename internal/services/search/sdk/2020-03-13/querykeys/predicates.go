package querykeys

type QueryKeyOperationPredicate struct {
	Key  *string
	Name *string
}

func (p QueryKeyOperationPredicate) Matches(input QueryKey) bool {

	if p.Key != nil && (input.Key == nil && *p.Key != *input.Key) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
