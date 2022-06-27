package model

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/config"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

type userInfoDM struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserInfoDM(ctx context.Context) UserInfoInterface {
	return &userInfoDM{
		ctx: ctx,
		db:  config.GetDBs().UserInfoDB,
	}
}

func (dm *userInfoDM) GetUserInfo(
	ID *uint64,
	userID *uint64,
) (*UserInfo, error) {
	var userInfo UserInfo

	if userID == nil && ID == nil {
		return nil, cerr.New(
			"userID and ID cannot both be empty",
			http.StatusBadRequest,
		)
	}

	baseQuery := fmt.Sprintf(
		`SELECT * FROM %s WHERE `,
		dm.getTableName(),
	)

	whereCols := make([]string, 0)
	args := make([]interface{}, 0)

	if ID != nil {
		whereCols = append(whereCols, "id = ?")
		args = append(args, *ID)
	}

	if userID != nil {
		whereCols = append(whereCols, "user_id = ?")
		args = append(args, *userID)
	}

	where := strings.Join(whereCols, " AND ")
	query := baseQuery + where

	err := dm.db.QueryRow(
		query,
		args...,
	).Scan(
		&userInfo.ID,
		&userInfo.UserID,
		&userInfo.Username,
		&userInfo.Bio,
		&userInfo.Phone,
		&userInfo.Email,
		&userInfo.Image,
		&userInfo.CreatedAt,
		&userInfo.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, cerr.New(
			fmt.Sprintf("query userInfo err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return &userInfo, nil
}

func (dm *userInfoDM) CreateUserInfo(
	req *CreateUserInfoRequest,
) (*UserInfo, error) {
	if req.UserID == nil {
		return nil, cerr.New(
			"userID cannot be nil",
			http.StatusBadRequest,
		)
	}

	existingUser, err := dm.GetUserInfo(nil, req.UserID)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing userInfo err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingUser != nil {
		return nil, cerr.New(
			"userInfo already exist",
			http.StatusBadRequest,
		)
	}

	query := fmt.Sprintf(
		`
		INSERT INTO %s (user_id, username, bio, phone, email)
		VALUES(?, ?, ?, ?, ?)
		`, dm.getTableName(),
	)

	result, err := dm.db.Exec(
		query,
		req.UserID,
		req.Username,
		req.Bio,
		req.Phone,
		req.Email,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("insert user info into db err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	lastInsertID, _ := result.LastInsertId()

	return dm.GetUserInfo(util.Uint64Ptr(uint64(lastInsertID)), nil)
}

func (dm *userInfoDM) UpdateUserInfo(
	req *UpdateUserInfoRequest,
) (*UserInfo, error) {
	if req.UserID == nil {
		return nil, cerr.New(
			"userID cannot be nil",
			http.StatusBadRequest,
		)
	}

	tx, err := dm.db.BeginTx(dm.ctx, nil)
	if err != nil {
		return nil, cerr.New(
			"start transaction error",
			http.StatusBadGateway,
		)
	}
	defer tx.Rollback()

	sqlQuery := fmt.Sprintf(
		`
		SELECT * FROM %s WHERE user_id = ? limit 1
		FOR UPDATE
		`, dm.getTableName(),
	)

	var existingUser UserInfo
	err = tx.QueryRowContext(
		dm.ctx,
		sqlQuery,
		*req.UserID).Scan(
		&existingUser.ID,
		&existingUser.UserID,
		&existingUser.Username,
		&existingUser.Bio,
		&existingUser.Phone,
		&existingUser.Email,
		&existingUser.Image,
		&existingUser.CreatedAt,
		&existingUser.UpdatedAt,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing userInfo err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingUser.ID == nil {
		return nil, cerr.New(
			"userInfo does not exist",
			http.StatusBadRequest,
		)
	}

	if req.Username != nil {
		existingUser.Username = req.Username
	}

	if req.Bio != nil {
		existingUser.Bio = req.Bio
	}

	if req.Phone != nil {
		existingUser.Phone = req.Phone
	}

	if req.Email != nil {
		existingUser.Email = req.Email
	}

	sqlQuery = fmt.Sprintf(
		`
		UPDATE %s
		SET username = ?, bio = ?, phone = ?, email = ?
		WHERE user_id = ?
		`,
		dm.getTableName(),
	)

	_, err = tx.ExecContext(
		dm.ctx,
		sqlQuery,
		existingUser.Username,
		existingUser.Bio,
		existingUser.Phone,
		existingUser.Email,
		existingUser.UserID,
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("updatte user info err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	err = tx.Commit()
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("commit update transaction err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return &existingUser, nil
}

func (dm *userInfoDM) getTableName() string {
	return "user_info_tab"
}
