# Defitinion

This is a Software used to allow online book's reservation.

The application uses, Hexagonal architecture and DDD (Domain Driven Design)

# Entities 


| **users**  | **books**      | **users_book_reservations** |
|------------|----------------|-----------------------------|
| id         | id             | id                          |
| name       | title          | user_id                     |
| email      | author         | book_id                     |
| cpf        | published_date | created_at                  |
| birth_date | genre          | updated_at                  |
| password   | created_at     | updated_at                  |
| created_at | updated_at     |                             |
| updated_at |      		  |                             |

# Business Logic

- User creation.

- Authetication/Authorization - Users need to authenticate to our platform to use the app features.
The main roles will be: admin and common.

- Creation and maintaneance of books (just Admins).

- Reservation of books, one book can be reserved only by one User at a time. Other users
can enter in a queue to get the current reserved book.

- User reserve the book for a determined time, if its not returned in time a fine will be applied. 

# Improvments

- Review on books where Users can talk about the book and give ratings.

teste, err := time.Parse("2006-01-02", "2024-08-27")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(teste)