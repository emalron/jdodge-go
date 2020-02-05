package interfaces

import(
    "net/http"
    "jdodge-go/domain"
    "jdodge-go/usecases"
    "fmt"
    "io"
    "os"
)

type RankInteractor interface {
    ShowAllRanks()([]domain.Rank,error)
}

type WebserviceHandler struct {
    RankInteractor *usecases.RankInteractor
}

func (handle WebserviceHandler) ShowAll(w http.ResponseWriter, r *http.Request) {
    fmt.Println("you hit the ShowAll()")
    ranks, err := handle.RankInteractor.ShowAllRanks()
    fmt.Println("how many rank records? ", len(ranks))
    if err != nil {
        fmt.Fprintf(os.Stderr, "ShowAll error ", err)
    }
    for _, rank := range ranks {
        w.Header().Set("Content-Type", "text/html;charset=utf-8")
        io.WriteString(w, fmt.Sprintf("%s %d<br>", rank.Name, rank.Score))
    }
}

func (handle WebserviceHandler) ShowByID(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    ranks, _ := handle.RankInteractor.ShowByID(id)
    for _, rank := range ranks {
        io.WriteString(w, fmt.Sprintf(rank.Name, rank.Score))
    }
}
