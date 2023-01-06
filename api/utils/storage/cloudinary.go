package storage

import (
	"context"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryClient struct {
	CloudinaryClient *cloudinary.Cloudinary
}

func newCloudinaryStorage() (*CloudinaryClient, error) {
	cld, err := cloudinary.New()
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
		return nil, err
	}
	return &CloudinaryClient{CloudinaryClient: cld}, err
}

func (cs *CloudinaryClient) UplodaImage(ctx context.Context, imagePath string, publicID string) (*uploader.UploadResult, error) {
	uploadResult, err := cs.CloudinaryClient.Upload.Upload(
		ctx,
		imagePath,
		uploader.UploadParams{PublicID: publicID})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
		return nil, err
	}
	return uploadResult, err
}

// func Storage() {
// 	// Prints something like:
// 	// https://res.cloudinary.com/<your cloud name>/image/upload/v1615875158/logo.png

// 	// uploadResult contains useful information about the asset, like Width, Height, Format, etc.
// 	// See uploader.UploadResult struct for more details.

// 	// Now we can use Admin API to see the details about the asset.
// 	// The request can be customised by providing AssetParams.
// 	asset, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "logo"})
// 	if err != nil {
// 		log.Fatalf("Failed to get asset details, %v\n", err)
// 	}

// 	// Print some basic information about the asset.
// 	log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)

// 	// Cloudinary also provides a very flexible Search API for filtering and retrieving
// 	// information on all the assets in your product environment with the help of query expressions
// 	// in a Lucene-like query language.
// 	searchQuery := search.Query{
// 		Expression: "resource_type:image AND uploaded_at>1d AND bytes<1m",
// 		SortBy:     []search.SortByField{{"created_at": search.Descending}},
// 		MaxResults: 30,
// 	}

// 	searchResult, err := cld.Admin.Search(ctx, searchQuery)

// 	if err != nil {
// 		log.Fatalf("Failed to search for assets, %v\n", err)
// 	}

// 	log.Printf("Assets found: %v\n", searchResult.TotalCount)

// 	for _, asset := range searchResult.Assets {
// 		log.Printf("Public ID: %v, URL: %v\n", asset.PublicID, asset.SecureURL)
// 	}
// }
