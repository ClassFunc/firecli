package main

import (
	"fmt"
	"github.com/classfunc/firecli"
)

const serviceAccountFilePath = "./YOUR_SERVICE_ACCOUNT_KEY.json"

func main() {
	firecli.WithCredentialsFile(serviceAccountFilePath)
	fmt.Println(
		firecli.App,
		firecli.Auth,
		firecli.Firestore,
		firecli.Storage,
		firecli.DefaultBucket,
	)

	firecli.EnvCollPrefix = "GO_COL_"
	fmt.Println(firecli.CollPathFromEnv("GO_COL_USERS"))

	firecli.CollRef("GO_COL_USERS" /* or "users"*/)
	firecli.CollGetAllDocs("users")
	firecli.CollGetAllDocsData("users")

	firecli.DocRef("users", "123")
	firecli.GetDoc("users", "123")
	firecli.GetDocData("users", "123")
	firecli.GetDocDataTo("users", "123", nil)

	firecli.SubCollRef("users", "123", "roles")
	firecli.SubCollGetAllDocs("users", "123", "roles")
	firecli.SubCollGetAllDocsData("users", "123", "roles")

	firecli.SubDocRef("users", "123", "roles", "456")
	firecli.GetSubDoc("users", "123", "roles", "456")
	firecli.GetSubDocData("users", "123", "roles", "456")

	type User struct {
		Name string
	}
	var adam = new(User)
	firecli.GetDocDataTo("users", "123", adam)

	// for collectionGroup
	// https://firebase.google.com/docs/firestore/query-data/queries?authuser=0&hl=en#collection-group-query
	firecli.CollGroupRef("CollId_or_subCollId")
	firecli.CollGroupGetAllDocs("CollId_or_subCollId")
	firecli.CollGroupGetAllDocsData("CollId_or_subCollId")
}
