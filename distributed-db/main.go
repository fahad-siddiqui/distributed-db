package main

import (
    "github.com/gin-gonic/gin"
    "my-distributed-db/kvstore"
    "net/http"
)

func main() {
    kv := kvstore.NewKVStore()

    r := gin.Default()
    
    //Endpoint to fetch the value of the given key
    r.GET("/key/:key", func(c *gin.Context) {
        key := c.Param("key")
        value, exists := kv.Get(key)
        if exists {
            c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
        } else {
            c.JSON(http.StatusNotFound, gin.H{"message": "Key not found"})
        }
    })
    
    //Endpoint to retrieve all the key-value pairs
    r.GET("/all", func(c *gin.Context) {
        keyValues := kv.GetAllKeyValues()
        c.JSON(http.StatusOK, keyValues)
    })

    //Endpoint to enter a key-value pair
    r.POST("/key", func(c *gin.Context) {
        var data struct {
            Key   string `json:"key"`
            Value string `json:"value"`
        }
        if err := c.ShouldBindJSON(&data); err == nil {
            kv.Set(data.Key, data.Value)
            c.JSON(http.StatusCreated, gin.H{"message": "Key created/updated"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
        }
    })

    // Endpoint to delete a key-value pair by key.
    r.DELETE("/key/:key", func(c *gin.Context) {
        key := c.Param("key")
        kv.Delete(key)
        c.JSON(http.StatusOK, gin.H{"message": "Key deleted"})
    })

    r.Run(":8080")
}