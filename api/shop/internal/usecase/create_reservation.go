package usecase

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/takeuchima0/async_serverless_application_sample/api/shop/internal/gen"
)

func (u reservationUseCase) CreateReservation(ctx *gin.Context, request gen.CreateReservationRequestObject) (gen.CreateReservationResponseObject, error) {

	time.Sleep(300 * time.Millisecond)

	reservationID := uuid.New().String()

	return gen.CreateReservation201JSONResponse{
		ReservationId: reservationID,
	}, nil
}
