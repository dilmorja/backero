package types

type String_t string

func (String_t) Str() string { return "string" }

const String String_t = ""