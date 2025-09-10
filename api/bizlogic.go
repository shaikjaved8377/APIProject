package api

import (
	"Project/dataservice"
	"Project/model"
	"Project/queue"
	"database/sql"
	"fmt"

	"github.com/IBM/sarama"
)

type IbizLogic interface {
	CreateBookLogic(book model.Book) error
}

type bizlogic struct {
	DB       *sql.DB
	Producer sarama.SyncProducer
}

func NewBizLogic(db *sql.DB, producer sarama.SyncProducer) *bizlogic {
	return &bizlogic{DB: db, Producer: producer}
}

func (bl *bizlogic) CreateBookLogic(book model.Book) error {
	//validations
	if book.Title == "" {
		return fmt.Errorf("title should be present")
	}
	if err := dataservice.CreateBook(bl.DB, book); err != nil {
		return err
	}

	// produce the message to kafka
	message := fmt.Sprintf("Book created: %s by %s", book.Title, book.Author)
	err := queue.ProduceKafkaMessage("book_created_topic", message, bl.Producer)
	if err != nil {
		return fmt.Errorf("failed to produce kafka message: %v", err)
	}
	return nil
}
