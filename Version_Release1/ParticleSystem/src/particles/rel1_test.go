package particles

import (
	"testing"
	"project-particles/config"
)

////////////////////////////////////////////////////////////////
//SPAWN INIT
////////////////////////////////////////////////////////////////
func Test_Spawn_None(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=0
	config.General.SpawnRate=0
	s := NewSystem()
	if s.Content.Len()!=0{
		t.Error("Il ne devrais pas avoir de particule, mais il y en a",s.Content.Len())
	}
}

func Test_Spawn_Negative(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=-1
	config.General.SpawnRate=0
	s := NewSystem()
	if s.Content.Len()!=0{
		t.Error("Il ne devrais pas avoir de particule, mais il y en a",s.Content.Len())
	}
}

func Test_Spawn_Positive(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=1
	config.General.SpawnRate=0
	s := NewSystem()
	if s.Content.Len()!=1{
		t.Error("Il ne devrais avoir 1 particule, mais il y en a",s.Content.Len())
	}
}
////////////////////////////////////////////////////////////////
//GENERATE
////////////////////////////////////////////////////////////////
func Test_Generate_None(t *testing.T) {
	config.General.InitNumParticles=0
	config.General.SpawnRate=0
	s := NewSystem()
	s.Update()
	if s.Content.Len()!=0{
		t.Error("Il ne devrais pas avoir particule, mais il y en a",s.Content.Len())
	}
}

func Test_Generate_Negative(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=0
	config.General.SpawnRate=-1
	s := NewSystem()
	s.Update()
	if s.Content.Len()!=0{
		t.Error("Il ne devrais pas avoir particule, mais il y en a",s.Content.Len())
	}
}

func Test_Generate_Positive(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=0
	config.General.SpawnRate=1
	s := NewSystem()
	s.Update()
	if s.Content.Len()!=1{
		t.Error("Il devrais avoir 1 particule, mais il y en a",s.Content.Len())
	}
}

func Test_Generate_Float(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=0
	config.General.SpawnRate=0.5
	s := NewSystem()
	s.Update()
	s.Update()
	if s.Content.Len()!=1{
		t.Error("Il devrais avoir 1 particule, mais il y en a",s.Content.Len())
	}
}
////////////////////////////////////////////////////////////////
//RANDOM SPAWN
////////////////////////////////////////////////////////////////
func Test_RandomSpawn(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=1
	config.General.SpawnRate=0
	config.General.RandomSpawn=true
	s := NewSystem()
	particle := s.Content.Front().Value.(*Particle)
	if particle.PositionX==float64(config.General.SpawnX) || particle.PositionY==float64(config.General.SpawnY){
		t.Error("Il devrais avoir une position random pour la particule, mais il est en X",particle.PositionX,"Y",particle.PositionY)
	}
}

func Test_No_RandomSpawn(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=1
	config.General.SpawnRate=0
	config.General.RandomSpawn=false
	s := NewSystem()
	particle := s.Content.Front().Value.(*Particle)
	if float64(particle.PositionX)!=float64(config.General.SpawnX) || float64(particle.PositionY)!=float64(config.General.SpawnY){
		t.Error("Il devrais avoir une position donné pour la particule, mais il est en n'est pas aux bonnes coordonnées:",particle.PositionX,particle.PositionY)
	}
}
////////////////////////////////////////////////////////////////
//MOVE Y
////////////////////////////////////////////////////////////////
func Test_Move(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=1
	config.General.SpawnRate=0
	config.General.RandomSpawn=false
	s := NewSystem()
	s.Update()
	particle := s.Content.Front().Value.(*Particle)
	if particle.PositionX==float64(config.General.SpawnX) || particle.PositionY==float64(config.General.SpawnY){
		t.Error("Il devrais avoir un changement de position pour la particule, mais elle n'a pas changé")
	}
}
func Test_RandomMove(t *testing.T) {
	config.Get("../config.json")
	config.General.InitNumParticles=2
	config.General.SpawnRate=0
	config.General.RandomSpawn=false
	s := NewSystem()
	s.Update()
	var tab []float64
	for e := s.Content.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Particle)
		tab = append(tab,v.VelocityX)
		tab = append(tab,v.VelocityY)
	}
	if tab[0] == tab[2] || tab[1] == tab[3] {
		t.Error("Il devrais avoir un changement de position différent pour chaque particule, mais elles se déplace de manière identiques")
	}
}