package user

//go:generate go run github.com/dmarkham/enumer -type=Status -json
type Status int

const (
	STATUS_INVALID Status = iota
	STATUS_ENABLED
	STATUS_DISABLED
)

func (e Status) ToInt32() int32 {
	return int32(e)
}

func StatusFromString(in string) Status {
	switch in {
	case "enabled":
		return STATUS_ENABLED
	case "disabled":
		return STATUS_DISABLED
	}
	return STATUS_INVALID
}

func StatusFromPointerString(in *string) Status {
	if in == nil {
		return STATUS_INVALID
	}
	return StatusFromString(*in)
}

func (e Status) String() string {
	switch e {
	case STATUS_ENABLED:
		return "enabled"
	case STATUS_DISABLED:
		return "disabled"
	}

	return "invalid"
}

func (e Status) StringPtr() *string {
	val := e.String()
	return &val
}
