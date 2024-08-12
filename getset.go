package main

import (
	"errors"
	"time"
)

type Date struct {
	day, month, year int // private fields
}

// Public setter methods to control access
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day value")
	}

	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid day value")
	}

	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {
	if year < 1876 || year > time.Now().Year() {
		return errors.New("invalid day value")
	}

	d.year = year
	return nil
}

// Public getter methods
func (d *Date) GetDay() int {
	return d.day
}

func (d *Date) GetMonth() int {
	return d.month
}

func (d *Date) GetYear() int {
	return d.year
}
