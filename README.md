# Defitinion

This is a Software used to allow online book's reservation.

The application uses, Hexagonal architecture and DDD (Domain Driven Design)

# Entities 


| **Users**       | **Books**     |
| -------------- | -------------- |
| ID             | ID             |
| Name           | Title          |
| Email          | Author         |
| CPF            | Published Date |
| Birth Date     | Genre          |
|                |                |


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