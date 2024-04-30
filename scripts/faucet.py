from web3 import Web3

# Connect to Ganache
ganache_url = "http://localhost:8545"
web3 = Web3(Web3.HTTPProvider(ganache_url))

# Check if the connection is successful
if web3.is_connected():
    print("Connected to Ganache")
else:
    print("Failed to connect to Ganache")

def send_ether():
    # Fetch accounts from Ganache
    accounts = web3.eth.accounts

    # Define sender and receiver
    sender = accounts[0]  # Typically the first account is used
    receiver = "0xc0E3Cad8CAF06C53588EFE33D160316a7B9e7cE8"  # Replace with your address
    amount = web3.to_wei(10, 'ether')  # Amount to send

    # Send Ether
    txn_hash = web3.eth.send_transaction({
        'from': sender,
        'to': receiver,
        'value': amount
    })

    # Wait for transaction to be mined
    receipt = web3.eth.wait_for_transaction_receipt(txn_hash)

    print(f"Transaction successful with hash: {receipt.transactionHash.hex()}")

if __name__ == "__main__":
    send_ether()
