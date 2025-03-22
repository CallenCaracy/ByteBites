package guestsession

import (
	"context"
	"fmt"
	"time"

	pb "order-service/pb"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GuestSessionServiceServer struct {
	pb.UnimplementedGuestSessionServiceServer
	DB *pgx.Conn
}

func (s *GuestSessionServiceServer) CreateGuestSession(ctx context.Context, req *pb.CreateGuestSessionRequest) (*pb.CreateGuestSessionResponse, error) {
	sessionToken := req.SessionToken
	if sessionToken == "" {
		sessionToken = uuid.New().String()
	}

	query := `
		INSERT INTO guest_sessions (session_token, started_at, status)
		VALUES ($1, $2, $3)
	`
	_, err := s.DB.Exec(ctx, query, sessionToken, time.Now(), "active")
	if err != nil {
		return nil, fmt.Errorf("failed to create guest session: %w", err)
	}

	return &pb.CreateGuestSessionResponse{
		SessionToken: sessionToken,
		Status:       "active",
	}, nil
}

func (s *GuestSessionServiceServer) GetGuestSession(ctx context.Context, req *pb.GetGuestSessionRequest) (*pb.GetGuestSessionResponse, error) {
	var sessionToken, status string
	var startedAt time.Time
	var endedAt *time.Time

	query := `
		SELECT id, session_token, started_at, ended_at, status
		FROM guest_sessions
		WHERE id = $1
	`
	row := s.DB.QueryRow(ctx, query, req.GuestId)
	err := row.Scan(&req.GuestId, &sessionToken, &startedAt, &endedAt, &status)
	if err != nil {
		return nil, fmt.Errorf("failed to get guest session: %w", err)
	}

	var sessionEndTime *timestamppb.Timestamp
	if endedAt != nil {
		sessionEndTime = timestamppb.New(*endedAt)
	} else {
		sessionEndTime = nil
	}

	return &pb.GetGuestSessionResponse{
		GuestId:          req.GuestId,
		SessionToken:     sessionToken,
		SessionStartTime: timestamppb.New(startedAt),
		SessionEndTime:   sessionEndTime,
		Status:           status,
	}, nil
}

func (s *GuestSessionServiceServer) UpdateGuestSession(ctx context.Context, req *pb.UpdateGuestSessionRequest) (*pb.UpdateGuestSessionResponse, error) {
	// If there is no inputted update for status, then that means it supposed to be updated by default 'complete' assuming the guest has finished their session/order.
	status := req.Status
	if status == "" {
		status = "complete"
	}

	query := `
        UPDATE guest_sessions
        SET status = $3, ended_at = NOW()
        WHERE id = $1 AND session_token = $2
    `
	_, err := s.DB.Exec(ctx, query, req.GuestId, req.SessionToken, status)
	if err != nil {
		return nil, fmt.Errorf("failed to update guest session: %w", err)
	}

	return &pb.UpdateGuestSessionResponse{
		SessionToken:   req.SessionToken,
		SessionEndTime: timestamppb.Now(),
		Status:         status,
	}, nil
}

func (s *GuestSessionServiceServer) DeleteGuestSession(ctx context.Context, req *pb.DeleteGuestSessionRequest) (*pb.DeleteGuestSessionResponse, error) {
	query := `
		DELETE FROM guest_sessions
		WHERE id = $1
	`
	_, err := s.DB.Exec(ctx, query, req.GuestId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete guest session: %w", err)
	}

	return &pb.DeleteGuestSessionResponse{Status: "Deleted"}, nil
}

// Checks the status of the guest session if its still active or not. For order service logic use.
func (s *GuestSessionServiceServer) CheckGuestStatus(ctx context.Context, req *pb.CheckGuestStatusRequest) (*pb.CheckGuestStatusResponse, error) {
	var status string

	query := `
		SELECT status
		FROM guest_sessions
		WHERE id = $1
	`
	row := s.DB.QueryRow(ctx, query, req.GuestId)
	err := row.Scan(&status)
	if err != nil {
		return nil, fmt.Errorf("failed to check guest status: %w", err)
	}

	return &pb.CheckGuestStatusResponse{Status: status}, nil
}
