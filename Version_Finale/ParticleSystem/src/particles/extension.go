package particles

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"project-particles/config"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Cette fonction effectue la gestion des attributs (Position,Direction,Couleurs,Opacité,Temps de vie)
// de la particule quand elle spawn selon les règles établie dans le fichier config.json
// elle retourne les valeurs que doit prendre les attributs
// elle gère aussi la mort au spawn, ce qui est utile si jamais la particule spawn en dehors de l'écran
func GimmeRandom() (rot, posx, posy, scalex, scaley, op, colorR, colorG, colorB, DX, DY, lf float64, dead bool) {
	//Partie graphique
	rand.Seed(time.Now().UnixNano())
	//Variable modifiables
	colorR = config.General.DefaultColorR
	colorG = config.General.DefaultColorG
	colorB = config.General.DefaultColorB
	DX = float64(rand.Float64() - 0.5)
	DY = float64(rand.Float64() - 0.5)
	lf = 1
	if config.General.LifetimeEnabled {
		minl := config.General.LifetimeMin
		maxl := config.General.LifetimeMax
		lf = minl + rand.Float64()*(maxl-minl)
	}
	//partie avec mort possible
	//scale+pos
	scalex = config.General.DefaultScaleX
	scaley = config.General.DefaultScaleY
	op = config.General.DefaultOpacity
	if config.General.SpawnType == "default" {
		posx = float64(config.General.SpawnX)
		posy = float64(config.General.SpawnY)
	} else if config.General.SpawnType == "random" {
		posy = rand.Float64() * float64(config.General.WindowSizeY)
		posx = rand.Float64() * float64(config.General.WindowSizeX)
	} else if config.General.SpawnType == "middle" {
		posx = float64(config.General.WindowSizeX / 2)
		posy = float64(config.General.WindowSizeY / 2)
	} else if config.General.SpawnType == "cursor" {
		a, b := ebiten.CursorPosition()
		posx, posy = float64(a), float64(b)
	}

	if config.General.Shape == "circle" {
		var angle float64 = rand.Float64() * math.Pi * 2
		prevposx, prevposy := posx, posy
		posx = math.Cos(angle)*config.General.ShapeSize + float64(prevposx)
		posy = math.Sin(angle)*config.General.ShapeSize + float64(prevposy)
	} else if config.General.Shape == "triangle" {
		prevposx, prevposy := posx, posy
		var Ax, Ay float64 = config.General.ShapeSize + float64(prevposx), float64(prevposy) + config.General.ShapeSize/2
		var Bx, By float64 = float64(prevposx) - config.General.ShapeSize, float64(prevposy) + config.General.ShapeSize/2
		var Cx, Cy float64 = float64(prevposx), float64(prevposy) - config.General.ShapeSize
		var r1, r2 float64 = rand.Float64(), rand.Float64()
		var sqrtR1 float64 = math.Sqrt(r1)

		posx, posy = (1-sqrtR1)*Ax+(sqrtR1*(1-r2))*Bx+(sqrtR1*r2)*Cx, (1-sqrtR1)*Ay+(sqrtR1*(1-r2))*By+(sqrtR1*r2)*Cy

	} else if config.General.Shape == "square" {
		prevposx, prevposy := posx, posy
		var angle1, angle2 = rand.Float64() * math.Pi * 2, rand.Float64() * math.Pi * 2
		posx, posy = math.Cos(angle1)*config.General.ShapeSize+float64(prevposx), math.Sin(angle2)*config.General.ShapeSize+float64(prevposy)

	}
	if config.General.OutsideForbidden {
		if !(IsInsideWindow(posx, posy, scalex, scaley)) {
			dead = true
		}
	}
	if config.General.RangeValue {
		//color
		min := config.General.DefaultColorG - config.General.RangeColorR
		max := config.General.DefaultColorG + config.General.RangeColorR
		colorR = min + rand.Float64()*(max-min)
		min = config.General.DefaultColorG - config.General.RangeColorG
		max = config.General.DefaultColorG + config.General.RangeColorG
		colorG = min + rand.Float64()*(max-min)
		min = config.General.DefaultColorB - config.General.RangeColorB
		max = config.General.DefaultColorB + config.General.RangeColorB
		colorB = min + rand.Float64()*(max-min)
		//opacity
		min = config.General.DefaultOpacity - config.General.RangeOpacity
		max = config.General.DefaultOpacity + config.General.RangeOpacity
		op = min + rand.Float64()*(max-min)
		//direction
		min = DY - config.General.RangeDirectionX
		max = DX + config.General.RangeDirectionX
		DX = min + rand.Float64()*(max-min)
		min = DY - config.General.RangeDirectionY
		max = DY + config.General.RangeDirectionY
		DY = min + rand.Float64()*(max-min)
		//scale
		min = config.General.DefaultScaleX - config.General.RangeScaleX
		max = config.General.DefaultScaleX + config.General.RangeScaleX
		scalex = min + rand.Float64()*(max-min)
		min = config.General.DefaultScaleY - config.General.RangeScaleY
		max = config.General.DefaultScaleY + config.General.RangeScaleY
		scaley = min + rand.Float64()*(max-min)
		//rotation
	}
	return rot, posx, posy, scalex, scaley, op, colorR, colorG, colorB, DX, DY, lf, dead
}

