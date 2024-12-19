
import time
import numpy as np
from bfs.seq import seq_bfs
from bfs.par import par_bfs


def generate_cubic_graph(side_length):
    graph = {}
    for x in range(side_length):
        for y in range(side_length):
            for z in range(side_length):
                neighbors = []
                for dx, dy, dz in [(1, 0, 0), (-1, 0, 0), (0, 1, 0),
                                   (0, -1, 0), (0, 0, 1), (0, 0, -1)]:
                    nx, ny, nz = x + dx, y + dy, z + dz
                    if 0 <= nx < side_length and 0 <= ny < side_length and 0 <= nz < side_length:
                        neighbors.append((nx, ny, nz))
                graph[(x, y, z)] = neighbors
    return graph


def test_correctness():
    small_graph = {
        (0, 0): [(0, 1), (1, 0)],
        (0, 1): [(0, 0), (1, 1)],
        (1, 0): [(0, 0), (1, 1)],
        (1, 1): [(0, 1), (1, 0)]
    }
    start_node = (0, 0)

    result_seq = seq_bfs(small_graph, start_node)
    result_par = par_bfs(small_graph, start_node)

    # print(*result_par, *result_seq)
    assert result_seq == result_par, "Diff"

    side_length = 10
    large_graph = generate_cubic_graph(side_length)
    start_node = (0, 0, 0)
    
    result_seq = seq_bfs(large_graph, start_node)
    result_par = par_bfs(large_graph, start_node)

    assert result_seq == result_par, "Diff"

    print("Same results!")


def test():
    side_length = 10
    start_node = (0, 0, 0)
    graph = generate_cubic_graph(side_length)

    seq_times = []
    for i in range(5):
        start_time = time.time()
        seq_bfs(graph, start_node)
        elapsed_time = time.time() - start_time
        seq_times.append(elapsed_time)

    par_times = []
    for i in range(5):
        start_time = time.time()
        par_bfs(graph, start_node, num_processes=4)
        elapsed_time = time.time() - start_time
        par_times.append(elapsed_time)

    seq_avg_time = sum(seq_times) / len(seq_times)
    par_avg_time = sum(par_times) / len(par_times)

    speedup = seq_avg_time / par_avg_time if par_avg_time > 0 else float('inf')
    
    return seq_avg_time, par_avg_time, speedup

if __name__ == "__main__":
    print("Correctness...")
    test_correctness()
    print("Perfect!")

    print("\nTests:")
    seq_avg_time, par_avg_time, speedup = test()

    print(f"Av time (seq): {seq_avg_time:.4f} s")
    print(f"Av time (par): {par_avg_time:.4f} s")
    print(f"{speedup:.2f}x")



