# 🗜 Clamp

🗜 Clamp is useful to replace environment variables in any file using go template syntax

## How to use from pipe

    echo "{{ .USER }}" | clamp
    # JulienBreux

## How to use from file

    echo "{{ .USER }}" > test.txt
    clamp test.txt
    # /Users/julienbreux
