package main

import (
	"fmt"
	"image/color"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var marge int = 15

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 20, 255})
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		//Next := e.Next()
		p, ok := e.Value.(*particles.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
			if config.General.Debug { //ajout des vecteurs de directions
				if !p.Dead {
					ebitenutil.DrawLine(screen, p.PositionX+10*p.ScaleX/2, p.PositionY+10*p.ScaleY/2, p.PositionX+(p.DirectionX*50), p.PositionY+(p.DirectionY*50), color.RGBA{uint8(p.ColorRed * 255), uint8(p.ColorGreen * 255), uint8(p.ColorBlue * 255), uint8(p.Opacity * 255)})
				}
			}
		}
		//e = Next
	}
	if config.General.KeyInput { //gestion de configuration dynamique
		particles.ManageDynamicEdit(screen)
	}
	if config.General.Debug { //affichage debug
		ebitenutil.DebugPrintAt(screen, "----------------------------------", 5, 15)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("FPS: ", ebiten.ActualTPS()), 5, 25)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombre de particules: ", g.system.Content.Len()), 5, 40)
		ebitenutil.DebugPrintAt(screen, "----------------------------------", 5, 50)
	}
	ebitenutil.DebugPrintAt(screen, "----------------------------------", 5, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Etat de la configuration: ", config.General.ShowConfigStatus), 5, 60)
	ebitenutil.DebugPrintAt(screen, "----------------------------------", 5, 70)
	if config.General.ShowConfigStatus { //affichage du statut de la config
		decallageconfig := 65
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("DebugEnabled: ", config.General.Debug), 5, marge*1+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnType: ", config.General.SpawnType), 5, marge*2+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnX: ", config.General.SpawnX, " SpawnY: ", config.General.SpawnY), 5, marge*3+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("SpawnRate: ", config.General.SpawnRate), 5, marge*4+decallageconfig)
		//gravity
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityEnabled: ", config.General.GravityEnabled), 5, marge*5+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("GravityForceX: ", config.General.GravityForceX, " GravityForceY: ", config.General.GravityForceY), 5, marge*6+decallageconfig)
		//border
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("OutsideForbidden: ", config.General.OutsideForbidden), 5, marge*7+decallageconfig)
		//lifetime
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("LifetimeEnabled: ", config.General.LifetimeEnabled), 5, marge*8+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("LifetimeMax: ", config.General.LifetimeMax, " LifetimeMin: ", config.General.LifetimeMin), 5, marge*9+decallageconfig)
		//decay
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("VariationEnabled: ", config.General.DecayEnabled), 5, marge*10+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("ScaleX: ", config.General.DecayScaleX, " ScaleY: ", config.General.DecayScaleY), 5, marge*11+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("couleur (RGB): ", config.General.DecayColorR, config.General.DecayColorG, config.General.DecayColorB), 5, marge*12+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Opacity: ", config.General.DecayOpacity, " Rotation: ", config.General.DecayRotation), 5, marge*13+decallageconfig)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Temps pour l'atteinte de l'objectif: ", config.General.DecayTime), 5, marge*14+decallageconfig)
	}
}
