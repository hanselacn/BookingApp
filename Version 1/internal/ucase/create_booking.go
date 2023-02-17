package ucase

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gitlab.privy.id/go_graphql/internal/appctx"
	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/internal/repositories"
)

type createBooking struct {
	bookingRepository repositories.Booking
}

func (cb createBooking) Serve(data *appctx.Data) appctx.Response {
	a := &entity.Booking{}

	err := data.Cast(&a)

	if err != nil {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Booking").WithState("AddBookingFailed").WithMessage("Adding Book Failed").WithError(err.Error())
	}

	user := entity.User{}

	user, err = cb.bookingRepository.VerifyBookingBy(data.Request.Context(), a.BookedBy)

	if user.Username != a.BookedBy {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Booking").WithState("AddBookingFailed").WithMessage("Adding Book Failed").WithError(err.Error())
	}

	day := entity.Day{}

	day, err = cb.bookingRepository.VerifyBookingDay(data.Request.Context(), a.BookedDay)

	if day.DayName != a.BookedDay {
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Booking").WithState("AddBookingFailed").WithMessage("Adding Book Failed").WithError(err.Error())
	}

	booking := entity.Booking{}

	booking, _ = cb.bookingRepository.VerifyBookingEntity(data.Request.Context(), a.BookedDay, a.Session, a.BookedRoom)

	if booking.Session == a.Session {
		err = fmt.Errorf("session already booked")
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Booking").WithState("AddBookingFailed").WithMessage("Adding Book Failed").WithError(err.Error())
	}

	id, err := cb.bookingRepository.Create(data.Request.Context(), a)

	if err != nil {
		fmt.Println(err)
		return *appctx.NewResponse().WithCode(http.StatusBadRequest).WithStatus("FAIL").WithEntity("Booking").WithState("AddBookingFailed").WithMessage("Adding Book Failed").WithError(err.Error())
	}

	type resp struct {
		ID uuid.UUID `json:"id"`
	}

	var res resp
	res.ID = *id

	return *appctx.NewResponse().WithCode(http.StatusCreated).WithStatus("SUCCESS").WithEntity("Booking").WithState("AddBookingSuccess").WithMessage("Adding Booking Success").WithData(res)
}

func NewCreateBooking(bookingRepository repositories.Booking) *createBooking {
	return &createBooking{
		bookingRepository: bookingRepository,
	}
}
