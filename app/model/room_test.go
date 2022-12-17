package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/google/go-cmp/cmp"
)

func TestResult(t *testing.T) {

	// 比較を行うため事前に生成しておく
	uid1 := uuid.New()
	uid2 := uuid.New()

	time1 := time.Now()

	tests := []struct{
		name string
		room *RoomDetails
		want []*Result
	}{
		{
			name: "empty case",
			room: &RoomDetails{},
			want: nil,
		},
		{
			name: "one event",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id: uid1,
						Name: "event1",
						EventType: EventTypeOuter,
						EventAt: time1,
						Txns: []*Transaction{
							{
								Id: uid2,
								Amount: 100,
								Payer: uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount: 100,
					Payer: uid1,
					Receiver: uid2,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.room.Results()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("differs: (-want +got) %s", diff)
			}
		})
	}
}
