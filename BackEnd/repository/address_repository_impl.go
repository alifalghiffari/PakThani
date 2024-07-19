package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepositoryImpl() AddressRepository {
	return &AddressRepositoryImpl{}
}

func (repository *AddressRepositoryImpl) AddAddress(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "INSERT INTO addresses(user_id, nama_penerima, kabupaten, kecamatan, kelurahan, alamat, no_hp, note) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, address.UserId, address.NamaPenerima, address.Kabupaten, address.Kecamatan, address.Kelurahan, address.Alamat, address.NoHP, address.Note)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	address.Id = int(id)
	return address
}

func (repository *AddressRepositoryImpl) UpdateAddress(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "UPDATE addresses SET nama_penerima = ?, kabupaten = ?, kecamatan = ?, kelurahan = ?, alamat = ?, no_hp = ?, note = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, address.NamaPenerima, address.Kabupaten, address.Kecamatan, address.Kelurahan, address.Alamat, address.NoHP, address.Note, address.Id)
	helper.PanicIfError(err)

	return address
}

func (repository *AddressRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Address, error) {
	SQL := `
		SELECT a.id, a.user_id, a.nama_penerima, a.kabupaten, a.kecamatan, a.kelurahan, a.alamat, a.no_hp, a.note, u.id, u.username, u.email, u.role
		FROM addresses a
		LEFT JOIN user u ON a.user_id = u.id
		WHERE u.id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []domain.Address
	for rows.Next() {
		address := domain.Address{}
		user := domain.User{}
		if err := rows.Scan(&address.Id, &address.UserId, &address.NamaPenerima, &address.Kabupaten, &address.Kecamatan, &address.Kelurahan, &address.Alamat, &address.NoHP, &address.Note, &user.Id, &user.Username, &user.Email, &user.Role); err != nil {
			return nil, err
		}
		address.User = append(address.User, user)
		addresses = append(addresses, address)
	}

	return addresses, nil
}
