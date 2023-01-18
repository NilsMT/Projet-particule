package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	//spawn
	InitNumParticles int
	SpawnX, SpawnY   int
	SpawnRate        float64
	//extension
	//gravité
	GravityEnabled               bool
	GravityForceX, GravityForceY float64
	//hors de l'écran
	OutsideForbidden bool
	//durée de vie
	LifetimeMax, LifetimeMin float64
	LifetimeEnabled          bool
	//decay
	DecayEnabled                                                                                            bool
	DecayOpacity, DecayRotation, DecayColorR, DecayColorG, DecayColorB, DecayTime, DecayScaleX, DecayScaleY float64
	//input
	KeyInput bool
	//optionnel
	ShowConfigStatus bool
	//methode de spawn
	SpawnType, Shape string
	ShapeSize        float64
	//default value
	DefaultRotation                             float64
	DefaultColorR, DefaultColorG, DefaultColorB float64
	DefaultScaleX, DefaultScaleY                float64
	DefaultOpacity                              float64
	//interval value
	RangeValue                                                                                                                     bool
	RangeRotation, RangeDirectionX, RangeDirectionY, RangeScaleY, RangeScaleX, RangeColorR, RangeColorG, RangeColorB, RangeOpacity float64
}

var General Config
