package guestsession

import (
	"context"
	"fmt"
	"time"

	pb "order-service/pb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GuestSessionServiceServer struct {
	pb.UnimplementedGuestSessionServiceServer
	DB *pgx.Conn
}

// CreateGuestSession creates a new guest session.
func (s *GuestSessionServiceServer) CreateGuestSession(ctx context.Context, req *pb.CreateGuestSessionRequest) (*pb.CreateGuestSessionResponse, error) {
	// Generate a new guest session ID if not provided.
	var guestSessionID string
	if req.GuestSessionId == "" {
		guestSessionID = uuid.New().String()
	} else {
		guestSessionID = req.GuestSessionId
	}

	// Generate a guest ID if not provided.
	guestID := req.GuestId
	if guestID == "" {
		guestID = uuid.New().String()
	}

	// Insert a new guest session into the database.
	query := `
		INSERT INTO guest_sessions (id, guest_id, session_token, started_at, status)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := s.DB.Exec(ctx, query, uuid.New().String(), guestID, guestSessionID, time.Now(), "active")
	if err != nil {
		return nil, fmt.Errorf("failed to create guest session: %w", err)
	}

	return &pb.CreateGuestSessionResponse{
		GuestSessionId: guestSessionID,
		Status:         "active",
	}, nil
}

// (Implement Get, Update, Delete similarly as needed.)
