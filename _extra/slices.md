# Slices

```
                s1 := make([]int, 10)        
              array, len=10, cap=10
               │                    
               └──────────────┬──────────────────────────┐
                              │                          │
                              ↓                          ↓ 
            Underlying array [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
               indices        0  1  2  3  4  5  6  7  8  9
                                       ↑        ↑
                                       │        │
                  ┌────────────────────┴────────┘
                  │
               array, len=4 (7-3), cap=7 (10-3)
                 s2 := s1[3:7]

```
