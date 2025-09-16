package pagination

type sort string

func (s *sort) String() string {
	if s == nil {
		return string(ASC)
	}
	return string(*s)
}

const (
	DESC sort = "DESC"
	ASC  sort = "ASC"
)

type Pagination struct {
	Size int64 `json:"size"`
	Page int64 `json:"page"`
	Sort sort  `json:"sort"`
}

func (o *Pagination) Offset() int64 {
	return (o.Page - 1) * o.Size
}
func (o *Pagination) Limit() int64 {
	return o.Size
}
func (o *Pagination) Order() string {
	return o.Sort.String()
}
func (o *Pagination) Get() *Pagination {
	if o == nil {
		return nil
	}
	if o.Page == 0 || o.Size == 0 {
		return nil
	}
	return o
}

func New(page int64, size int64) *Pagination {
	if page == 0 || size == 0 {
		return nil
	}
	return &Pagination{
		Size: size,
		Page: page,
	}
}
