package server

import (
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"net/http"
	"strconv"
)

func (a App) GetDay(writer http.ResponseWriter, request *http.Request) {
	var err error
	var todaysPrompts *domain.Prompt
	d := request.URL.Query().Get("day")
	if d == "" {
		todaysPrompts, err = a.prompts.GetToday()
	} else {
		daysBack, err := strconv.Atoi(d)
		if err != nil {
			a.writeError(fmt.Errorf("%s could not be converted to a number", d), writer)
			return
		}
		todaysPrompts, err = a.prompts.GetDay(daysBack)
	}

	if err != nil {
		a.writeError(err, writer)
		return
	}

	a.writeJSON(todaysPrompts, writer)
}
