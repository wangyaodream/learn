import logging


def main() -> None:
    """
    debug的等级会逐级向下，如果level是WARNING，debug和info的内容将不会呈现
    """
    logging.basicConfig(level=logging.DEBUG)


    logging.debug("This is a debug message")
    logging.info("This is a debug message")
    logging.warning("This is a debug message")
    logging.error("This is a debug message")
    logging.critical("This is a debug message")



if __name__ == "__main__":
    main()

