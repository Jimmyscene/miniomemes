import boto3

from flask import Flask, jsonify
from botocore.client import Config


app = Flask(__name__)


@app.route('/')
def hello():
    client = boto3.client(
            service_name='s3',
            # * Minio expects no protocol schema, boto3 expects to have it.
            endpoint_url="http://minio:9000",
            aws_access_key_id="accesskey",
            aws_secret_access_key="secretkey",
            config=Config(signature_version='s3v4')
    )
    data = {
        bucket["Name"]: {
            object["Key"]: client.generate_presigned_url(
                "get_object", Params={
                    "Bucket": bucket["Name"],
                    "Key": object["Key"]
                }
            ).replace("http://minio:9000/", "/minio/")
            for object in client.list_objects(Bucket=bucket["Name"]).get("Contents", [])
        }
        for bucket in client.list_buckets()["Buckets"]
    }
    return jsonify(data)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8080)
