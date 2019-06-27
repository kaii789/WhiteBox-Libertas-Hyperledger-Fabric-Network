# Copyright 2019 Sipher Inc.
#
# SPDX-License-Identifier: Apache-2.0
#
# This script uninstalls the test app by bringing down the test network and deleting any keys in the wallet.

# Exit on first error
set -e

# Set relative path correctly
cd $(dirname "$0")

# Bring down test network
cd ../libertas-dev-network
./downDevelopmentNetwork.sh

# Remove keys in wallet
cd ../app/javascript/
rm -rf ./wallet/*

echo "Squeaky clean!"