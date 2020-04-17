package weather

type Weather interface {
	GetTempByCity(city string) (string, error)
}
