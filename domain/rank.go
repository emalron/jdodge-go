package domain

type Rank struct {
    Name string
    Score int
    ReplayData string
    Time string
}

type RankRepository interface {
    FindAll() ([]Rank, error)
    FindByID(id string) ([]Rank, error)
}
