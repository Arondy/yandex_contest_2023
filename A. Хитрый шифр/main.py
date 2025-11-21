class Candidate:
    def __init__(self, data: list[str]):
        self.surname, self.name, self.patronymic, self.birth_day, self.birth_month, self.birth_year = data

    def create_code(self) -> str:
        number_sum = len(set(self.surname + self.name + self.patronymic))

        for let in self.birth_day + self.birth_month:
            number_sum += int(let) * 64

        number_sum += (ord(self.surname[0].lower()) - ord('a') + 1) * 256
        code = hex(number_sum).upper()[2:]
        code = code[-3:]

        if len(code) < 3:
            code = (3 - len(code)) * '0' + code

        return code


def main(filename: str = "input.txt"):
    with open(filename) as f:
        next(f)

        for line in f:
            candidate = Candidate(line.strip().split(","))
            print(candidate.create_code(), end=" ")


if __name__ == "__main__":
    main()
