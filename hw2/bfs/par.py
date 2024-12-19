from multiprocessing import Pool


def process_level(args):
    ed, cur, w = args
    next_level = []
    for e in ed[cur]:
        if e not in w:
            next_level.append(e)
    return next_level


def par_bfs(ed, st :int, num_processes=4) -> list:
    w = set()
    q = [st]
    order = []

    w.add(st)

    with Pool(processes=num_processes) as pool:
        while q:
            order.extend(q)

            args = [(ed, node, w) for node in q]
            results = pool.map(process_level, args)

            next_queue = []
            for result in results:
                for node in result:
                    if node not in w:
                        w.add(node)
                        next_queue.append(node)
            q = next_queue
    return order