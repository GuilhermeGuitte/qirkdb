package database

import (
    "unsafe"
    "fmt"
    "sync"
)

type Row struct {
    value string
    ttl   int
}

type Database struct {
    table map[string]Row
    sync.RWMutex
}

/**
 * Constructor
 */
func New() *Database {
    return &Database{
        table: make(map[string]Row),
    }
}

/**
 * Set key => value to database.
 */
func (database *Database) Set(key string, value string) {
    database.Lock()
    defer database.Unlock()

    row      := Row{}
    row.value = value

    database.table[key] = row
}

/**
 * Get the key given.
 */
func (database Database) Get(key string) (string, bool) {
    database.RLock()
    defer database.RUnlock()

    value := database.table[key].value
    empty := len(value) == 0

    return value, empty
}

/**
 * Delete a Key given at database.
 */
func (database *Database) Del(key string) {
    database.Lock()
    defer database.Unlock()

    delete(database.table, key)
}

/**
 * Verify if a Key exists
 */
func (database *Database) Exists(key string) (bool) {
    _ , empty := database.Get(key)

    return !empty
}

/**
 * Setting ttl for a key
 */
func (database *Database) setTTL(key string, ttl int) (bool) {
    _ , empty := database.Get(key)

    if empty {
        return false
    }

    row    := database.table[key]
    row.ttl = ttl
    database.table[key] = row

    return true
}

/**
 * Getting ttl for a key
 */
func (database *Database) getTTL(key string) (int, bool) {
    _ , empty := database.Get(key)

    return database.table[key].ttl, empty
}

/**
 * Return a size of the value
 */
func (database Database) getSize(key string) (string) {
    value , _ := database.Get(key)

    size := (float64(unsafe.Sizeof(value)) / 1024)

    return fmt.Sprintf("%f KB", size)
}
