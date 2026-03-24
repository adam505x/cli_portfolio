"""
split_earth.py

Reads tui/assets/earth.txt and splits it into individual frame files
in tui/pages/frames_earth/frameXX.txt

Frame structure: 58 lines per frame, separated by 1 blank line (44 frames total).
"""

import os

INPUT = "tui/assets/earth.txt"
OUTPUT_DIR = "tui/pages/frames_earth"

with open(INPUT, "r", encoding="utf-8") as f:
    raw_lines = f.read().splitlines()

FRAME_HEIGHT = 58
PERIOD = FRAME_HEIGHT + 1  # 58 content lines + 1 blank separator

# Validate
expected_frames = (len(raw_lines) + 1) // PERIOD
print(f"Total lines: {len(raw_lines)}")
print(f"Expected frames: {expected_frames}")

os.makedirs(OUTPUT_DIR, exist_ok=True)

for i in range(expected_frames):
    start = i * PERIOD
    frame_lines = raw_lines[start : start + FRAME_HEIGHT]
    filename = os.path.join(OUTPUT_DIR, f"frame{i+1:02d}.txt")
    with open(filename, "w", encoding="utf-8") as f:
        f.write("\n".join(frame_lines) + "\n")

print(f"Written {expected_frames} frames to {OUTPUT_DIR}/")
