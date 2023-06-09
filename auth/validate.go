package auth

import (
	"ca_parks/core/module/user/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (i *Implementation) Validate(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("origin")
	method := r.Method
	fmt.Printf("method: %v\n", method)
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Accept", "application/json, multipart/mixed")
	w.Header().Set("Content-Type", "application/json")
	if method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	claims, err := i.ClaimsFromHeader(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, time.Millisecond*500)
	userRes, err := i.core.User().FetchUserByEmail(ctx, types.FetchUserByEmailRequest{
		Email: claims.Email,
		Limit: 1,
	})
	defer cancelFn()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("error getting user by email: %s, %v\n", claims.Email, err)
		return
	}

	if len(userRes.Results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Printf("user not found: %s\n", claims.Email)
		return
	}

	user := userRes.Results[0]
	user.Password = ""
	resJson, _ := json.Marshal(user)
	w.Write(resJson)
}

func (i *Implementation) ClaimsFromHeader(w http.ResponseWriter, r *http.Request) (*Claims, error) {
	tknStr, err := i.TokenFromHeader(r)
	if err != nil {
		fmt.Printf("error getting token from header\n")
		return nil, err
	}

	// Initialize a new instance of `Claims`
	claims, err := ParseToken(tknStr, &w)
	if err != nil {
		fmt.Printf("error parsing token: %v\n", err)
		return nil, err
	}

	return claims, nil
}

func (i *Implementation) TokenFromHeader(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "bearer ")
	if len(splitToken) < 2 {
		return "", errors.New("invalid token len")
	}

	return splitToken[1], nil
}
