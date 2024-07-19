package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-workshop/go-api-ecom/model/domain"
	"time"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(username, password, email, role, no_telepon, is_verified, verification_token) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Email, user.Role, user.NoTelepon, user.IsVerified, user.VerificationToken)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, email domain.User) domain.User {
	SQL := "select id, username, password, email, no_telepon, role, is_verified, verification_token from user where email = ? and is_verified = true"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&email.Id, &email.Username, &email.Password, &email.Email, &email.NoTelepon, &email.Role, &email.IsVerified, &email.VerificationToken)
		if err != nil {
			panic(err)
		}
		return email
	} else {
		panic("email or password is incorrect")
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
    SQL := "SELECT id, username, password, email, no_telepon, role, is_verified, verification_token FROM user WHERE id = ?"
    row := tx.QueryRowContext(ctx, SQL, userId)

    user := domain.User{}
    err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role, &user.IsVerified, &user.VerificationToken)
    if err != nil {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("user not found")
        }
        return user, err
    }

    return user, nil
}


func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, username, password, email, no_telepon, role, is_verified, verification_token from user"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role, &user.IsVerified, &user.VerificationToken)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindByRole(ctx context.Context, tx *sql.Tx, role bool) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role, is_verified, verification_token from user where role = ?"
	rows, err := tx.QueryContext(ctx, SQL, role)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role, &user.IsVerified, &user.VerificationToken)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role, is_verified, verification_token from user where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role, &user.IsVerified, &user.VerificationToken)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindByEmailGoogle(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	query := "SELECT id, username, email, role, no_telepon FROM user WHERE email = ?"
	var user domain.User
	err := tx.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.NoTelepon)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}

func (repository *UserRepositoryImpl) VerifyEmail(ctx context.Context, tx *sql.Tx, verificationtoken string) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role, is_verified, verification_token from user where verification_token = ?"
	rows, err := tx.QueryContext(ctx, SQL, verificationtoken)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role, &user.IsVerified, &user.VerificationToken)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) UpdateVerificationStatus(ctx context.Context, tx *sql.Tx, email string, verificationtoken string) (domain.User, error) {
	SQL := "update user set is_verified = true where email = ? and verification_token = ?"
	_, err := tx.ExecContext(ctx, SQL, email, verificationtoken)
	if err != nil {
		panic(err)
	}

	user, err := repository.FindByEmail(ctx, tx, email)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) InsertResetPassword(ctx context.Context, tx *sql.Tx, token domain.ResetPassword) (domain.ResetPassword, error) {
	SQL := "insert into reset_password(user_id, token, expired_at) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, token.UserId, token.Token, token.Expired_at)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	token.Id = int(id)
	return token, nil
}

func (repository *UserRepositoryImpl) FindByToken(ctx context.Context, tx *sql.Tx, token string) (domain.ResetPassword, error) {
    SQL := "SELECT id, user_id, token, expired_at FROM reset_password WHERE token = ?"
    rows, err := tx.QueryContext(ctx, SQL, token)
    if err != nil {
        return domain.ResetPassword{}, fmt.Errorf("error querying token: %v", err)
    }
    defer rows.Close()

    resetPassword := domain.ResetPassword{}
    if rows.Next() {
        var expiredAt []uint8 // Use []uint8 to store datetime from database
        err := rows.Scan(&resetPassword.Id, &resetPassword.UserId, &resetPassword.Token, &expiredAt)
        if err != nil {
            return domain.ResetPassword{}, fmt.Errorf("error scanning token row: %v", err)
        }
        
        // Convert []uint8 (datetime from database) to time.Time
        resetPassword.Expired_at, err = time.Parse("2006-01-02 15:04:05", string(expiredAt))
		if err != nil {
			return domain.ResetPassword{}, fmt.Errorf("error parsing expired_at: %v", err)
		}

        return resetPassword, nil
    }

    return domain.ResetPassword{}, fmt.Errorf("token not found: %s", token)
}


func (repository *UserRepositoryImpl) DeletedByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "delete from reset_password where user_id =?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	if err != nil {
		panic(err)
	}

	user, err := repository.FindById(ctx, tx, userId)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
    SQL := "UPDATE user SET password = ? WHERE id = ?"
    _, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
    if err != nil {
        fmt.Println("UpdatePassword SQL error:", err)
        return user, err
    }

    return user, nil
}

