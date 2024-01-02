package interface_demo

// DB 数据接口
type DB interface {
	Get(key string) (int, error)
	Add(key string, value int) error
}

func NewDB(name string) DB {
	if name == "my" {
		return new(MyDB)
	}
	return nil
}

type MyDB struct {
}

func (m *MyDB) Get(key string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MyDB) Add(key string, value int) error {
	//TODO implement me
	panic("implement me")
}

// GetFromDB 根据key从DB查询数据的函数
func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key); err == nil {
		return v
	}
	return -1
}

func AddToDB(db DB, key string, value int) error {
	if err := db.Add(key, value); err != nil {
		return err
	}
	return nil
}
