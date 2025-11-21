from bisect import bisect_left, bisect_right
from collections import defaultdict


def build_prefixsums(f, N: int):
    start2cost = defaultdict(int)
    end2time = defaultdict(int)
    start2cost[0] = 0
    end2time[0] = 0

    for i in range(N):
        line = f.readline()
        start, end, cost = [int(x) for x in line.split()]

        start2cost[start] += cost
        end2time[end] += end - start

    s2c_keys = sorted(start2cost.keys())
    e2t_keys = sorted(end2time.keys())

    for i in range(1, len(s2c_keys)):
        start2cost[s2c_keys[i]] += start2cost[s2c_keys[i - 1]]
    for i in range(1, len(e2t_keys)):
        end2time[e2t_keys[i]] += end2time[e2t_keys[i - 1]]

    return start2cost, s2c_keys, end2time, e2t_keys


def binary_search_closest(value: int, sorted_array: list[int], left_bound: bool) -> int:
    if left_bound:
        return bisect_left(sorted_array, value)
    else:
        return bisect_right(sorted_array, value) - 1


def main():
    with open("input.txt") as f:
        N = int(f.readline())
        start2cost, s2c_keys, end2time, e2t_keys = build_prefixsums(f, N)

        # Пропускаем число запросов
        f.readline()

        for line in f:
            start, end, type_ = [int(x) for x in line.split()]

            if type_ == 1:
                left_index = binary_search_closest(start, s2c_keys, True)
                right_index = binary_search_closest(end, s2c_keys, False)
                print(start2cost[s2c_keys[right_index]] - start2cost[s2c_keys[left_index - 1]], end=" ")
            else:
                left_index = binary_search_closest(start, e2t_keys, True)
                right_index = binary_search_closest(end, e2t_keys, False)
                print(end2time[e2t_keys[right_index]] - end2time[e2t_keys[left_index - 1]], end=" ")


if __name__ == "__main__":
    main()
