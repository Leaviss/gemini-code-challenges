

def manage_inventory():
    """Manage the server inventory"""
    servers = {
    "darkstar" : "192.168.1.1"
    }

    s = input("Enter server name: ")

    if s in servers:
        print(f"Server IP for {s} is {servers[s]}\n")
    else:
        print("Server not found.")
        ip = input(f"Enter IP for {s}: ")
        servers[s] = ip
        print(f"Added: {s} -> {ip}\n")
    
    print("Current Inventory: ")
    for s in servers:
        print(f"{s}: {servers[s]}")


def main():
    """This is the programs main function."""
    manage_inventory()


if __name__ == "__main__":
    main()