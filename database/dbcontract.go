package database

import "fiscaluno-ws/database/filter"

type DBcontract interface {
    Save(object interface{})
    Find(node string, conditions[] filter.Filter) (interface{}, error)
    Update(newValue, conditions[] filter.Filter)
    Delete(conditions[] interface{})
}