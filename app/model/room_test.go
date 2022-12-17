package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestResult(t *testing.T) {
	tests := []RoomDetails{
		RoomDetails{
			Id:        uuid.New(),
			Name:      "test",
			CreatedAt: time.Now(),
			Members: []*Member{
				&Member{
					Id:   uuid.New(),
					Name: "test1",
				}, {
					Id:   uuid.New(),
					Name: "test2",
				},
			},
			Events: []*Event{
				&Event{
					Id:        uuid.New(),
					Name:      "test",
					Amount:    100,
					EventType: EventTypeOuter,
					EventAt:   time.Now(),
					Txns: []*Transaction{
						&Transaction{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
						{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
					},
					CreatedAt: time.Now(),
				},
				{
					Id:        uuid.New(),
					Name:      "test",
					Amount:    100,
					EventType: EventTypeInner,
					EventAt:   time.Now(),
					Txns: []*Transaction{
						&Transaction{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
						{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
					},
				},
			},
		}, {
			Id:        uuid.New(),
			Name:      "test",
			CreatedAt: time.Now(),
			Members: []*Member{
				&Member{
					Id:   uuid.New(),
					Name: "test1",
				}, {
					Id:   uuid.New(),
					Name: "test2",
				},
			},
			Events: []*Event{
				&Event{
					Id:        uuid.New(),
					Name:      "test",
					Amount:    100,
					EventType: EventTypeOuter,
					EventAt:   time.Now(),
					Txns: []*Transaction{
						&Transaction{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
						{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
					},
					CreatedAt: time.Now(),
				},
				{
					Id:        uuid.New(),
					Name:      "test",
					Amount:    100,
					EventType: EventTypeInner,
					EventAt:   time.Now(),
					Txns: []*Transaction{
						&Transaction{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
						{
							Id:       uuid.New(),
							Amount:   100,
							Payer:    uuid.New(),
							Receiver: uuid.New(),
						},
					},
				},
			},
		},
	}

	if got := tests.Results(); got != 0 {
		t.Errorf("got: %d, want: 0", got)
	}
}
