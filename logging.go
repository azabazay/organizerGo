package main

import (
	"context"
	"fmt"
	"time"

	types "github.com/abai/organizer/types"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *types.CatFact, err error) {
	defer func(start time.Time) {
		fmt.Printf("fact=%s err=%v took=%v\n", fact.Fact, err, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}

func (s *LoggingService) GetUser(userID string) (user *types.User, err error) {
	defer func(start time.Time) {
		fmt.Printf("userID=%s err=%v took=%v\n", userID, err, time.Since(start))
	}(time.Now())

	return s.next.GetUser(userID)
}

func (s *LoggingService) CreateEvent(event *types.TimeTableItem) (user *types.TimeTableItem, err error) {
	defer func(start time.Time) {
		fmt.Printf("event=%v err=%v took=%v\n", event, err, time.Since(start))
	}(time.Now())

	return s.next.CreateEvent(event)
}

func (s *LoggingService) DeleteEvent(eventID string) (user *types.TimeTableItem, err error) {
	defer func(start time.Time) {
		fmt.Printf("eventID=%s err=%v took=%v\n", eventID, err, time.Since(start))
	}(time.Now())

	return s.next.DeleteEvent(eventID)
}

func (s *LoggingService) GetEvent(eventID string) (user *types.TimeTableItem, err error) {
	defer func(start time.Time) {
		fmt.Printf("eventID=%s err=%v took=%v\n", eventID, err, time.Since(start))
	}(time.Now())

	return s.next.GetEvent(eventID)
}

func (s *LoggingService) GetUserEvents(userID string) (user *[]types.TimeTableItem, err error) {
	defer func(start time.Time) {
		fmt.Printf("userID=%s err=%v took=%v\n", userID, err, time.Since(start))
	}(time.Now())

	return s.next.GetUserEvents(userID)
}
