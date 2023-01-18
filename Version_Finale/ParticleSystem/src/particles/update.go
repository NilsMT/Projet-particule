package particles

import (
	"project-particles/config"
)

var count float64 = 0
var deathcount int = 0

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	//génération
	if config.General.SpawnRate > 0 {
		count += config.General.SpawnRate
		//spawn
		for count >= 1 {
			rot, posx, posy, scalex, scaley, op, colorR, colorG, colorB, DX, DY, lf, dead := GimmeRandom()
			//génération de la particule
			if deathcount <= 0 { //push normalement
				s.Content.PushFront(&Particle{
					PositionX: posx,
					PositionY: posy,
					ScaleX:    scalex, ScaleY: scaley,
					ColorRed: colorR, ColorGreen: colorG, ColorBlue: colorB,
					Opacity:    op,
					DirectionX: DX,
					DirectionY: DY,
					Lifetime:   lf,
					Rotation:   rot,
					Dead:       dead,
				})
			} else { //va en back, modifie et move front
				j := s.Content.Back()
				u := j.Value.(*Particle)
				u.PositionX, u.PositionY, u.ScaleX, u.ScaleY = posx, posy, scalex, scaley
				u.ColorRed, u.ColorGreen, u.ColorBlue, u.Opacity = colorR, colorG, colorB, op
				u.DirectionX, u.DirectionY, u.Lifetime, u.Dead = DX, DY, lf, dead
				deathcount -= 1
				s.Content.MoveToFront(j)
			}
			e := s.Content.Front()
			KillIfNeeded(s, e, &deathcount)
			count -= 1
		}
	}
	//
	//parcours la liste et change toute les positions
	for e := s.Content.Front(); e != nil; {
		next := e.Next()
		v := e.Value.(*Particle)
		if !v.Dead {
			//extension
			//gravité
			if config.General.GravityEnabled {
				v.DirectionY += config.General.GravityForceY
				v.DirectionX += config.General.GravityForceX
			}
			//temps de vie
			if config.General.LifetimeEnabled {
				if v.Lifetime > 0 {
					v.Lifetime -= 1
					v.Opacity -= v.Opacity / float64(v.Lifetime)
				} else {
					v.Lifetime = 0
				}
				//quand temps de vie expire il commence à devenir transparent
			}
			//decay
			if config.General.DecayEnabled {
				Decaying(v)
			}
			//par défault
			v.PositionY += v.DirectionY
			v.PositionX += v.DirectionX
			//kill
			KillIfNeeded(s, e, &deathcount)
		}
		e = next
	}
}
