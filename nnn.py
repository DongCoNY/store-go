import json
import requests

res = requests.post('https://soroban-testnet.stellar.org', json={
    "jsonrpc": "2.0",
    "id": 8675309,
    "method": "getTransaction",
    "params": {
        "hash": "fa83440afa28eb3c02e70da0a6befb38008ec2e0a7956ce74e74340d0a9eab4f"
    }
})
print(json.dumps(res.json(), indent=4))


# startLedger = 487649
# ====================
# res = requests.post('https://soroban-testnet.stellar.org', json={
#     "jsonrpc": "2.0",
#     "id": 8675309,
#     "method": "getEvents",
#     "params": {
#         "startLedger": startLedger,
#         "pagination": {
#             "limit": 10000
#         }
#     }
# })

# print(json.dumps(res.json(), indent=4))

# ===============
# import json, requests
# res = requests.post('https://soroban-testnet.stellar.org', json={
#     "jsonrpc": "2.0",
#     "id": 8675309,
#     "method": "getLedgerEntries",
#     "params": {
#     }
# })
# print(json.dumps(res.json(), indent=4))



