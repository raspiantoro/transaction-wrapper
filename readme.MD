# Transaction Wrapper Example

This example is shows you how to perform transactions between repository in golang application using gorm


## Migrations
This project have no database migrations, run `./sql/db.sql` manually on your mysql database

## Test Endpoint
Import postman collection from `./postman` directory

## Example
This example is taken from person service

### Without Transaction
```golang
func (ps *personService) CreatePerson(ctx context.Context, person payload.CreatePersonRequests) (err error) {

	user := &model.User{
		ID:       uuid.NewV4().String(),
		UserName: person.UserName,
		Password: person.Password,
	}

	profile := &model.Profile{
		ID:        uuid.NewV4().String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Age:       person.Age,
	}

	err = ps.Repository.User.CreateUser(ctx, user)
	if err != nil {
		return
	}

	// commenting the line below to throw error on CreateProfile
	// will resulting inconsistency on database
	profile.UserID = user.ID

	err = ps.Repository.Profile.CreateProfile(ctx, profile)

	return
}
```

### With Transaction
```golang
func (ps *personService) CreatePersonTx(ctx context.Context, person payload.CreatePersonRequests) (err error) {

	user := &model.User{
		ID:       uuid.NewV4().String(),
		UserName: person.UserName,
		Password: person.Password,
	}

	profile := &model.Profile{
		ID:        uuid.NewV4().String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		Age:       person.Age,
	}

	err = ps.Repository.WithTransactions(func(tx *gorm.DB) (err error) {

		err = ps.Repository.User.CreateUser(ctx, user, repository.WithDBInstance(tx))
		if err != nil {
			return
		}

		// commenting the line below to throw error on CreateProfile
		// won't resulting inconsistency on database, CreateUser will be rollback
		profile.UserID = user.ID

		err = ps.Repository.Profile.CreateProfile(ctx, profile, repository.WithDBInstance(tx))

		return
	})

	return
}
```