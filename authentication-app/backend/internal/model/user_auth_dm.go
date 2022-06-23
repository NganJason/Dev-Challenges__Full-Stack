package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/config"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

type userAuthDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserAuthDM(ctx context.Context) UserAuthInterface {
	return &userAuthDM{
		ctx: ctx,
		db:  config.GetDBs().AuthDB,
	}
}

func (dm *userAuthDM) GetUserAuth(
	id *uint64,
	username *string,
	externalID *string,
	authMethod *int,
) (*UserAuth, error) {
	var userAuth UserAuth

	if username == nil && id == nil && externalID == nil {
		return nil, cerr.New(
			"must provide either userID, username or externalID",
			http.StatusBadRequest,
		)
	}

	if authMethod == nil {
		return nil, cerr.New(
			"auth_method cannot be empty",
			http.StatusBadRequest,
		)
	}

	baseQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE `,
		dm.getTableName(),
	)

	whereCols := make([]string, 0)
	args := make([]interface{}, 0)

	if username != nil {
		whereCols = append(whereCols, "username = ?")
		args = append(args, *username)
	}

	if externalID != nil {
		whereCols = append(whereCols, "external_id = ?")
		args = append(args, *username)
	}

	if authMethod != nil {
		whereCols = append(whereCols, "auth_method = ?")
		args = append(args, *authMethod)
	}

	where := strings.Join(whereCols, " AND ")
	query := baseQuery + where

	err := dm.db.QueryRow(
		query,
		args...,
	).Scan(
		&userAuth,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("query userAuth err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return &userAuth, nil
}

func (dm *userAuthDM) CreateUserAuth(req *CreateUserAuthRequest) (uint64, error) {
	if req.Username == "" {
		return 0, cerr.New(
			"username cannot be empty for user auth creation",
			http.StatusBadRequest,
		)
	}

	if req.AuthMethod == 0 {
		return 0, cerr.New(
			"invalid auth method",
			http.StatusBadRequest,
		)
	}

	existingUser, err := dm.GetUserAuth(nil, &req.Username, req.ExternalID, &req.AuthMethod)
	if err != nil {
		return 0, cerr.New(
			fmt.Sprintf("get existing user err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingUser != nil {
		return 0, cerr.New(
			"user already exist",
			http.StatusBadRequest,
		)
	}

	query := fmt.Sprintf(
		`
		INSERT INTO %s (username, auth_method, external_id, hashed_password, salt)
		VALUES(?, ?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		query,
		req.Username,
		req.AuthMethod,
		req.ExternalID,
		req.HashedPassword,
		req.Salt,
	)
	if err != nil {
		return 0, cerr.New(
			fmt.Sprintf("insert user auth into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	return uint64(lastInsertID), nil
}

func (dm *userAuthDM) getTableName() string {
	return "user_auth_tab"
}
