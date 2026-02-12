

def request_server() -> str:
    """A user is given a prompt for requesting a server."""
    server = input("Enter a new server name: ")
    return server

def manage_servers(server: str) -> list:
    """This function will manage a list of servers."""
    servers = []
    static_servers = ["darkstar", "ai-core", "omni-1"]
    for s in static_servers:
        servers.append(s)
    servers.append(server)
    return servers


def main():
    """This is my main function."""
    server = request_server()
    servers = manage_servers(server)
    print("Server List:")

    for i, server in enumerate(servers):
        print(f"[{i}] {server} is online!")

    

if __name__ == "__main__":
    main()