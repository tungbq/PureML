def parse_version_label(label: str):
    version = "latest"
    branch = None
    separator_count = label.count(":")

    if separator_count == 0:
        name = label
    elif separator_count == 1:
        name, branch, version = label.split(":")
    elif separator_count == 2:
        name, branch = label.split(":")

    return name, branch, version
