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
	userID *uint64,
	loginID *string,
	authMethod *int,
) (*UserAuth, error) {
	var userAuth UserAuth

	if userID == nil && loginID == nil {
		return nil, cerr.New(
			"must provide either userID or loginID",
			http.StatusBadRequest,
		)
	}

	if userID == nil && authMethod == nil {
		return nil, cerr.New(
			"auth_method and userID cannot both be empty",
			http.StatusBadRequest,
		)
	}

	baseQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE `,
		dm.getTableName(),
	)

	whereCols := make([]string, 0)
	args := make([]interface{}, 0)

	if userID != nil {
		whereCols = append(whereCols, "id = ?")
		args = append(args, *userID)
	}

	if loginID != nil {
		whereCols = append(whereCols, "login_id = ?")
		args = append(args, *loginID)
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
		&userAuth.ID,
		&userAuth.LoginID,
		&userAuth.AuthMethod,
		&userAuth.HashedPassword,
		&userAuth.Salt,
		&userAuth.CreatedAt,
		&userAuth.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, cerr.New(
			fmt.Sprintf("query userAuth err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return &userAuth, nil
}

func (dm *userAuthDM) CreateUserAuth(req *CreateUserAuthRequest) (uint64, error) {
	if req.AuthMethod == 0 {
		return 0, cerr.New(
			"invalid auth method",
			http.StatusBadRequest,
		)
	}

	existingUser, err := dm.GetUserAuth(nil, req.LoginID, &req.AuthMethod)
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
		INSERT INTO %s (login_id, auth_method, hashed_password, salt)
		VALUES(?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		query,
		req.LoginID,
		req.AuthMethod,
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
