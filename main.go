package main

import (
	"fmt"
	"time"
)

type Game struct{

    isRunning bool

}

func NewGame() *Game{
    return &Game{
        isRunning: false,
    }
}


func (g *Game) Start(){
    g.isRunning = true

    g.gameLoop()
}


func (g *Game) gameLoop(){
    interval := 1 * time.Second

    for {
        
        fmt.Println("The game is Running")

        time.Sleep(interval)

    }
}

type Player struct {
    Health      uint
    Name        string
    AttackPower int
}

func NewPlayer(name string, hp uint, ap int) *Player{
    return &Player{
        Name: name,
        Health: hp,
        AttackPower: ap,

    }
}



//func (player *Player) die(){
func (p *Player) die(){
    p.Health = 0
}



func main(){


    game := NewGame()

    game.Start()
    fmt.Printf("%+v", game)




    //The ampersant is a Pointer with the * in the func
    //player_a := &Player{
    //  Health: 70,
    // Name: "Pichera",
    // AttackPower: 100,
    //}

    //player_b := &Player{
    //    Health: 120,
    //    Name: "Destroyer",
    //    AttackPower: 50,
    //}

    //player_a := NewPlayer("Pichera", 70, 100)
    // player_b := NewPlayer("Destroyer", 120, 50)

    // player_a.die()    
    // player_b.die()    

    // fmt.Println("The player's health  is = ", player_a.Health)

}
