package repository

import (
    "auth/src/infrastructure/mongo"
)

const (
    tableName := "user"
    pk := "UserId"
)

func SetDefault() {
    mongo.SetDefault(tableName, pk)
}