// La fonction détermine si la particule est contenu dans l'écran ou non
// elle retourne un booléen qui vaut vrai si la particule est dans l'écran, sinon elle retourne faux
// la particule est considéré comme dehors quand elle n'est plus visible (et non quand elle est partiellement dehors)
func IsInsideWindow(posx, posy, sx, sy float64) bool {
	screenx := float64(config.General.WindowSizeX)
	screeny := float64(config.General.WindowSizeY)
	if posx < 0-sx || posx > screenx+sx || posy < 0-sy || posy > screeny+sy {
		return false
	}
	return true
}

// La fonction decaying prend en paramètre une particule
// Elle va faire la gestion des attributs de la particule en fonction du temps
// Les valeurs dans config sont les "objectifs" à atteindre en un certain temps
// Exemple :
// si config.General.DecayTime vaut 10 et si config.General.DecayColorR vaut 1
// il faut qu'en 10 frame, elle atteignent la couleur rouge à 1
// la fonction fonctionne comme le principe du tweening (utilisé pour des animations)
func Decaying(v *Particle) {
	//Rotation
	if v.Rotation > config.General.DecayRotation {
		v.Rotation -= (1 - config.General.DecayRotation) / config.General.DecayTime
	} else if v.Rotation < config.General.DecayRotation {
		v.Rotation += (config.General.DecayRotation) / config.General.DecayTime
	}
	//colors
	if v.ColorRed > config.General.DecayColorR {
		v.ColorRed -= (1 - config.General.DecayColorR) / config.General.DecayTime
	} else if v.ColorRed < config.General.DecayColorR {
		v.ColorRed += (config.General.DecayColorR) / config.General.DecayTime
	}
	if v.ColorGreen > config.General.DecayColorG {
		v.ColorGreen -= (1 - config.General.DecayColorG) / config.General.DecayTime
	} else if v.ColorGreen < config.General.DecayColorG {
		v.ColorGreen += (config.General.DecayColorG) / config.General.DecayTime
	}
	if v.ColorBlue > config.General.DecayColorB {
		v.ColorBlue -= (1 - config.General.DecayColorB) / config.General.DecayTime
	} else if v.ColorBlue < config.General.DecayColorB {
		v.ColorBlue += (config.General.DecayColorB) / config.General.DecayTime
	}
	//Opacity
	if v.Opacity > config.General.DecayOpacity {
		v.Opacity -= (1 - config.General.DecayOpacity) / config.General.DecayTime
	} else if v.Opacity < config.General.DecayOpacity {
		v.Opacity += (config.General.DecayOpacity) / config.General.DecayTime
	}
	//scale
	//X
	if v.ScaleX > config.General.DecayScaleX {
		v.ScaleX -= (1 - config.General.DecayScaleX) / config.General.DecayTime
	} else if v.ScaleX < config.General.DecayScaleX {
		v.ScaleX += (config.General.DecayScaleX) / config.General.DecayTime
	}
	//Y
	if v.ScaleY > config.General.DecayScaleY {
		v.ScaleY -= (1 - config.General.DecayScaleY) / config.General.DecayTime
	} else if v.ScaleY < config.General.DecayScaleY {
		v.ScaleY += (config.General.DecayScaleY) / config.General.DecayTime
	}
}

// initialisation des options et de la sélection
var selectedindextype int
var selectedindexbool int
var selectedindexint int
var selectedindexfloat int

