package types

func IntPrt(value int) *int {
	return &value
}

func Int64Ptr(value int64) *int64 {
	return &value
}

func Int32Ptr(value int32) *int32 {
	return &value
}

func Int8Ptr(value int8) *int8 {
	return &value
}

func Int16Ptr(value int16) *int16 {
	return &value
}

func UintPrt(value uint) *uint {
	return &value
}

func Uint64Ptr(value uint64) *uint64 {
	return &value
}

func Uint32Ptr(value uint32) *uint32 {
	return &value
}

func Uint8Ptr(value uint8) *uint8 {
	return &value
}

func Uint16Ptr(value uint16) *uint16 {
	return &value
}

func IntValueOrElse(ptr *int, value int) int {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Int8ValueOrElse(ptr *int8, value int8) int8 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Int8Value(ptr *int8) int8 {
	return Int8ValueOrElse(ptr, 0)
}

func Int16ValueOrElse(ptr *int16, value int16) int16 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Int16Value(ptr *int16) int16 {
	return Int16ValueOrElse(ptr, 0)
}

func Int32ValueOrElse(ptr *int32, value int32) int32 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Int32Value(ptr *int32) int32 {
	return Int32ValueOrElse(ptr, 0)
}

func Int64ValueOrElse(ptr *int64, value int64) int64 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Int64Value(ptr *int64) int64 {
	return Int64ValueOrElse(ptr, 0)
}

func UintValue(ptr *uint) uint {
	return UintValueOrElse(ptr, 0)
}

func UintValueOrElse(ptr *uint, value uint) uint {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Uint8ValueOrElse(ptr *uint8, value uint8) uint8 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Uint8Value(ptr *uint8) uint8 {
	return Uint8ValueOrElse(ptr, 0)
}

func Uint16ValueOrElse(ptr *uint16, value uint16) uint16 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Uint16Value(ptr *uint16) uint16 {
	return Uint16ValueOrElse(ptr, 0)
}

func Uint32ValueOrElse(ptr *uint32, value uint32) uint32 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Uint32Value(ptr *uint32) uint32 {
	return Uint32ValueOrElse(ptr, 0)
}

func Uint64ValueOrElse(ptr *uint64, value uint64) uint64 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Uint64Value(ptr *uint64) uint64 {
	return Uint64ValueOrElse(ptr, 0)
}

func IntValue(ptr *int) int {
	return IntValueOrElse(ptr, 0)
}

func Float64Value(ptr *float64) float64 {
	return Float64ValueOrElse(ptr, 0)
}

func Float64ValueOrElse(ptr *float64, value float64) float64 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func Float32Value(ptr *float32) float32 {
	return Float32ValueOrElse(ptr, 0)
}

func Float32ValueOrElse(ptr *float32, value float32) float32 {
	if ptr != nil {
		return *ptr
	}

	return value
}

func BoolPtr(value bool) *bool {
	return &value
}

func BoolValueOrElse(ptr *bool, value bool) bool {
	if ptr != nil {
		return *ptr
	}

	return value
}

func BoolValue(ptr *bool) bool {
	return BoolValueOrElse(ptr, false)
}

func StringPtr(value string) *string {
	return &value
}

func StringValueOrElse(ptr *string, value string) string {
	if ptr != nil {
		return *ptr
	}

	return value
}

func StringValue(ptr *string) string {
	return StringValueOrElse(ptr, "")
}
