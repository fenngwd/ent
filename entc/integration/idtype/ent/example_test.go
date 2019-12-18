// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.
	u0 := client.User.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u0)
	u2 := client.User.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("user created:", u2)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetName("string").
		SetSpouse(u0).
		AddFollowing(u2).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	u0, err = u.QuerySpouse().First(ctx)
	if err != nil {
		log.Fatalf("failed querying spouse: %v", err)
	}
	log.Println("spouse found:", u0)

	u2, err = u.QueryFollowing().First(ctx)
	if err != nil {
		log.Fatalf("failed querying following: %v", err)
	}
	log.Println("following found:", u2)

	// Output:
}
