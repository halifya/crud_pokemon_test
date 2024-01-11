package models

import (
	"time"

	"github.com/google/uuid"
)

type Pokemon struct {
	ID            uuid.UUID      `gorm:"type:char(36);primary_key;default:uuid()" json:"id,omitempty"`
	Num           string         `gorm:"type:varchar(100)" json:"num,omitempty"`
	Name          string         `gorm:"type:varchar(100)" json:"name,omitempty"`
	Img           string         `gorm:"type:varchar(255)" json:"img,omitempty"`
	Type          StringArray    `gorm:"type:varchar(100)" json:"type,omitempty"`
	Height        string         `gorm:"type:varchar(100)" json:"height,omitempty"`
	Weight        string         `gorm:"type:varchar(100)" json:"weight,omitempty"`
	Candy         string         `gorm:"type:varchar(100)" json:"candy,omitempty"`
	Egg           string         `gorm:"type:varchar(100)" json:"egg,omitempty"`
	Multipliers   FloatArray     `gorm:"type:varchar(100)" json:"multipliers,omitempty"`
	Weaknesses    StringArray    `gorm:"type:varchar(100)" json:"weaknesses,omitempty"`
	CandyCount    int            `gorm:"type:int" json:"candy_count,omitempty"`
	SpawnChance   float64        `gorm:"type:float" json:"spawn_chance,omitempty"`
	AvgSpawns     float64        `gorm:"type:float" json:"avg_spawns,omitempty"`
	SpawnTime     string         `gorm:"type:varchar(100)" json:"spawn_time,omitempty"`
	PrevEvolution EvolutionArray `gorm:"type:varchar(255)" json:"prev_evolution,omitempty"`
	NextEvolution EvolutionArray `gorm:"type:varchar(255)" json:"next_evolution,omitempty "`
	CreatedAt     time.Time      `gorm:"type:timestamp;not null" json:"createdAt,omitempty"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;not null" json:"updatedAt,omitempty"`
}

type Evolution struct {
	Num  string `gorm:"type:varchar(50)" json:"num"`
	Name string `gorm:"type:varchar(50)" json:"name"`
}

type CreatePokemonSchema struct {
	Num           string      `gorm:"type:varchar(100)" json:"num,omitempty"`
	Name          string      `gorm:"type:varchar(100)" json:"name,omitempty"`
	Img           string      `gorm:"type:varchar(255)" json:"img,omitempty"`
	Type          []string    `gorm:"type:varchar(100)" json:"type,omitempty"`
	Height        string      `gorm:"type:varchar(100)" json:"height,omitempty"`
	Weight        string      `gorm:"type:varchar(100)" json:"weight,omitempty"`
	Candy         string      `gorm:"type:varchar(100)" json:"candy,omitempty"`
	Egg           string      `gorm:"type:varchar(100)" json:"egg,omitempty"`
	Multipliers   []float64   `gorm:"type:varchar(100)" json:"multipliers,omitempty"`
	Weaknesses    []string    `gorm:"type:varchar(100)" json:"weaknesses,omitempty"`
	CandyCount    int         `gorm:"type:int" json:"candy_count,omitempty"`
	SpawnChance   float64     `gorm:"type:float" json:"spawn_chance,omitempty"`
	AvgSpawns     float64     `gorm:"type:float" json:"avg_spawns,omitempty"`
	SpawnTime     string      `gorm:"type:varchar(100)" json:"spawn_time,omitempty"`
	PrevEvolution []Evolution `json:"prev_evolution"`
	NextEvolution []Evolution `json:"next_evolution"`
}

type UpdatePokemonSchema struct {
	Num           string         `gorm:"type:varchar(100)" json:"num,omitempty"`
	Name          string         `gorm:"type:varchar(100)" json:"name,omitempty"`
	Img           string         `gorm:"type:varchar(255)" json:"img,omitempty"`
	Type          StringArray    `gorm:"type:varchar(100)" json:"type,omitempty"`
	Height        string         `gorm:"type:varchar(100)" json:"height,omitempty"`
	Weight        string         `gorm:"type:varchar(100)" json:"weight,omitempty"`
	Candy         string         `gorm:"type:varchar(100)" json:"candy,omitempty"`
	Egg           string         `gorm:"type:varchar(100)" json:"egg,omitempty"`
	Multipliers   FloatArray     `gorm:"type:float" json:"multipliers,omitempty"`
	Weaknesses    StringArray    `gorm:"type:varchar(100)" json:"weaknesses,omitempty"`
	CandyCount    int            `gorm:"type:int" json:"candy_count,omitempty"`
	SpawnChance   float64        `gorm:"type:float" json:"spawn_chance,omitempty"`
	AvgSpawns     float64        `gorm:"type:float" json:"avg_spawns,omitempty"`
	SpawnTime     string         `gorm:"type:varchar(100)" json:"spawn_time,omitempty"`
	PrevEvolution EvolutionArray `json:"prev_evolution"`
	NextEvolution EvolutionArray `json:"next_evolution"`
	UpdatedAt     time.Time      `gorm:"type:timestamp;not null" json:"updatedAt,omitempty"`
}
