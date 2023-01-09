package particles

import (
	"project-particles/config"
	"time"
	"math/rand"
)

func GimmeRandom() (posx,posy,scalex,scaley,colorR,colorG,colorB,op,VX,VY float64) {
	rand.Seed(time.Now().UnixNano())
	//Variable modifiables
	scalex = 1
	scaley = 1
	colorR = 1
	colorG = 1
	colorB = 0
	VX = float64(rand.Float64()-0.5)
	VY = float64(rand.Float64()-0.5)
	posx = float64(config.General.SpawnX)
	posy = float64(config.General.SpawnY)
	op = 1
	if config.General.RandomSpawn {
		posx=rand.Float64()*float64(config.General.WindowSizeX)
		posy=rand.Float64()*float64(config.General.WindowSizeY)
	}
	return posx,posy,scalex,scaley,colorR,colorG,colorB,op,VX,VY
}