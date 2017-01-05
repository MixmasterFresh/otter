package operation

import(
    "testing"
)

func BenchmarkPush(b *testing.B) {
    buf := &operationBuffer{}
    buf.init()
    op := &Operation{}
    for n := 0; n < b.N; n++ {
		buf.push(op)
	}
}

func BenchmarkPushPushPop(b *testing.B){
    buf := &operationBuffer{}
    buf.init()
    op := &Operation{}
    buf.push(op)
    for n := 0; n < b.N; n++ {
		buf.push(op)
        buf.pop()
	}
}

func BenchmarkPushPop(b *testing.B){
    buf := &operationBuffer{}
    buf.init()
    op := &Operation{}
    for n := 0; n < b.N; n++ {
		buf.push(op)
        buf.pop()
	}
}


