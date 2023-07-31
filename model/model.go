package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title  string
	Genre  string
	Rating float32
}

func CreateMovie(db *gorm.DB, movie *Movie) error {
	return db.Create(movie).Error
}

func GetMovies(db *gorm.DB, movies *[]Movie) error {
	return db.Find(movies).Error
}

func GetMovie(db *gorm.DB, movie *Movie, id string) error {
	return db.Where("id = ?", id).First(movie).Error
}

func UpdateMovie(db *gorm.DB, movie *Movie) error {
	return db.Save(movie).Error
}

func DeleteMovie(db *gorm.DB, id string) error {
	return db.Where("id = ?", id).Delete(&Movie{}).Error
}
