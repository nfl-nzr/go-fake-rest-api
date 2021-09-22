package db

type Database struct {
	Data *map[string]interface{}
}

func (d *Database) Connect(dsn string) error {
	if err := d.LoadFile(dsn); err != nil {
		return err
	}
	return nil
}
