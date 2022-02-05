#!/usr/bin/env bash

OSCAL_JSONS_URL=https://raw.githubusercontent.com/usnistgov/OSCAL/main/json/schema/oscal_complete_schema.json

# download OSCAL JSON schema
wget $OSCAL_JSONS_URL \
    -O vault/oscal_schema.json

# fix double-escaped regex
sed -i 's/\\\\u/\\u/g' vault/oscal_schema.json
