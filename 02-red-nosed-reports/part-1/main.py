safe_reports = 0

with open("./input") as f:
    for line in f.readlines():
        split_line = line.split()

        all_increasing = None
        is_report_safe = True

        for i in range(1, len(split_line)):
            prev = int(split_line[i - 1])
            current = int(split_line[i])

            distance = abs(current - prev)

            is_increasing = current - prev > 0

            if all_increasing is None:
                all_increasing = is_increasing

            if distance < 1 or distance > 3:
                is_report_safe = False
                break

            if is_increasing is not all_increasing:
                is_report_safe = False
                break

        if is_report_safe:
            safe_reports = safe_reports + 1

print(safe_reports)
