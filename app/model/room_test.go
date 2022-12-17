package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestResult(t *testing.T) {

	// 比較を行うため事前に生成しておく
	uid1 := uuid.New()
	uid2 := uuid.New()
	uid3 := uuid.New()
	uid4 := uuid.New()
	uid5 := uuid.New()
	uid6 := uuid.New()
	// uid7 := uuid.New()
	// uid8 := uuid.New()
	// uid9 := uuid.New()

	time1 := time.Now()

	tests := []struct {
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
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   100,
					Payer:    uid1,
					Receiver: uid2,
				},
			},
		},
		{
			name: "two events",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   200,
					Payer:    uid1,
					Receiver: uid2,
				},
			},
		},
		{
			name: "two events with different payer",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid2,
								Receiver: uid1,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: nil,
		},
		{
			name: "two events with different receiver",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid3,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   100,
					Payer:    uid1,
					Receiver: uid2,
				}, {
					Amount:   100,
					Payer:    uid1,
					Receiver: uid3,
				},
			},
		},
		{
			name: "two events with different amount",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   200,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   300,
					Payer:    uid1,
					Receiver: uid2,
				},
			},
		},
		{
			name: "two events with different payer",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid3,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   100,
					Payer:    uid1,
					Receiver: uid2,
				}, {
					Amount:   100,
					Payer:    uid3,
					Receiver: uid2,
				},
			},
		},
		{
			name: "three events with same payer and receiver",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event3",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   300,
					Payer:    uid1,
					Receiver: uid2,
				},
			},
		},
		{
			name: "three events with same payer and different receivers",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid3,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event3",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid4,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   100,
					Payer:    uid1,
					Receiver: uid2,
				}, {
					Amount:   100,
					Payer:    uid1,
					Receiver: uid3,
				}, {
					Amount:   100,
					Payer:    uid1,
					Receiver: uid4,
				},
			},
		},
		{
			name: "three events with different payers and same receivers",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid3,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event3",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid4,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   100,
					Payer:    uid1,
					Receiver: uid2,
				}, {
					Amount:   100,
					Payer:    uid3,
					Receiver: uid2,
				}, {
					Amount:   100,
					Payer:    uid4,
					Receiver: uid2,
				},
			},
		},
		{
			name: "three events with different payers and different receivers",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid3,
								Receiver: uid4,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event3",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid5,
								Receiver: uid6,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: []*Result{
				{
					Amount:   100,
					Payer:    uid1,
					Receiver: uid2,
				}, {
					Amount:   100,
					Payer:    uid3,
					Receiver: uid4,
				}, {
					Amount:   100,
					Payer:    uid5,
					Receiver: uid6,
				},
			},
		},
		{
			name: "three events with different payers and different receivers in three members",
			room: &RoomDetails{
				Events: []*Event{
					{
						Id:        uid1,
						Name:      "event1",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid1,
								Receiver: uid2,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event2",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid2,
								Receiver: uid3,
							},
						},
						CreatedAt: time1,
					},
					{
						Id:        uid1,
						Name:      "event3",
						EventType: EventTypeOuter,
						EventAt:   time1,
						Txns: []*Transaction{
							{
								Id:       uid2,
								Amount:   100,
								Payer:    uid3,
								Receiver: uid1,
							},
						},
						CreatedAt: time1,
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.room.Results()
			// if diff := cmp.Diff(tt.want, got); diff != "" {
			// 	t.Errorf("differs: (-want +got) %s", diff)
			// }
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("differs: (-want +got) %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
