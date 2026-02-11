def prompt_password() -> str:
    password = input("Enter Password: ")
    return password


def main():
    secret_password = "secret123"
    password = ""
    while password != secret_password:
        password = prompt_password()
        if password != secret_password:
            print("Incorrect, try again.")
        else:
            print("Access Granted!")
    
    


if __name__ == "__main__":
    main()