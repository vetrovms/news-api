package request

const (
	LocEn      = "en"
	LocUk      = "uk"
	DefaultLoc = LocEn
)

func LocInWhiteList(loc string) bool {
	if loc == LocEn {
		return true
	}
	if loc == LocUk {
		return true
	}
	return false
}
