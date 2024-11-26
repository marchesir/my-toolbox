# Pumumi Python program to manage my-aws-nest, provides baseline configurable infra.
import pulumi
from pulumi_aws import s3
import random
import re

# Globals.
AWS_REGION_REGEX_STRICT = r"^(us|eu|ap|ca|sa)-(east|west|north|south|central|southeast|northeast)-\d{1}$"
ADJECTIVES = ["Cool","Fast","Bold","Calm","Sharp","Quick","Keen","Brave","Happy","Proud"]
NOUNS = ["Panda","Tiger","Wolf","Bear","Eagle","Fox","Deer","Owl","Lion","Hawk"]


def generate_docker_style_name() -> str:
    """
    Generates a random Docker-style name (e.g., 'fluffy-lion').
    """
    adjective = random.choice(ADJECTIVES)
    noun = random.choice(NOUNS)
    return f"{adjective}-{noun}"


def is_valid_aws_region(region: str) -> bool:
    """
    Validates if the given string is a valid, known AWS region code.
    
    Args:
        region (str): The AWS region code to validate.
    
    Returns:
        bool: True if valid, False otherwise.
    """
    return bool(re.match(AWS_REGION_REGEX_STRICT, region))


def validate_and_apply_logic(required_keys: list[str]) -> dict:
    """
    Validates required keys in Pulumi config and applies logic based on the configuration values.

    Args:
        required_keys (list[str]): A list of configuration keys that must be present.
    
    Returns:
        dict: A dictionary of processed configuration values.

    Raises:
        KeyError: If a required key is missing from the configuration.
    """
    
    config = pulumi.Config(pulumi.get_project())    
    stack_config = {}

    # Validate required keys.
    for key in required_keys:
        value = config.get(key)
        if value is None:
            raise KeyError(f"Required configuration key '{key}' is missing.")
        stack_config[key] = value

    # Apply logic based on configuration values.
    processed_config = {}
    for key, value in stack_config.items():
        if key == "region" and is_valid_aws_region(value): 
            print("good")            
        #elif key == "enable_feature_x":
        #    # Example logic: Convert "true"/"false" strings to booleans
        #    processed_config["enable_feature_x"] = value.lower() == "true"
        #else:
            # Store unprocessed keys directly
        processed_config[key] = value

    return processed_config


# Example usage in Pulumi program
if __name__ == "__main__":
    try:
        # Specify the required keys        
        required_keys = ["region", "service", "count"]     
        # Validate and apply logic
        processed_config = validate_and_apply_logic(required_keys)
        # Output the processed config
        pulumi.log.info(f"Processed Pulumi configuration: {processed_config}")

    except KeyError as e:
        pulumi.log.error(str(e))
    except ValueError as e:
        pulumi.log.error(str(e))
    except Exception as e:
        pulumi.log.error(f"Unexpected error: {e}")









# Create a configuration object to access the stack's config.
#config = pulumi.Config()

# Get s3_base config.
#s3_base = config.require("myapp") 

# mimic UIIX RAND.
#unix_rand = random.randint(0, 32767)

# Create an AWS resource (S3 Bucket)
#bucket = s3.Bucket(f"{s3_base}{unix_rand}")
# Export the name of the bucket
#pulumi.export('bucket_name', bucket.id)
     

