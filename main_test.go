package goban
import "testing"

func TestSet(t *testing.T) {
    Setup("./etc/redis.json")
    _ = Set("test", "test123")
    value,_ := Get("test")
    if value != "test123" {
        t.Error("illegal value", value)
    }
}
func TestZrevrank(t *testing.T) {
    Setup("./etc/redis.json")
    key := "rank_test"
    m := map[string]int {
        "Mike":  10,
        "Jobs":  50,
        "Megu": 100,
        "Jake":  70,
    }
    for k,v := range m {
        Zadd(key,v,k)
    }
    v,_ := Zrevrank(key, "Mike")
    if (v != 3) {
        t.Error("illegal value Mike", v)
    }
    v,_ = Zrevrank(key, "Jobs")
    if (v != 2) {
        t.Error("illegal value Jobs", v)
    }
    v,_ = Zrevrank(key, "Megu")
    if (v != 0) {
        t.Error("illegal value Megu", v)
    }
    v,_ = Zrevrank(key, "Jake")
    if (v != 1) {
        t.Error("illegal value Jake", v)
    }
}