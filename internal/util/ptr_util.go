package util

func ToString(p *string) (v string) {
	if p == nil {
		return v
	}

	return *p
}