// liste des options (triée par type)
var confbool []*bool = []*bool{&config.General.Debug, &config.General.GravityEnabled, &config.General.OutsideForbidden, &config.General.LifetimeEnabled, &config.General.RangeValue, &config.General.DecayEnabled, &config.General.ShowConfigStatus}
var confboolname []string = []string{"Debug", "GravityEnabled", "OutsideForbidden", "LifetimeEnabled", "DecayEnabled", "RangeValue", "ShowConfigStatus"}
var confint []*int = []*int{&config.General.SpawnX, &config.General.SpawnY}
var confintname []string = []string{"SpawnX", "SpawnY"}

var conffloat []*float64 = []*float64{&config.General.ShapeSize, &config.General.SpawnRate, &config.General.GravityForceX, &config.General.GravityForceY, &config.General.LifetimeMax, &config.General.LifetimeMin, &config.General.DecayTime, &config.General.DecayOpacity, &config.General.DecayScaleX, &config.General.DecayScaleY, &config.General.DecayColorR, &config.General.DecayColorG, &config.General.DecayColorB, &config.General.DecayRotation, &config.General.DefaultRotation, &config.General.DefaultColorR, &config.General.DefaultColorG, &config.General.DefaultColorB, &config.General.DefaultScaleX, &config.General.DefaultScaleY, &config.General.DefaultOpacity, &config.General.RangeRotation, &config.General.RangeDirectionX, &config.General.RangeDirectionY, &config.General.RangeScaleY, &config.General.RangeScaleX, &config.General.RangeColorR, &config.General.RangeColorG, &config.General.RangeColorB, &config.General.RangeOpacity}
var conffloatname []string = []string{"ShapeSize", "SpawnRate", "GravityForceX", "GravityForceY", "LifetimeMax", "LifetimeMin", "DecayTime", "DecayOpacity", "DecayScaleX", "DecayScaleY", "DecayColorR", "DecayColorG", "DecayColorB", "DecayRotation", "DefaultRotation", "DefaultColorR", "DefaultColorG", "DefaultColorB", "DefaultScaleX", "DefaultScaleY", "DefaultOpacity", "RangeRotation", "RangeDirectionX", "RangeDirectionY", "RangeScaleY", "RangeScaleX", "RangeColorR", "RangeColorG", "RangeColorB", "RangeOpacity"}
var stringcycle []string = []string{"default", "middle", "random", "cursor"}
var stringcycleindex int
var stringshcycle []string = []string{"default", "circle", "triangle", "square"}
var stringshcycleindex int
var typename []string = []string{"bool", "int", "float", "string", "string"}

// cette fonction permet de déterminé la sélection avec les touches
// elle évite d'avoir des erreurs en indexant à une valeur hors champs
func ManageSelection(add bool) {
	if add {
		if selectedindextype == 0 {
			if selectedindexbool == len(confbool)-1 {
				selectedindexbool = 0
			} else {
				selectedindexbool += 1
			}
		} else if selectedindextype == 1 {
			if selectedindexint == len(confint)-1 {
				selectedindexint = 0
			} else {
				selectedindexint += 1
			}
		} else if selectedindextype == 2 {
			if selectedindexfloat == len(conffloat)-1 {
				selectedindexfloat = 0
			} else {
				selectedindexfloat += 1
			}
		}
	} else {
		if selectedindextype == 0 {
			if selectedindexbool == 0 {
				selectedindexbool = len(confbool) - 1
			} else {
				selectedindexbool -= 1
			}
		} else if selectedindextype == 1 {
			if selectedindexint == 0 {
				selectedindexint = len(confint) - 1
			} else {
				selectedindexint -= 1
			}
		} else if selectedindextype == 2 {
			if selectedindexfloat == 0 {
				selectedindexfloat = len(conffloat) - 1
			} else {
				selectedindexfloat -= 1
			}
		}
	}
}

