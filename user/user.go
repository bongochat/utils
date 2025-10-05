package user

import (
	"strconv"
	"strings"

	"github.com/bongochat/utils/resterrors"
	"github.com/gin-gonic/gin"
)

func GetID(userIdParam string) (int64, resterrors.RestError) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, resterrors.NewBadRequestError("user id should be a number", "")
	}
	return userId, nil
}

func GetAccountNumber(c *gin.Context) (int64, resterrors.RestError) {
	accountNumberAny, exists := c.Get("accountnumber")
	if !exists {
		return 0, resterrors.NewUnauthorizedError("account number not found in context", "")
	}

	// Type assertion: userId should be int64 as set in middleware
	accountNumber, ok := accountNumberAny.(int64)
	if !ok {
		return 0, resterrors.NewUnauthorizedError("account number has invalid type", "")
	}

	return accountNumber, nil
}

func GetUserID(c *gin.Context) (int64, resterrors.RestError) {
	userIdAny, exists := c.Get("userId")
	if !exists {
		return 0, resterrors.NewUnauthorizedError("user id not found in context", "")
	}

	// Type assertion: userId should be int64 as set in middleware
	userId, ok := userIdAny.(int64)
	if !ok {
		return 0, resterrors.NewUnauthorizedError("user id has invalid type", "")
	}

	return userId, nil
}

func GetAccessToken(c *gin.Context) (string, resterrors.RestError) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", resterrors.NewUnauthorizedError("authorization header missing", "")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", resterrors.NewUnauthorizedError("invalid token format", "")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return "", resterrors.NewUnauthorizedError("token is empty", "")
	}

	return token, nil
}

func GetClientID(clientIdParam string) (string, resterrors.RestError) {
	if clientIdParam == "" {
		return "", resterrors.NewBadRequestError("client id should be a string", "")
	}
	return clientIdParam, nil
}
