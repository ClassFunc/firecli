# Singleton Firebase Admin Client

## Install:

```sh
go get -u github.com/classfunc/firecli
```

## Usage:

```go
package main

import (
    "github.com/classfunc/firecli"
)

func main() {

    // init firebase client
    firecli.WithCredentialsFile("SERVICE_ACCOUNT_FILE_PATH")
    
    // then you can use
    firecli.DB.Collection(...)
    firecli.Auth.CreateUser(...)
    firecli.DefaultBucket.Object("media/" + fileName)
    firecli.Storage.Bucket("name").Object("media/" + fileName)
	
	
    // settings firestore client
    firecli.EnvCollPrefix = "GO_COL_"
    fmt.Println(firecli.CollPathFromEnv("GO_COL_USERS"))
    
    // actions with firestore collections
    firecli.CollRef("GO_COL_USERS" /* or "users"*/)
    firecli.CollGetAllDocs("users")
    firecli.CollGetAllDocsData("users")
    
    // actions with firestore documents
    firecli.DocRef("users", "123")
    firecli.GetDoc("users", "123")
    firecli.GetDocData("users", "123")
    firecli.GetDocDataTo("users", "123", nil)
    
    // actions with firestore sub_collections
    firecli.SubCollRef("users", "123", "roles")
    firecli.SubCollGetAllDocs("users", "123", "roles")
    firecli.SubCollGetAllDocsData("users", "123", "roles")
    
    // actions with firestore sub_documents
    firecli.SubDocRef("users", "123", "roles", "456")
    firecli.GetSubDoc("users", "123", "roles", "456")
    firecli.GetSubDocData("users", "123", "roles", "456")

	type User struct {
        Name string
    }
    var adam = new(User)
    firecli.GetDocDataTo("users", "123", adam)
}
```

## Lisence:
MIT

## Authors:
ClassFunc Softwares JSC

Website: https://classfunc.com
