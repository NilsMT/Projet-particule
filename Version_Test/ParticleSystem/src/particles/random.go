package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

func GimmeRandom() (posx, posy, scalex, scaley, colorR, colorG, colorB, op, SX, SY, rot float64) {
	rand.Seed(time.Now().UnixNano())
	//Variable modifiables
	scalex = config.General.DefaultSizeX
	scaley = config.General.DefaultSizeY
	colorR = config.General.DefaultR
	colorG = config.General.DefaultG
	colorB = config.General.DefaultB
	SX = config.General.DefaultSX
	SY = config.General.DefaultSY
	posx = float64(config.General.SpawnX)
	posy = float64(config.General.SpawnY)
	op = config.General.DefaultOpacity
	rot = config.General.DefaultRotation

	if config.General.RandomSpawn {
		posx = rand.Float64() * float64(config.General.WindowSizeX)
		posy = rand.Float64() * float64(config.General.WindowSizeY)
	}

	if config.General.RandomSizeX {
		scalex = rand.Float64() * config.General.MaxSizeX
	}

	if config.General.RandomSizeY {
		scalex = rand.Float64() * config.General.MaxSizeX
	}

	if config.General.RandomColor {
		colorR = rand.Float64()
		colorG = rand.Float64()
		colorB = rand.Float64()
	}

	if config.General.RandomOpacity {
		op = rand.Float64()
	}

	if config.General.RandomSpeedX {
		SX = float64(rand.Float64()-0.5) * config.General.MaxSpeedX
	}

	if config.General.RandomSpeedY {
		SY = float64(rand.Float64()-0.5) * config.General.MaxSpeedY
	}
	if config.General.RandomRotation {
		rot = float64(rand.Float64()-0.5) * config.General.MaxRotation
	}

	return posx, posy, scalex, scaley, colorR, colorG, colorB, op, SX, SY, rot
}
