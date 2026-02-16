## Day 5: The Inventory System (Key-Value Stores)

**The Goal:**
Create a CLI tool to manage a lookup table of server names and their IP addresses using a key-value data structure (Map or Dictionary).

**Requirements:**
1. **Initialize** a key-value data structure to store Server Names (keys) and IP Addresses (values).
2. **Pre-populate** the structure with at least one known server (e.g., "darkstar" -> "192.168.1.1").
3. **Prompt** the user to enter a server name to look up.
4. **Check** if the server name exists in the data structure:
   - **If found:** Print the associated IP address.
   - **If not found:** Print "Server not found", prompt the user to enter an IP address for that server, and add the new pair to the data structure.
5. **Print** the updated list of all servers and IPs.

**Example Output:**
```text
Enter server name: redis-cache
Server not found.
Enter IP for redis-cache: 10.0.0.99
Added: redis-cache -> 10.0.0.99

Current Inventory:
darkstar: 192.168.1.1
redis-cache: 10.0.0.99