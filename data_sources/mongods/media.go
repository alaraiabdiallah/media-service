package mongods

import (
	"fmt"
	"github.com/alaraiabdiallah/media-service/helpers"
	"github.com/alaraiabdiallah/media-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mediaCollname string = "media"

func getMediaUrl(media_id interface{}) string {
	base_url := helpers.EnvGetByDefault("APP_BASE_URL","http://localhost")
	id := media_id.(primitive.ObjectID).Hex()
	return fmt.Sprintf("%v/media/%v",base_url, id)
}

func CreateMedia(media interface{}) error{
	coll, dbCtx := collection(mediaCollname)
	defer dbCtx.Disconnect()
	if _, err := coll.InsertOne(dbCtx.ctx, media); err != nil {
		return err
	}
	return nil
}

func FindAllMedia(filter interface{}, results *[]models.MediaDS) error {
	coll, dbCtx := collection(mediaCollname)
	defer dbCtx.Disconnect()
	cur, err := coll.Find(dbCtx.ctx, filter)
	if err != nil { return err }
	for cur.Next(dbCtx.ctx) {
		var res models.MediaDS
		err := cur.Decode(&res)
		if err != nil { return err }
		*results = append(*results, res)
	}
	if err := cur.Err(); err != nil { return err }
	return nil
}

func FindAllMediaLink(filter interface{}, results *[]models.MediaLink) error {
	var media []models.MediaDS
	if err := FindAllMedia(filter, &media); err != nil{return err}
	for _, res := range media {
		*results = append(*results, models.MediaLink{Url: getMediaUrl(res.Id)})
	}
	return nil
}

func FindMedia(filter interface{}, result *models.MediaDS) error {
	coll, dbCtx := collection(mediaCollname)
	defer dbCtx.Disconnect()
	if err := coll.FindOne(dbCtx.ctx,filter).Decode(result); err != nil { return err }
	return nil
}