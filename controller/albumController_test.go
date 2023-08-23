package controller_test

// import (
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// )

// func TestGetAlbum(t *testing.T) {
// 	t.Parallel()

// 	gin.SetMode(gin.TestMode)
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	type album struct {
// 		Id     string
// 		Title  string
// 		Artist string
// 		Price  float64
// 	}

// 	tests := []struct {
// 		Name  string
// 		input *gin.Context
// 		Want  []album
// 	}{
// 		{
// 			Name:  "test positive json output of albums",
// 			input: ctrl,
// 			Want: []album{
// 				{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 				{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 				{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.Name, func(t *testing.T) {
// 			t.Parallel()

// 		})
// 	}
// }
