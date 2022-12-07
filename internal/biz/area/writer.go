package area

import "time"

type Writer interface {
	SetRegionID(regionID string)
	SetTitle(title string)
	SetCityCode(cityCode string)
	SetZipCode(zipCode string)
	SetLevel(level int)
	SetUpdatedAt(t time.Time)
}

type Write func(w Writer)

func SetRegionID(regionID string) Write {
	return func(w Writer) {
		w.SetRegionID(regionID)
	}
}

func SetTitle(title string) Write {
	return func(w Writer) {
		w.SetTitle(title)
	}
}

func SetCityCode(cityCode string) Write {
	return func(w Writer) {
		w.SetCityCode(cityCode)
	}
}

func SetZipCode(zipCode string) Write {
	return func(w Writer) {
		w.SetZipCode(zipCode)
	}
}

func SetLevel(level int) Write {
	return func(w Writer) {
		w.SetLevel(level)
	}
}

func SetUpdatedAt(t time.Time) Write {
	return func(w Writer) {
		w.SetUpdatedAt(t)
	}
}
