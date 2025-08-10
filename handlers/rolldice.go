package handlers

// boilerplate incomming from opentelemetry documentation at https://opentelemetry.io/docs/languages/go/getting-started/

import (
	"kauanpecanha/devops-challenge/db"
	"kauanpecanha/devops-challenge/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

const name = "go.opentelemetry.io/otel/example/dice"

// opentelemetry integration within the rolldice feature
var (
	tracer  = otel.Tracer(name)
	meter   = otel.Meter(name)
	logger  = otelslog.NewLogger(name)
	rollCnt metric.Int64Counter
)

func init() {
	var err error
	rollCnt, err = meter.Int64Counter("dice.rolls",
		metric.WithDescription("The number of rolls by roll value"),
		metric.WithUnit("{roll}"))
	if err != nil {
		panic(err)
	}
}

func Play(c *gin.Context) {
	ctx, span := tracer.Start(c, "play")
	defer span.End()

	player := c.Param("player")

	if player == "" {
		player = "Anonymous Player"
	}

	roll := models.Roll{
		ID:        primitive.NewObjectID(),
		Player:    player,
		Result:    1 + rand.Intn(6),
		Timestamp: time.Now(),
	}

	_, err := db.RollsCollection.InsertOne(ctx, roll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no MongoDB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"player":    roll.Player,
		"result":    roll.Result,
		"timestamp": roll.Timestamp,
	})

}

func GetAllPlays(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "GetAllPlays")
	defer span.End()

	cursor, err := db.RollsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar rolls"})
		return
	}
	defer cursor.Close(ctx)

	rolls := make([]models.Roll, 0)
	if err := cursor.All(ctx, &rolls); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar rolls"})
		return
	}

	c.JSON(http.StatusOK, rolls)
}

func GetPlayByID(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "GetPlayByID")
	defer span.End()

	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var roll models.Roll
	err = db.RollsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&roll)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Roll não encontrado"})
		return
	}

	c.JSON(http.StatusOK, roll)
}

func UpdatePlay(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "UpdatePlay")
	defer span.End()

	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var roll models.Roll
	if err := c.ShouldBindJSON(&roll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.RollsCollection.UpdateOne(ctx,
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"result": roll.Result, "timestamp": time.Now()}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Atualizado com sucesso"})
}

func DeletePlay(c *gin.Context) {
	ctx, span := tracer.Start(c.Request.Context(), "DeletePlay")
	defer span.End()

	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	_, err = db.RollsCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar roll"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roll deletado com sucesso"})
}

func WelcomeHome(context *gin.Context) {
	context.String(http.StatusOK, "Hello, World!")
}
