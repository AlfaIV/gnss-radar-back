echo Test NOW
cat ./gnss-satellites/exampleReq/now  | curl -X POST http://localhost:8000/api/v1/satellites/now \
    -H "Content-Type: application/json" \
    -d @-

# echo Test TIME
# cat ./gnss-satellites/exampleReq/time  | curl -X POST http://localhost:8000/api/v1/satellites/time \
#     -H "Content-Type: application/json" \
#     -d @-


# >>> import time
# >>> time.time()
# 1744656431.9362886