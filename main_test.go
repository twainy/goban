package goban
import "testing"

func TestSet(t *testing.T) {
    _ = Set("test", "test123")
    value,_ := Get("test")
    if value != "test123" {
        t.Error("illegal value", value)
    }
}