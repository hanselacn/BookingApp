package ucase

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type getBookByName struct {
	repo repositories.Booking
}

func NewGetBookByName(repo repositories.Booking) *getBookByName {
	return &getBookByName{repo: repo}
}

func (u *getBookByName) Serve(data *appctx.Data) appctx.Response {

	bookedby := mux.Vars(data.Request)["bookedBy"]
	books, err := u.repo.GetBookByName(data.Request.Context(), bookedby)

	if err != nil {
		err := fmt.Errorf("getting books: %w", err)
		fmt.Println(err)

		return *appctx.NewResponse().WithCode(http.StatusInternalServerError).WithStatus("ERROR").WithMessage("Internal Server Error")
	}
	return *appctx.NewResponse().WithCode(http.StatusOK).WithStatus("SUCCESS").WithEntity("Bookings").WithState("GettingBookingSuccess").WithMessage("Getting Bookings Success").WithData(books)
}
