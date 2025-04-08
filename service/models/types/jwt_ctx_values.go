package types

type JWTContextValues struct {
	M map[string]any
}

func (j JWTContextValues) Get(key string) any {
	return j.M[key]
}
