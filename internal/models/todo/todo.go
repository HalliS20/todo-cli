package todo

type Todo struct {
	ID       uint   `gorm:"primaryKey"`
	Index    int    `gorm:"not null"`
	Title    string `gorm:"not null"`
	Done     bool   `gorm:"not null;default:false"`
	Dir      bool   `gorm:"not null;default:false"`
	ParentID uint   `gorm:"not null"`
}

type Todos []*Todo

func (a Todos) Len() int           { return len(a) }
func (a Todos) Swap(i, j int)      { (a)[i], (a)[j] = (a)[j], (a)[i] }
func (a Todos) Less(i, j int) bool { return (a)[i].Index < (a)[j].Index }
