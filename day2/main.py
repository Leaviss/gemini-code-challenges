def check_access() -> str:
    """Ask the user for their role."""
    
    role = input("Enter your role: ")

    if role.lower() == "admin":
        return print("Full Access Granted")
    elif role.lower() == "manager":
        return print("Limited Access Granted")
    else:
        return print("Access Denied")

def main():
    """execute main program"""
    check_access()


if __name__ == "__main__":
    main()