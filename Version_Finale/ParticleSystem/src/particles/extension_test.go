package particles

import (
	"project-particles/config"
	"testing"
)

////////////////////////////////////////////////////////////////
//GRAVITY
////////////////////////////////////////////////////////////////

//gravity = 0
func Test_Gravity_None(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = true, false, false, false, "default"
	config.General.GravityForceY = 0

	s := NewSystem()

	v := s.Content.Front().Value.(*Particle)
	v.DirectionY = 0
	basepos := v.PositionY
	s.Update()
	if v.PositionY != basepos {
		t.Error("la particule bouge alors qu'il n'y a pas de gravite")
	}
}

//gravity = -1
func Test_Gravity_Negative(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = true, false, false, false, "default"

	config.General.GravityForceY = -1

	s := NewSystem()

	v := s.Content.Front().Value.(*Particle)
	v.DirectionY = 0
	basepos := v.PositionY
	s.Update()
	if v.PositionY == basepos {
		t.Error("la particule ne bouge pas alors qu'il y a de la gravite")
	}
}

//gravity = 1
func Test_Gravity_Positive(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = true, false, false, false, "default"

	config.General.GravityForceY = 1

	s := NewSystem()

	v := s.Content.Front().Value.(*Particle)
	v.DirectionY = 0
	basepos := v.PositionY
	s.Update()
	if v.PositionY == basepos {
		t.Error("la particule ne bouge pas alors qu'il y a de la gravite")
	}
}

//gravity = 0
func Test_No_Gravity_None(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	config.General.GravityForceY = 0

	s := NewSystem()

	v := s.Content.Front().Value.(*Particle)
	v.DirectionY = 0
	basepos := v.PositionY
	s.Update()
	if v.PositionY != basepos {
		t.Error("la particule bouge alors qu'il n'y a pas de gravite d'active")
	}
}

//gravity = -1
func Test_No_Gravity_Negative(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	config.General.GravityForceY = -1

	s := NewSystem()

	v := s.Content.Front().Value.(*Particle)
	v.DirectionY = 0
	basepos := v.PositionY
	s.Update()
	if v.PositionY != basepos {
		t.Error("la particule bouge alors qu'il n'y a pas de gravite d'active")
	}
}

//gravity = 1
func Test_No_Gravity_Positive(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	config.General.GravityForceY = 1

	s := NewSystem()

	v := s.Content.Front().Value.(*Particle)
	v.DirectionY = 0
	basepos := v.PositionY
	s.Update()
	if v.PositionY != basepos {
		t.Error("la particule bouge alors qu'il n'y a pas de gravite d'active")
	}
}

////////////////////////////////////////////////////////////////
//OUTSIDE
////////////////////////////////////////////////////////////////

//spawn outside with
func Test_Inside_Inside(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, true, false, false, "default"

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle) 
	v.PositionY = float64(config.General.WindowSizeY)+ 10*v.ScaleY
	s.Update()
	if v.Dead != true {
		t.Error("la particule est considere comme dehors, en etant dedans")
	}
}

//spawn inside with
func Test_Inside_Outside(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, true, false, false, "default"

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)

	v.PositionY = float64(config.General.WindowSizeY)

	s.Update()
	if v.Dead != false {
		t.Error("la particule est considere comme dedans, en etant dehors")
	}
}

//spawn outside with no extension --> useless as default value say its true, so it will not change

//spawn outside with no extension
func Test_No_Inside_Outside(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)

	v.PositionY = float64(config.General.WindowSizeY) + 10*v.ScaleY

	s.Update()
	if v.Dead == true {
		t.Error("la particule est morte en etant dehors alors que l'extension ne marche pas")
	}
}

////////////////////////////////////////////////////////////////
//MIDDLE SPAWN
////////////////////////////////////////////////////////////////
//middle yes
func Test_Middle(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "middle"

	config.General.WindowSizeX, config.General.WindowSizeY = 250, 250

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	if v.PositionX != float64(config.General.WindowSizeX)/2 {
		t.Error("la particule n'est pas au milieu")
	}
}

//middle no
func Test_No_Middle(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	config.General.WindowSizeX, config.General.WindowSizeY = 250, 250

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	if v.PositionX == float64(config.General.WindowSizeX)/2 {
		t.Error("la particule est au milieu et elle ne devrais pas")
	}
}

////////////////////////////////////////////////////////////////
//LIFETIME
////////////////////////////////////////////////////////////////
//lifetime enabled
//lifetime a 0 -> update 1 fois -> check lifetime si = 0
func Test_Lifetime_None(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, true, false, "default"

	config.General.LifetimeMax = 0
	config.General.LifetimeMin = 0

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	s.Update()
	if v.Lifetime != 0 {
		t.Error("la particule est toujours en vie")
	}
}

//lifetime a 2 -> update 2 fois -> check lifetime si = 0
func Test_Lifetime_Pos(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, true, false, "default"

	config.General.LifetimeMax = 1
	config.General.LifetimeMin = 1
	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	s.Update()
	if v.Lifetime != 0 {
		t.Error("la particule est en vie")
	}
}

//lifetime disabled
//lifetime a 0 -> update 1 fois -> check lifetime si = 0
func Test_No_Lifetime_None(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	config.General.LifetimeMax = 0
	config.General.LifetimeMin = 0
	config.General.LifetimeEnabled = false
	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	s.Update()
	if v.Lifetime == 0 {
		t.Error("la particule est morte alors que l'extension n'est pas active", v.Lifetime)
	}
}

//lifetime a 2 -> update 2 fois -> check lifetime si = 0
func Test_No_Lifetime_Pos(t *testing.T) {
	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"

	config.General.LifetimeMax = 1
	config.General.LifetimeMin = 1
	config.General.LifetimeEnabled = false
	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	s.Update()
	if v.Lifetime == 0 {
		t.Error("la particule est morte alors que l'extension n'est pas active", v.Lifetime)
	}
}

////////////////////////////////////////////////////////////////
//DECAY
////////////////////////////////////////////////////////////////
//Test decay couleur
func Test_Decay(t *testing.T) {

	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, true, "default"

	config.General.DecayColorR = 0
	config.General.DecayTime = 1
	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	v.ColorRed = 1
	baser := v.ColorRed

	s.Update()
	if baser == v.ColorRed {
		t.Error("la particule n'a pas change de couleur")
	}
}

//Test pas decay couleur
func Test_No_Decay(t *testing.T) {

	config.General.SpawnRate = 0
	config.General.InitNumParticles = 1
	config.General.GravityEnabled, config.General.OutsideForbidden, config.General.LifetimeEnabled, config.General.DecayEnabled, config.General.SpawnType = false, false, false, false, "default"
	config.General.DecayColorR = 0
	config.General.DecayTime = 1

	s := NewSystem()
	v := s.Content.Front().Value.(*Particle)
	v.ColorRed = 1
	baser := v.ColorRed

	s.Update()
	if baser != v.ColorRed {
		t.Error("la particule a change de couleur alors que l'extension n'est pas active")
	}
}