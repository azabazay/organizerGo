package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	types "github.com/abai/organizer/types"
	"gorm.io/gorm"
)

type Service interface {
	// Example service methods.
	GetCatFact(context.Context) (*types.CatFact, error)

	// User Service methdos.
	GetUser(userID string) (*types.User, error)

	// Event Service methods.
	CreateEvent(event *types.TimeTableItem) (*types.TimeTableItem, error)
	DeleteEvent(eventID string) (*types.TimeTableItem, error)
	GetEvent(eventID string) (*types.TimeTableItem, error)
	GetUserEvents(eventID string) (*[]types.TimeTableItem, error)
}

type UserService struct {
	catFactUrl string
	DB         *gorm.DB
}

func NewUserService(catFactUrl string, db *gorm.DB) Service {
	return &UserService{
		catFactUrl: catFactUrl,
		DB:         db,
	}
}

func (s *UserService) GetCatFact(context.Context) (*types.CatFact, error) {
	resp, err := http.Get(s.catFactUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	fact := &types.CatFact{}
	if err := json.NewDecoder(resp.Body).Decode(fact); err != nil {
		return nil, err
	}

	return fact, nil
}

func (s *UserService) GetUser(userID string) (*types.User, error) {
	eventModel := &types.User{}

	err := s.DB.Where("id = ?", userID).First(eventModel).Error
	if err != nil {
		return nil, err
	}

	return eventModel, nil
}

func (s *UserService) CreateEvent(event *types.TimeTableItem) (*types.TimeTableItem, error) {
	err := s.DB.Create(&event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *UserService) DeleteEvent(eventID string) (*types.TimeTableItem, error) {
	eventModel := &types.TimeTableItem{}

	err := s.DB.Delete(eventModel, eventID).Error
	if err != nil {
		return nil, err
	}

	return eventModel, nil
}

func (s *UserService) GetEvent(eventID string) (*types.TimeTableItem, error) {
	eventModel := &types.TimeTableItem{}

	err := s.DB.Where("id = ?", eventID).First(eventModel).Error
	if err != nil {
		return nil, err
	}

	return eventModel, nil
}

func (s *UserService) GetUserEvents(userID string) (*[]types.TimeTableItem, error) {
	eventModel := &[]types.TimeTableItem{}
	idInt, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, err
	}

	err = s.DB.Find(eventModel, types.TimeTableItem{
		UserID: uint(idInt),
	}).Error
	if err != nil {
		return nil, err
	}

	return eventModel, nil
}
