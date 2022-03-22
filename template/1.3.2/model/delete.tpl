
func (m *default{{.upperStartCamelObject}}Model) Delete({{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne({{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}{{end}}

	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.Exec(func(conn *gorm.DB) *gorm.DB {
        return conn.Delete(&{{.upperStartCamelObject}}{}, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		_,err:=m.CachedConn.ExecNoCache(func(conn *gorm.DB) *gorm.DB {
		    return conn.Delete(&{{.upperStartCamelObject}}{}, {{.lowerStartCamelPrimaryKey}}){{end}}
		}
	return err
}