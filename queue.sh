#!/bin/bash

echo "[" > queue.json
for i in $(seq 1 1000); do
    echo "  {\"id\": $i, \"fileName\": \"index_$i.ts\", \"projectName\": \"project_$i\", \"processed\": false}" >> queue.json
    if [ $i -lt 1000 ]; then
        echo "," >> queue.json
    fi
done
echo "]" >> queue.json