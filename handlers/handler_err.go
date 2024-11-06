package handlers


import(
	"net/http"
	"github.com/Flamme97/todoListGo/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request){
	utils.RespondWithError(w, 400, "an Error occurred")
}