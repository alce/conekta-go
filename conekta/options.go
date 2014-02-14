package conekta

import "fmt"

type FilterOptions struct {
	q string
}

func NewFilterOptions() *FilterOptions {
	return &FilterOptions{
		q: "",
	}
}

func (f *FilterOptions) Eq(field string, val interface{}) *FilterOptions {
	f.append(field, val, "=")
	return f
}

func (f *FilterOptions) Gt(field string, val interface{}) *FilterOptions {
	f.append(field, val, ".gt=")
	return f
}

func (f *FilterOptions) Gte(field string, val interface{}) *FilterOptions {
	f.append(field, val, ".gte=")
	return f
}

func (f *FilterOptions) Lt(field string, val interface{}) *FilterOptions {
	f.append(field, val, ".lt=")
	return f
}

func (f *FilterOptions) Lte(field string, val interface{}) *FilterOptions {
	f.append(field, val, ".lte=")
	return f
}

func (f *FilterOptions) Ne(field string, val interface{}) *FilterOptions {
	f.append(field, val, ".ne=")
	return f
}

func (f *FilterOptions) Regex(field string, val interface{}) *FilterOptions {
	f.append(field, val, ".regex=")
	return f
}

func (f *FilterOptions) Limit(val int) *FilterOptions {
	f.append("limit", val, "=")
	return f
}

func (f *FilterOptions) Offset(val int) *FilterOptions {
	f.append("offset", val, "=")
	return f
}

func (f *FilterOptions) Sort(field, direction string) *FilterOptions {
	f.append("sort", field+"."+direction, "=")
	return f
}

func (f FilterOptions) Q() string {
	return f.q
}

func (f *FilterOptions) append(field string, val interface{}, op string) {
	condition := fmt.Sprintf("%s%s%v", field, op, val)

	if len(f.q) == 0 {
		f.q = fmt.Sprintf("?%s", condition)
	} else {
		f.q = fmt.Sprintf("%s&%s", f.q, condition)
	}
}
