# Copyright 2019 Google LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""A client for AWS KMS.

Currently works only in Python3 (see Bug 146480447)"""

from __future__ import absolute_import
from __future__ import division
# Placeholder for import for type annotations
from __future__ import print_function

import boto3
from botocore.exceptions import ClientError
import configparser
import os
import re

from typing import Text

from tink.python.aead import aead
from tink.python.integration.awskms.aws_kms_aead import AwsKmsAead
#from tink.python.core import tink_error

AWS_KEYURI_PREFIX = "aws-kms://"

class AwsKmsClient(object):
  """Basic AWS client for AEAD."""

  def __init__(self, key_uri: Text, credentials_path: Text):
    """Creates a new AwsKmsClient that is bound to the key specified in 'key_uri'.

    Uses the specifed credentials when communicating with the KMS. Either of
    arguments can be empty.

    If 'key_uri' is empty, then the client is not bound to any particular key.
    If 'credential_path' is empty, then default credentials will be used.

    Args:
      key_uri: Text, URI of the key the client should be bound to.
      credentials_path: Text, Path to the file with the access credentials.

    Raises:
      ValueError: If the key uri or credentials are invalid.
    """

    match = re.match('aws-kms://arn:aws:kms:([a-z0-9-]+):', key_uri)
    if match:
      self.key_uri = key_uri
      region = key_uri.split(":")[4]
    else:
      raise ValueError

    if not credentials_path:
      kms_client = boto3.client('kms', region_name=region)
    else:
      if not os.path.exists(credentials_path):
        raise ValueError

      if not os.path.isfile(credentials_path):
        raise ValueError

      cred_config = configparser.ConfigParser()
      cred_config.read(credentials_path)
      try:
        if 'aws_access_key_id' in cred_config['default']:
          aws_access_key_id = cred_config['default']['aws_access_key_id']

        if 'aws_secret_access_key' in cred_config['default']:
          aws_secret_access_key = cred_config['default']['aws_secret_access_key']
      except KeyError:
        raise ValueError
      kms_client = boto3.client('kms',
                                aws_access_key_id=aws_access_key_id,
                                aws_secret_access_key=aws_secret_access_key,
                                region_name=region)
      # Make a dummy request to check if the client is valid.
      try:
        kms_client.encrypt(Plaintext="clienttest",
                           KeyId=key_uri[len(AWS_KEYURI_PREFIX):])
      except ClientError:
        raise ValueError
    self.client = kms_client

  def does_support(self, key_uri: Text) -> bool:
    """Returns true iff this client supports KMS key specified in 'key_uri'.

    Args:
      key_uri: Text, URI of the key to be checked.

    Return:
      A boolean value which is true if the key is supported and false otherwise.
    """
    return key_uri.startswith(self.key_uri)

  def get_aead(self, key_uri: Text) -> aead.Aead:
    """Returns an Aead-primitive backed by KMS key specified by 'key_uri'.

    Args:
      key_uri: Text, URI of the key which should be used.

    Returns:
      The AEAD object...
    """

    key_name = key_uri[len(AWS_KEYURI_PREFIX):]
    return AwsKmsAead(key_name, self.client)