// cette fonction permet de déterminé et d'appliquer les modifications à effectuer sur la variable sélectionner
// la réaction est différente en prenant en compte deux chose :
// - l'argument, qui va déterminé si l'ont doit ajouté ou diminué la valeur
// (valable que pour les int et les float64, sinon il change à l'opposé pour les booléen)
// - le type de la sélection, qui va donc appliquer le changement sur la bonne liste
func EditConfig(add int) {
	if selectedindextype == 0 {
		value := confbool[selectedindexbool]
		*value = !*value
	} else if selectedindextype == 1 {
		value := confint[selectedindexint]
		if add == 1 {
			*value += 10
		} else {
			*value -= 10
		}
	} else if selectedindextype == 2 {
		value := conffloat[selectedindexfloat]
		if add == 1 {

			*value = math.Round((*value+float64(0.1))*100) / 100
		} else {
			*value = math.Round((*value-float64(0.1))*100) / 100
		}
	} else if selectedindextype == 3 {
		if add == 1 {
			if stringcycleindex == len(stringcycle)-1 {
				stringcycleindex = 0
			} else {
				stringcycleindex += 1
			}
			config.General.SpawnType = stringcycle[stringcycleindex]
		} else {
			if stringcycleindex == 0 {
				stringcycleindex = len(stringcycle) - 1
			} else {
				stringcycleindex -= 1
			}
			config.General.SpawnType = stringcycle[stringcycleindex]
		}
	} else if selectedindextype == 4 {
		if add == 1 {
			if stringshcycleindex == len(stringshcycle)-1 {
				stringshcycleindex = 0
			} else {
				stringshcycleindex += 1
			}
			config.General.Shape = stringshcycle[stringshcycleindex]
		} else {
			if stringcycleindex == 0 {
				stringshcycleindex = len(stringshcycle) - 1
			} else {
				stringshcycleindex -= 1
			}
			config.General.Shape = stringshcycle[stringshcycleindex]
		}
	}
}

// cette fonction est la fonction principale de la configuration dynamique
// elle permet de sélectionné une des liste contenant l'adresse des configuration
// ces liste sont triée par type car elle ne sont pas modifié de la même manière (mettre un false à un int ne marche pas par exemple)
// les touches E et A permettent de naviguer entre les types à sélectionné
// les touches flèches droite et gauche permettent de naviguer dans la liste sélectionné avec A et E
// les touches flèches haute et basse permettent de modifier en incrémentant (haut) et en diminuant (bas)
// pour les booléen, la réaction est la même pour flèche haut et bas : elle met la valeur opposé (donc false devient vrai, et vice versa)
// il est à noté que le navigation de chaque liste est indépendante

func ManageDynamicEdit(screen *ebiten.Image) {
	if inpututil.IsKeyJustReleased(ebiten.KeyQ) { //select type -
		if selectedindextype == 0 {
			selectedindextype = 4
		} else {
			selectedindextype -= 1
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyE) { //select type +
		if selectedindextype == 4 {
			selectedindextype = 0
		} else {
			selectedindextype += 1
		}
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) { //select option -
		ManageSelection(false)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) { //select option +
		ManageSelection(true)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) { //select value -
		EditConfig(-1)
	} else if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) { //select value +
		EditConfig(1)
	}
	var optname, optvalue string
	if selectedindextype == 0 {
		optname = confboolname[selectedindexbool]
		optvalue = strconv.FormatBool(*confbool[selectedindexbool])
	} else if selectedindextype == 1 {
		optname = confintname[selectedindexint]
		optvalue = strconv.Itoa(*confint[selectedindexint])
	} else if selectedindextype == 2 {
		optname = conffloatname[selectedindexfloat]
		optvalue = strconv.FormatFloat(*conffloat[selectedindexfloat], 'f', -1, 64)
	} else if selectedindextype == 3 {
		optname = "SpawnType"
		optvalue = config.General.SpawnType
	} else if selectedindextype == 4 {
		optname = "Shape"
		optvalue = config.General.Shape
	}
	//print
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Type: ", typename[selectedindextype], " Nom: ", optname, " Valeur: ", optvalue), 5, 5)
}

//cette fonction gère si la particule sélectionné doit être tué
//elle prend en argument le systeme qui contient la liste des particules, l'élément concerné et le nombre de mort actuel
func KillIfNeeded(s *System, e *list.Element, deathcount *int) {
	if e != nil {
		v := e.Value.(*Particle)
		//condition pour déclarer morte
		if (config.General.OutsideForbidden && !(IsInsideWindow(v.PositionX, v.PositionY, v.ScaleX, v.ScaleY))) || v.Opacity <= 0.05 {
			v.Dead = true
		}
		//recyclage
		if v.Dead {
			v.Opacity = 0
			s.Content.MoveToBack(e)
			*deathcount += 1
		}
	}
}

//
