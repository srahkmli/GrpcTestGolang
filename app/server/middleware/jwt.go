package middleware

import (
	"context"
	"github.com/google/uuid"
)

// global middleware example for all routes
func (m *middle) JWT(ctx context.Context) (context.Context, error) {
	//	meta, ok := metadata.FromIncomingContext(ctx)
	//	if !ok {
	//		// should fix with new error type
	//		return nil, status.New(codes.Unauthenticated, "").Err()
	//	}
	//
	//	authorization := meta.Get("authorization")
	//	if len(authorization) == 0 {
	//		return nil, status.New(codes.Unauthenticated, "no auth header provided").Err()
	//	}
	//
	//	strArr := strings.Split(authorization[0], " ")
	//
	//	if len(strArr) != 2 {
	//		return nil, status.New(codes.Unauthenticated, "malformed auth header").Err()
	//	}
	//	resp, err := oauth.JWT.VerifyToken(ctx, strArr[1])
	//	if err != nil {
	//		return nil, status.New(codes.Unauthenticated, "invalid token").Err()
	//	}
	//
	//	ctx = context.WithValue(ctx, "user_id", resp.UserID) // the user uuid
	//	ctx = context.WithValue(ctx, "dialer", resp.Dialer)  // the user mobile

	ctx = context.WithValue(ctx, "user_id", uuid.New().String()) // the user uuid
	ctx = context.WithValue(ctx, "dialer", "9100000")            // the user mobile
	return ctx, nil
}