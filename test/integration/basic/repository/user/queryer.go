// Code generated by nero, DO NOT EDIT.
package user

type Queryer struct {
	collection string
	columns    []string
	limit      uint64
	offset     uint64
	pfs        []PredFunc
	sfs        []SortFunc
}

func NewQueryer() *Queryer {
	return &Queryer{
		collection: collection,
		columns:    []string{"id", "uid", "email", "name", "age", "group_res", "kv", "updated_at", "created_at"},
	}
}

func (q *Queryer) Where(pfs ...PredFunc) *Queryer {
	q.pfs = append(q.pfs, pfs...)
	return q
}

func (q *Queryer) Sort(sfs ...SortFunc) *Queryer {
	q.sfs = append(q.sfs, sfs...)
	return q
}

func (q *Queryer) Limit(limit uint64) *Queryer {
	q.limit = limit
	return q
}

func (q *Queryer) Offset(offset uint64) *Queryer {
	q.offset = offset
	return q
}
