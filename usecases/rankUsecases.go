package usecases

import(
    "fmt"
    "jdodge-go/domain"
)

type User struct {
    ID string
}

type UserRepository interface {
    FindByID(id string) (User, error)
}

type RankInteractor struct {
    RankRepository domain.RankRepository
    UserRepository UserRepository
}

func (Interactor *RankInteractor) ShowAllRanks() ([]domain.Rank, error) {
    ranks, err := Interactor.RankRepository.FindAll()
    if err != nil {
        return make([]domain.Rank,0), err
    }
    return ranks, nil
}

func (Interactor *RankInteractor) ShowByID(id string) ([]domain.Rank, error) {
    user, userErr := Interactor.UserRepository.FindByID(id)
    if userErr != nil {
        userErr := fmt.Errorf("no such user")
        return nil, userErr
    }
    ranks, ranksErr := Interactor.RankRepository.FindByID(user.ID)
    if ranksErr != nil {
        ranksErr := fmt.Errorf("%s has no rank records", user.ID)
        return nil, ranksErr
    }
    return ranks, nil
}

