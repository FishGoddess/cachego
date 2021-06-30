 ```mermaid
graph TB
A[Cache]---B1[Segment-1]
A---B2[Segment-2]
A---B3[Segment-3]
A---B4[Segment-...]

B1---D1[data:map]
B1---D2[lock:*sync.Mutex]

D1---k1
k1---v1
D1---k2
k2---v2

```