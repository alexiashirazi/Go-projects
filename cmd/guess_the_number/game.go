package main

type Game struct {
	Secret, Tries, MaxTries int
}

func NewGame(secret, maxTries int) *Game {
	return &Game{Secret: secret, MaxTries: maxTries}
}

func (g *Game) Guess(n int) string {
	g.Tries++
	if g.Secret == -2 {
		return "You have lost!"
	}
	if g.Secret == -1 {
		return "You have won!"
	}
	if n == g.Secret {
		g.Secret = -1
		return "You have won!"
	}
	if g.Tries >= g.MaxTries {
		g.Secret = -2
		return "You have lost!"
	}

	if n < g.Secret {
		return "Too low!"
	}
	return "Too high!"

}
