package main

import (
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "log"
    "net/http"
    "context"
    "time"
    "math/rand"
    "strconv"
)

var rdb *redis.Client
var ctx = context.Background()

func main() {
    r := gin.Default()

    // Connect to Redis
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })

    // Routes
    r.POST("/start", startGame)
    r.POST("/draw", drawCard)
    r.GET("/leaderboard", getLeaderboard)

    // Start the server
    log.Fatal(r.Run(":8080"))
}

func startGame(c *gin.Context) {
    username := c.PostForm("username")
    deck := []string{"cat", "defuse", "shuffle", "exploding", "cat"}
    rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

    // Save deck and reset defuses for user
    rdb.Set(ctx, username+":deck", deck, 0)
    rdb.Set(ctx, username+":defuses", "0", 0)

    c.JSON(http.StatusOK, gin.H{"message": "Game started", "deck": deck})
}

func drawCard(c *gin.Context) {
    username := c.PostForm("username")

    deck, err := rdb.Get(ctx, username+":deck").Result()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Game not started"})
        return
    }

    cards := deckToSlice(deck)
    if len(cards) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "You won!"})
        updateScore(username)
        return
    }

    drawnCard := cards[0]
    newDeck := cards[1:]

    // Shuffle card handling
    if drawnCard == "shuffle" {
        newDeck = []string{"cat", "defuse", "shuffle", "exploding", "cat"}
        rand.Shuffle(len(newDeck), func(i, j int) { newDeck[i], newDeck[j] = newDeck[j], newDeck[i] })
    }

    // Exploding card handling
    if drawnCard == "exploding" {
        defuses, _ := rdb.Get(ctx, username+":defuses").Result()
        defuseCount, _ := strconv.Atoi(defuses)
        if defuseCount > 0 {
            defuseCount--
            rdb.Set(ctx, username+":defuses", strconv.Itoa(defuseCount), 0)
            c.JSON(http.StatusOK, gin.H{"message": "Bomb defused! Draw again.", "deck": newDeck})
        } else {
            c.JSON(http.StatusOK, gin.H{"message": "You lost!", "deck": newDeck})
            return
        }
    }

    // Save updated deck
    rdb.Set(ctx, username+":deck", sliceToDeck(newDeck), 0)
    c.JSON(http.StatusOK, gin.H{"drawnCard": drawnCard, "deck": newDeck})
}

func getLeaderboard(c *gin.Context) {
    users, _ := rdb.Keys(ctx, "*:score").Result()
    leaderboard := make(map[string]string)
    for _, user := range users {
        score, _ := rdb.Get(ctx, user).Result()
        leaderboard[user] = score
    }
    c.JSON(http.StatusOK, leaderboard)
}

func updateScore(username string) {
    rdb.Incr(ctx, username+":score")
}

func deckToSlice(deck string) []string {
    return strings.Split(deck, ",")
}

func sliceToDeck(deck []string) string {
    return strings.Join(deck, ",")
}
