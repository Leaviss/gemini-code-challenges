def prompt_user() -> str:
    """Ask the user for a noun, verb, and adjective."""
    noun = input("Provide a noun: ")
    verb = input("Provide a verb: ")
    adj = input("Provide a adj: ")

    return print(f"The {adj} {noun} likes to {verb} the cloud server.")

def main():
    """execute main program"""
    print("I need your help telling a story.")
    prompt_user()


if __name__ == "__main__":
    main()
    