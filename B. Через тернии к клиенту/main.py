class Time:
    def __init__(self):
        self.day = 0
        self.hour = 0
        self.minute = 0

    def add(self, day, hour, minute):
        self.day += day
        self.hour += hour
        self.minute += minute

    def sub(self, day, hour, minute):
        self.day -= day
        self.hour -= hour
        self.minute -= minute

    def count_minutes(self) -> int:
        return (self.day * 24 + self.hour) * 60 + self.minute


class LogEntry:
    def __init__(self, data: list[str]):
        self.day = int(data[0])
        self.hour = int(data[1])
        self.minute = int(data[2])
        self.id = int(data[3])
        self.letter = data[4]


def main(filename: str):
    rockets_time: dict[int, Time] = {}

    with open(filename) as f:
        next(f)

        for line in f:
            log = LogEntry(line.split())

            if log.letter == "B":
                continue

            if log.id not in rockets_time:
                rockets_time[log.id] = Time()

            if log.letter == "A":
                rockets_time[log.id].sub(log.day, log.hour, log.minute)
            else:
                rockets_time[log.id].add(log.day, log.hour, log.minute)

    for key in sorted(rockets_time):
        print(rockets_time[key].count_minutes(), end=" ")


if __name__ == "__main__":
    main("input.txt")
