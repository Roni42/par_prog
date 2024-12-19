
from collections import deque

def seq_bfs(ed, st :int) -> list:
    q = deque([st])
    w = set([st])
    order = [st]
    while q:
        cur = q.popleft()
        w.add(cur)
        for e in ed[cur]:
            if e not in w and e not in q:
                order += [e]
                q.append(e)
    return order