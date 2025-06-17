package util

func PointerString(s string) *string    { return &s }
func PointerInt64(i int64) *int64       { return &i }
func PointerFloat64(i float64) *float64 { return &i }
func PointerInt(i int) *int             { return &i }
func PointerBool(b bool) *bool          { return &b }
