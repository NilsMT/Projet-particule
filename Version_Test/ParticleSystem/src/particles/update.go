package particles

import (
	"math/rand"
	"project-particles/config"
	"time"
)

var count float64 = 0
var offsetX float64 = 10
var offsetY float64 = 10

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	rand.Seed(time.Now().UnixNano())
	//génération
	if config.General.SpawnRate > 0 {
		count += config.General.SpawnRate
		//spawn
		for count >= 1 {
			posx, posy, scalex, scaley, colorR, colorG, colorB, op, SX, SY, rot := GimmeRandom()
			//génération de la particule
			s.Content.PushFront(&Particle{
				PositionX: posx,
				PositionY: posy,
				ScaleX:    scalex, ScaleY: scaley,
				ColorRed: colorR, ColorGreen: colorG, ColorBlue: colorB,
				Opacity:  op,
				SpeedX:   SX,
				SpeedY:   SY,
				Rotation: rot,
			})
			count -= 1
		}
	}
	//
	//parcours la liste et change toute les positions
	for e := s.Content.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Particle)
		//extensions
		//debug color (direction)
		if config.General.StatusDebug {
			if v.SpeedX > 0 && v.SpeedY > 0 { //rouge ++
				v.ColorRed, v.ColorGreen, v.ColorBlue = 1, 0, 0
			} else if v.SpeedX > 0 && v.SpeedY < 0 { //bleu +-
				v.ColorRed, v.ColorGreen, v.ColorBlue = 0, 0, 1
			} else if v.SpeedX < 0 && v.SpeedY < 0 { //vert --
				v.ColorRed, v.ColorGreen, v.ColorBlue = 0, 1, 0
			} else if v.SpeedX < 0 && v.SpeedY > 0 { //jaune -+
				v.ColorRed, v.ColorGreen, v.ColorBlue = 1, 1, 0
			} else if v.SpeedX == 0 { //cyan 0?
				v.ColorRed, v.ColorGreen, v.ColorBlue = 0, 1, 1
			} else if v.SpeedX == 0 { //rose ?0
				v.ColorRed, v.ColorGreen, v.ColorBlue = 1, 0, 1
			} else { //blanc 00
				v.ColorRed, v.ColorGreen, v.ColorBlue = 1, 1, 1
			}
		}
		//rebond

		partsizeX := 10 * v.ScaleX
		partsizeY := 10 * v.ScaleY
		if config.General.Bounce {
			if v.PositionX <= 0 || v.PositionX >= float64(config.General.WindowSizeX)-partsizeX {
				v.SpeedX = -(v.SpeedX)
				//debug color (rebord)
				if config.General.StatusDebug {
					v.ColorRed, v.ColorGreen, v.ColorBlue = 1, 0.4, 0 //orange
				}
			}
			if v.PositionY <= 0 || v.PositionY >= float64(config.General.WindowSizeY)-partsizeY {
				v.SpeedY = -(v.SpeedY)
				//debug color (rebord)
				if config.General.StatusDebug {
					v.ColorRed, v.ColorGreen, v.ColorBlue = 1, 0.4, 0 //orange
				}
				v.SpeedY = v.SpeedY / 1.1
			} else {
				if config.General.Gravity {
					v.SpeedY += config.General.GravityForce
				}
			}
		}
		//gravité
		//ajout à pos = vers le bas

		//rotation
		if config.General.ContinuousRotation {
			v.Rotation += v.Rotation / 10
		}
		// par défaut
		//mouvement
		if v.PositionY > 0 {
			v.PositionY += v.SpeedY
		}
		if v.PositionX > 0 {
			v.PositionX += v.SpeedX
		}
	}
}
