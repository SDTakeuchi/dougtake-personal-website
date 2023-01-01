def is_empty(v) -> bool:
    return \
        v == 0 or \
        v == '' or \
        v is None or \
        len(v) == 0
