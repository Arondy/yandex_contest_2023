CAP_LIMIT = 10**9


def build_recipes(lines: list[list[int]]):
    unreachable = set()
    saved = {1: [1, 0], 2: [0, 1]}

    for ind, line in enumerate(lines):
        potion = ind + 3
        if potion in unreachable or potion in saved:
            continue

        stack = [potion]
        temp = {potion: [len(line) - 1, [0, 0]]}

        while stack:
            curr_potion = stack[-1]
            curr_index, curr_reqs = temp[curr_potion]

            if curr_index < 0:
                saved[curr_potion] = curr_reqs

                stack.pop()
                temp.pop(curr_potion)

                if stack:
                    parent_potion = stack[-1]
                    parent_reqs = temp[parent_potion][1]

                    parent_reqs[0] += curr_reqs[0]
                    parent_reqs[1] += curr_reqs[1]

                    if parent_reqs[0] > CAP_LIMIT or parent_reqs[1] > CAP_LIMIT:
                        unreachable.update(stack)
                        break
                continue

            recipe_line = lines[curr_potion - 3]
            x = recipe_line[curr_index]
            temp[curr_potion][0] -= 1

            if x in saved:
                curr_reqs[0] += saved[x][0]
                curr_reqs[1] += saved[x][1]

                if curr_reqs[0] > CAP_LIMIT or curr_reqs[1] > CAP_LIMIT:
                    unreachable.update(stack)
                    break
            elif x in unreachable or x in temp:
                unreachable.update(stack)
                break
            else:
                stack.append(x)
                new_line = lines[x - 3]
                temp[x] = [len(new_line) - 1, [0, 0]]

    return saved, unreachable


def main():
    with open("input.txt") as f:
        N = int(f.readline())
        lines = [[int(x) for x in f.readline().split()[1:]] for i in range(N - 2)]
        recipes, unreachable = build_recipes(lines)

        # пропускаем кол-во вопросов
        f.readline()

        for line in f:
            a, b, potion = (int(x) for x in line.split())

            if potion in unreachable:
                print("0", end="")
            elif a >= recipes[potion][0] and b >= recipes[potion][1]:
                print("1", end="")
            else:
                print("0", end="")


if __name__ == "__main__":
    main()
