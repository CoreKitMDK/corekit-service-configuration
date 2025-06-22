package configuration

type IConfiguration interface {
	Get(key string) (string, error)
	GetFromJson(key string, structure any) error
}
