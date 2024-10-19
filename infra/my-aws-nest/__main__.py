# Pumumi Python program to manage my-aws-nest, provides baseline configurable infra.

import pulumi
from pulumi_aws import s3
import random

# Create a configuration object to access the stack's config.
config = pulumi.Config()

# Get s3_base config.
s3_base = config.require("s3-base") 

# mimic UIIX RAND.
unix_rand = random.randint(0, 32767)

# Create an AWS resource (S3 Bucket)
bucket = s3.Bucket(f"{s3_base}{unix_rand}")
# Export the name of the bucket
pulumi.export('bucket_name', bucket.id)
