package database

import (
	"gorm-terminal/internal/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitializeDB はデータベース接続を初期化し、マイグレーションを実行します。
func InitializeDB(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})//Openは指定されたデータベースに接続する
    if err != nil {
        return nil, err
    }

    // モデルに基づいてテーブルを自動生成
    err = db.AutoMigrate(
        &model.User{},  // Userモデルを指定
        // 他のモデルもここに追加
    )
    if err != nil {
        return nil, err
    }

    return db, nil
}
