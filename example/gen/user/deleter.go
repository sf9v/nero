// Code generated by nero, DO NOT EDIT.
package user

type Deleter struct {
	collection string
	pfs        []PredicateFunc
}

func NewDeleter() *Deleter {
	return &Deleter{
		collection: collection,
	}
}

func (d *Deleter) Where(pfs ...PredicateFunc) *Deleter {
	d.pfs = append(d.pfs, pfs...)
	return d
}
