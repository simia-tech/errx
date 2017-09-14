package errx

type locationer interface {
	Location() (string, int)
}
