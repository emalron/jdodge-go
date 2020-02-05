package interfaces

import(
    "jdodge-go/domain"
    "jdodge-go/usecases"
    "fmt"
)

type DBHandler interface {
    Execute(statement string)
    Query(statement string) Row
}

type Row interface {
    Scan(dest ...interface{}) error
    Next() bool
}

type DBRepo struct {
    dbHandlers map[string]DBHandler
    dbHandler DBHandler
}

type DBRankRepo DBRepo
type DBUserRepo DBRepo

func NewDBRankRepo(dbHandlers map[string]DBHandler) *DBRankRepo {
    dbRankRepo := new(DBRankRepo)
    dbRankRepo.dbHandlers = dbHandlers
    dbRankRepo.dbHandler = dbHandlers["DBRankRepo"]
    return dbRankRepo
}

func NewDBUserRepo(dbHandlers map[string]DBHandler) *DBUserRepo {
    dbUserRepo := new(DBUserRepo)
    dbUserRepo.dbHandlers = dbHandlers
    dbUserRepo.dbHandler = dbHandlers["DBUserRepo"]
    return dbUserRepo
}

func (repo *DBRankRepo) FindAll() ([]domain.Rank, error) {
    rows := repo.dbHandler.Query("SELECT name, score, replay_data, time FROM view_ranking")
    results := make([]domain.Rank, 0)
    rank := domain.Rank{}
    for rows.Next() {
        scanErr := rows.Scan(&rank.Name, &rank.Score, &rank.ReplayData, &rank.Time)
        if scanErr != nil {
            return results, scanErr
        }
        results = append(results, rank)
    }
    return results, nil
}

func (repo *DBRankRepo) FindByID(id string) ([]domain.Rank, error) {
    rows := repo.dbHandler.Query(fmt.Sprintf(`SELECT name, score, replay_data, time FROM view_ranking WHERE id = '%s'`, id))
    results := make([]domain.Rank, 0)
    rank := domain.Rank{}
    for rows.Next() {
        scanErr := rows.Scan(rank)
        if scanErr != nil {
            return results, scanErr
        }
        results = append(results, rank)
    }
    return results, nil
}

func (repo *DBUserRepo) FindByID(id string) (usecases.User, error) {
    rows := repo.dbHandler.Query(fmt.Sprintf(`SELECT id FROM users WHERE id = '%s'`, id));
    user := usecases.User{}
    for rows.Next() {
        scanErr := rows.Scan(user)
        if scanErr != nil {
            return user, scanErr
        }
    }
    return user, nil
}
