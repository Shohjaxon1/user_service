package postgres

import (
	"database/sql"
	"time"
	pb "user_service/genproto/user_service"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo creates a new userRepocmd/main.go
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

// Create inserts a new user into the database
func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	query := `
		INSERT INTO users (id, first_name, last_name, phone_number, password, gender, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id, first_name, last_name, phone_number, password, gender, created_at
	`
	err := r.db.QueryRow(query, user.Id, user.FirstName, user.LastName, user.PhoneNumber, user.Password, user.Gender, time.Now()).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Password,
		&user.Gender,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetById retrieves a user by their ID
func (r *userRepo) GetById(req *pb.GetByIdRequest) (*pb.User, error) {
	query := "SELECT id, first_name, last_name, phone_number, password, gender, created_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, req.Id)

	user := &pb.User{}
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Password,
		&user.Gender,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetByFirstName(req *pb.GetByFirstNameRequest) (*pb.User, error) {
	query := "SELECT id, first_name, last_name, phone_number, password, gender, created_at FROM users WHERE first_name = $1"
	row := r.db.QueryRow(query, req.FirstName)

	user := &pb.User{}
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Password,
		&user.Gender,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetByLastName(req *pb.GetByLastNameRequest) (*pb.User, error) {
	query := "SELECT id, first_name, last_name, phone_number, password, gender, created_at FROM users WHERE last_name = $1"
	row := r.db.QueryRow(query, req.LastName)

	user := &pb.User{}
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Password,
		&user.Gender,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetByPhoneNumber(req *pb.GetByPhoneNumberRequest) (*pb.User, error) {
	query := "SELECT id, first_name, last_name, phone_number, password, gender, created_at FROM users WHERE phone_number = $1"
	row := r.db.QueryRow(query, req.PhoneNumber)

	user := &pb.User{}
	err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.Password,
		&user.Gender,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates a user's details in the database
func (r *userRepo) Update(req *pb.User) (*pb.User, error) {
	query := `
        UPDATE users
        SET first_name = $2, last_name = $3, phone_number = $4, password = $5, gender = $6, updated_at = $7
        WHERE id = $1
        RETURNING id, first_name, last_name, phone_number, password, gender, updated_at
    `
	row := r.db.QueryRow(query, req.Id, req.FirstName, req.LastName, req.PhoneNumber, req.Password, req.Gender, time.Now())
	updatedUser := &pb.User{}
	err := row.Scan(
		&updatedUser.Id,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.PhoneNumber,
		&updatedUser.Password,
		&updatedUser.Gender,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// Delete removes a user from the database
func (r *userRepo) Delete(req *pb.GetByIdRequest) (*pb.User, error) {
	query := "DELETE FROM users WHERE id = $1 RETURNING id, first_name, last_name, phone_number, password, gender, created_at, updated_at"
	row := r.db.QueryRow(query, req.Id)
	var updatedAt sql.NullTime
	deletedUser := &pb.User{}
	err := row.Scan(
		&deletedUser.Id,
		&deletedUser.FirstName,
		&deletedUser.LastName,
		&deletedUser.PhoneNumber,
		&deletedUser.Password,
		&deletedUser.Gender,
		&deletedUser.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	if updatedAt.Valid {
		deletedUser.UpdatedAt = updatedAt.Time.String()
	} else {
		deletedUser.UpdatedAt = ""
	}
	return deletedUser, nil
}

func (r *userRepo) GetAll(req *pb.GetAllRequest) (*pb.AllUsers, error) {
	intLimit := cast.ToInt(req.Limit)
	intPage := cast.ToInt(req.Page)
	offset := (intPage - 1) * intLimit

	query := `
        SELECT id, first_name, last_name, phone_number, password, gender, created_at, updated_at
        FROM users
        LIMIT $1 OFFSET $2
    `
	rows, err := r.db.Query(query, intLimit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*pb.User{}

	for rows.Next() {
		var updatedAt sql.NullTime
		user := &pb.User{}
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.Password,
			&user.Gender,
			&user.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time.String()
		} else {
			user.UpdatedAt = ""
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.AllUsers{Users: users}, nil
}
