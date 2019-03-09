package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	prisma "github.com/sample-graphql-prisma/generated/prisma-client"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := Resolver{
		Prisma: client,
	}

	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver})))

	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"context"
// 	"fmt"

// 	prisma "github.com/sample-graphql-prisma/generated/prisma-client"
// )

// func main() {
// 	client := prisma.New(nil)
// 	ctx := context.TODO()

// 	// Create a new user with two posts
// 	name := "Bob"
// 	email := "bob@prisma.io"
// 	title1 := "Join us for GraphQL Conf in 2019"
// 	title2 := "Subscribe to GraphQL Weekly for GraphQL news"
// 	newUser, err := client.CreateUser(prisma.UserCreateInput{
// 		Name:  name,
// 		Email: &email,
// 		Posts: &prisma.PostCreateManyWithoutAuthorInput{
// 			Create: []prisma.PostCreateWithoutAuthorInput{
// 				prisma.PostCreateWithoutAuthorInput{
// 					Title: title1,
// 				},
// 				prisma.PostCreateWithoutAuthorInput{
// 					Title: title2,
// 				},
// 			},
// 		},
// 	}).Exec(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Created new user: %+v\n", newUser)

// 	allUsers, err := client.Users(nil).Exec(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", allUsers)

// 	allPosts, err := client.Posts(nil).Exec(ctx)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%+v\n", allPosts)
// }

// // package main

// // import (
// // 	"context"
// // 	"fmt"
// // 	"log"

// // 	prisma "github.com/sample-graphql-prisma/generated/prisma-client"
// // )

// // func main() {
// // 	client := prisma.New(nil)
// // 	ctx := context.TODO()

// // 	name := "Alice"
// // 	newUser, err := client.CreateUser(prisma.UserCreateInput{
// // 		Name: name,
// // 	}).Exec(ctx)

// // 	if err != nil {
// // 		log.Panic(err)
// // 	}
// // 	fmt.Printf("Create new user %+v\n", newUser)

// // 	id := "cjt0yz4s700340853yb0erxi6"
// // 	// userByID, err := client.User(prisma.UserWhereUniqueInput{
// // 	// 	ID: &id,
// // 	// }).Exec(ctx)

// // 	// if err != nil {
// // 	// 	recover()
// // 	// }

// // 	// fmt.Printf("User by ID: %+v", userByID)

// // 	filter := "Alice"
// // 	posts, err := client.Users(&prisma.UsersParams{
// // 		Where: &prisma.UserWhereInput{
// // 			Name: &filter,
// // 		},
// // 	}).Exec(ctx)

// // 	if err != nil {
// // 		log.Panic(err)
// // 	}

// // 	fmt.Printf("User by filter: %+v", posts)

// // 	newName := "Bob"
// // 	updatedUser, err := client.UpdateUser(prisma.UserUpdateParams{
// // 		Where: prisma.UserWhereUniqueInput{
// // 			ID: &id,
// // 		},
// // 		Data: prisma.UserUpdateInput{
// // 			Name: &newName,
// // 		},
// // 	}).Exec(ctx)

// // 	if err != nil {
// // 		log.Panic(err)
// // 	}

// // 	fmt.Println("User updated: %+v", updatedUser)

// // 	deletedUser, err := client.DeleteUser(prisma.UserWhereUniqueInput{
// // 		ID: &id,
// // 	}).Exec(ctx)

// // 	if err != nil {
// // 		log.Panic(err)
// // 	}
// // 	fmt.Printf("Deleted a user %+v ", deletedUser)

// // 	users, err := client.Users(nil).Exec(ctx)
// // 	if err != nil {
// // 		log.Panic(err)
// // 	}
// // 	fmt.Printf("%+v\n", users)
// // }
