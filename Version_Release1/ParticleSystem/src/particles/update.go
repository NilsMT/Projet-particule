package particles

import (
	"project-particles/config"
	"math/rand"
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
		for count>=1 {
			posx,posy,scalex,scaley,colorR,colorG,colorB,op,VX,VY := GimmeRandom()
			//génération de la particule
			s.Content.PushFront(&Particle{
				PositionX: posx,
				PositionY: posy,
				ScaleX:    scalex, ScaleY: scaley,
				ColorRed: colorR, ColorGreen: colorG, ColorBlue: colorB,
				Opacity:   op,
				VelocityX: VX,
				VelocityY: VY,
			})
			count-=1
		}
	}
	//
	//parcours la liste et change toute les positions
	for e := s.Content.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Particle)
		v.PositionY += v.VelocityY
		v.PositionX += v.VelocityX
	}
}
