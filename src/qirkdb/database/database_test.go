package database

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestSetKeyValueAtDatabase(t *testing.T) {
    // Range
    database := New()
    database.Set("key", "value")

    // Act
    value, _ := database.Get("key")

    // Assert
    assert.Equal(t, "value", value)
}

func TestReturnFalseIfRemoveNotExistsKey(t *testing.T) {
    // Range
    database := New()

    // Act
    database.Del("key")
    _, empty := database.Get("key")

    // Assert
    assert.True(t, empty)
}

func TestReturnTrueRemovingAKey(t *testing.T) {
    // Range
    database := New()

    // Act
    database.Set("key", "value")
    database.Del("key")
    _, empty := database.Get("key")

    // Assert
    assert.True(t, empty)
}

func TestVerifyIfExists(t *testing.T) {
    // Range
    database := New()

    // Act
    var result bool = database.Exists("Key")

    assert.False(t, result)

    database.Set("key", "value")

    result = database.Exists("key")

    // Assert
    assert.True(t, result)
}

func TestSetTllForExistentARow(t *testing.T) {
    // Range
    database := New()
    database.Set("key", "value")

    // Act
    database.setTTL("key", 100) // seconds

    // Assert
    result, _ := database.getTTL("key")
    assert.Equal(t, 100, result)
}


func TestSetTllForNoExistentARow(t *testing.T) {
    // Range
    database := New()

    // Act
    var isOk bool =  database.setTTL("key", 100) // seconds

    // Assert
    assert.False(t, isOk)
}
