package controller

import (
	"chotot_product_ltruong/token"
	"context"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a server interceptor for authentication and authorization
type AuthInterceptor struct {
	jwtManager token.Maker
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager token.Maker) *AuthInterceptor {
	return &AuthInterceptor{jwtManager}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		ctx, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
//func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
//	return func(
//		srv interface{},
//		stream grpc.ServerStream,
//		info *grpc.StreamServerInfo,
//		handler grpc.StreamHandler,
//	) error {
//		log.Println("--> stream interceptor: ", info.FullMethod)
//
//		err := interceptor.authorize(stream.Context(), info.FullMethod)
//		if err != nil {
//			return err
//		}
//
//		return handler(srv, stream)
//	}
//}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.VerifyToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	//add new value to incoming ctx
	md.Append(UserIDCtx, strconv.Itoa(claims.UserID))
	md.Append(IsAdminCtx, strconv.FormatBool(claims.IsAdmin))
	newCtx := metadata.NewIncomingContext(ctx, md)
	return newCtx, nil
}
