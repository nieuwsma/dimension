package api

import (
	"dimension/internal/session"
	"github.com/gin-gonic/gin"
)

var TokenProvider session.SessionProvider

// todo come back to this function... i need to be able to handle tokens across the hierarchy
// either add a var to this to pass in what context is needed, OR create separate storage functions
func AuthMiddleware(context string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//session := c.GetHeader("x-auth-session")
		//
		//// Basic session presence check
		//if session == "" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "x-auth-session required"})
		//	c.Abort()
		//	return
		//}
		//
		//gameID, playerID, trainID := extractResourceDetails(c)
		//
		//// Check if the session is valid and fetch associated resources
		//allowedResources, err := fetchResourcesForToken(session)
		//if err != nil {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
		//	c.Abort()
		//	return
		//}
		//
		//// Extract resource details from the path
		//resourceType, resourceID := extractResourceFromPath(c.FullPath())
		//
		//// Check if the session has permission for the resource
		//if !isTokenAllowedForResource(allowedResources, resourceType, resourceID) {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token doesn't have access to this resource"})
		//	c.Abort()
		//	return
		//}

		c.Next() //this makes sure we keep processing it!
	}
}

func extractResourceDetails(c *gin.Context) (gameID, playerID, trainID string) {
	gameID = c.Param("gameId")
	playerID = c.Param("playerId")
	trainID = c.Param("trainID")
	return
}
