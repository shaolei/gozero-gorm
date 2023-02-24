
type (
	{{.lowerStartCamelObject}}Model interface{
		{{.method}}
	}

	default{{.upperStartCamelObject}}Model struct {
		{{if .withCache}}zgorm.CachedConn{{else}}conn *gorm.DB{{end}}
		table string
	}

	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)
