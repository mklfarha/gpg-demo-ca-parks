package park

//go:generate go run github.com/dmarkham/enumer -type=LinksType -json
type LinksType int

const (
	LINKS_TYPE_INVALID LinksType = iota
	LINKS_TYPE_BROCHURE
	LINKS_TYPE_MAP
	LINKS_TYPE_INSTAGRAM
	LINKS_TYPE_FACEBOOK
	LINKS_TYPE_YOUTUBE
	LINKS_TYPE_RESERVATION
	LINKS_TYPE_CHARGERS
	LINKS_TYPE_EBIKES
)

func (e LinksType) ToInt32() int32 {
	return int32(e)
}

func LinksTypeFromString(in string) LinksType {
	switch in {
	case "brochure":
		return LINKS_TYPE_BROCHURE
	case "map":
		return LINKS_TYPE_MAP
	case "instagram":
		return LINKS_TYPE_INSTAGRAM
	case "facebook":
		return LINKS_TYPE_FACEBOOK
	case "youtube":
		return LINKS_TYPE_YOUTUBE
	case "reservation":
		return LINKS_TYPE_RESERVATION
	case "chargers":
		return LINKS_TYPE_CHARGERS
	case "ebikes":
		return LINKS_TYPE_EBIKES
	}
	return LINKS_TYPE_INVALID
}

func LinksTypeFromPointerString(in *string) LinksType {
	if in == nil {
		return LINKS_TYPE_INVALID
	}
	return LinksTypeFromString(*in)
}

func (e LinksType) String() string {
	switch e {
	case LINKS_TYPE_BROCHURE:
		return "brochure"
	case LINKS_TYPE_MAP:
		return "map"
	case LINKS_TYPE_INSTAGRAM:
		return "instagram"
	case LINKS_TYPE_FACEBOOK:
		return "facebook"
	case LINKS_TYPE_YOUTUBE:
		return "youtube"
	case LINKS_TYPE_RESERVATION:
		return "reservation"
	case LINKS_TYPE_CHARGERS:
		return "chargers"
	case LINKS_TYPE_EBIKES:
		return "ebikes"
	}

	return "invalid"
}

func (e LinksType) StringPtr() *string {
	val := e.String()
	return &val
}
