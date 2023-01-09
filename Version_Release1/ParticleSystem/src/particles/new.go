package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()
	for i := 0; i < config.General.InitNumParticles; i++ {
		posx,posy,scalex,scaley,colorR,colorG,colorB,op,VX,VY := GimmeRandom()
		//génération de la particule
		l.PushFront(&Particle{
			PositionX: posx,
			PositionY: posy,
			ScaleX:    scalex, ScaleY: scaley,
			ColorRed: colorR, ColorGreen: colorG, ColorBlue: colorB,
			Opacity:   op,
			VelocityX: VX,
			VelocityY: VY,
		})
	}
	return System{Content: l}
}
